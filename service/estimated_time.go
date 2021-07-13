package service

import (
	"context"
	"github.com/google/uuid"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/tec"

	logger "github.com/sirupsen/logrus"
)

// ALGORITHM
// 1. Get Stage ID from Step
// 2. Fetch Template ID from this Stage ID
// 3. Fetch All Stages and Steps from Template ID
// 4. Iterate over the stages and steps and calculate time accordingly
// 5. Update time in DB
func updateEstimatedTimeByStageID(ctx context.Context, s db.Storer, stageID uuid.UUID) (err error) {
	stage, err := s.ShowStage(ctx, stageID)
	if err != nil {
		logger.Errorln(err)
		return
	}
	return updateEstimatedTimeByTemplateID(ctx, s, stage.TemplateID)
}

func getHomingAndLidTempTime(ctx context.Context, lidTemp int64, estimatedTime *float64) (err error) {

	// Calculate Homing Time as its included in Experiment Time
	*estimatedTime += float64(config.GetHomingTime())
	logger.Infoln("Estimated Time for Homing RTPCR: ", config.GetHomingTime())

	// Calculate Lid Temp Time
	// NOTE: This is where most variance exists for estimated time
	// TODO: Handle this in a better and accurate way
	// here 0.5 is the rate of heating/ cooling per sec
	*estimatedTime += math.Abs(float64(lidTemp)-config.GetRoomTemp()) / 0.5
	logger.Infoln("Estimated Time for Lid Temp Reaching: ", math.Abs(float64(lidTemp)-config.GetRoomTemp())/0.5)
	return
}

func updateEstimatedTimeByTemplateID(ctx context.Context, s db.Storer, templateID uuid.UUID) (err error) {

	temp, err := s.ShowTemplate(ctx, templateID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Templates Data")
		return
	}

	roomTemp := config.GetRoomTemp()
	currentTemp := roomTemp

	// Set Only Lid Temp and Homing Time
	var estimatedTime float64
	getHomingAndLidTempTime(ctx, temp.LidTemp, &estimatedTime)

	ss, err := s.ListStages(ctx, templateID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Stages Data")
		return
	}

	for _, stage := range ss {
		// Get stepID from first stage
		steps, err := s.ListSteps(ctx, stage.ID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching Steps Data")
			return err
		}

		// Calculate Hold Stage
		tp := 0.0
		for _, st := range steps {
			tp += math.Abs(currentTemp-float64(st.TargetTemperature)) / float64(st.RampRate)
			tp += float64(st.HoldTime)
			currentTemp = float64(st.TargetTemperature)
		}

		if stage.Type == cycle {
			estimatedTime += tp * float64(stage.RepeatCount)
		} else {
			estimatedTime += tp
		}
	}

	// Last step to go back to Room Temp
	estimatedTime += math.Abs(currentTemp-roomTemp) / tec.RoomTempRamp
	logger.Infoln("Estimated Time : ", estimatedTime)

	return s.UpdateEstimatedTime(ctx, templateID, int64(estimatedTime))
}
