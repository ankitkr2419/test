package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func pidCalibrationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]

		if deps.PlcDeck[deck].IsRunInProgress() {
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].PIDCalibration(req.Context())

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStartHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if plc.LidPidTuningInProgress {
			logger.Errorln(responses.LidPIDTuningPresentError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.LidPIDTuningPresentError.Error()})
			return
		}

		// TODO: Logging this API
		go deps.Plc.LidPIDCalibration()

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStopHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if !plc.LidPidTuningInProgress {
			logger.Errorln(responses.LidPidTuningStartError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.LidPidTuningStartError.Error()})
			return
		}

		// Stop the lid PID Tuning
		err := deps.Plc.Stop()
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.LidPidTuningNotOffError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, MsgObj{Msg: responses.LidPidTuningNotOffError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.LidPIDTuningStopped})
	})
}

func shakerHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]
		var err error

		var shObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(shObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		if deps.PlcDeck[deck].IsRunInProgress() || deps.PlcDeck[deck].IsShakerInProgress() {
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		go deps.PlcDeck[deck].Shaking(shObj, true)

		logger.Infoln(responses.ShakingSuccess)
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: responses.ShakingSuccess})
	})
}

func heaterHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]
		var err error

		var hObj db.Heating
		err = json.NewDecoder(req.Body).Decode(&hObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.HeatingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(hObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.HeatingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		if deps.PlcDeck[deck].IsRunInProgress() || deps.PlcDeck[deck].IsHeaterInProgress() {
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		go deps.PlcDeck[deck].Heating(hObj, true)

		logger.Infoln(responses.HeatingSuccess)
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: responses.HeatingSuccess})
	})
}

func dyeToleranceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var dyeWellTolerance db.DyeWellTolerance
		err := json.NewDecoder(req.Body).Decode(&dyeWellTolerance)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DyeToleranceDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DyeToleranceDecodeError.Error()})
			return
		}

		go deps.Plc.HomingRTPCR()
		go toleranceCalulation(deps, dyeWellTolerance)
		logger.Infoln(responses.DyeToleranceProgressSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, MsgObj{Msg: "dye tolerance calculation in progress"})

	})
}

func toleranceCalulation(deps Dependencies, dwtol db.DyeWellTolerance) {
	cycleCount := config.GetOpticalCalibrationCycleCount()
	//validate the kit id
	plc.ExperimentRunning = true
	dye, err := deps.Store.ShowDye(context.Background(), dwtol.DyeID)
	defer func() {
		if err != nil {
			deps.WsErrCh <- err
			return
		}
	}()
	knownValue, valid := validateandGetKitID(dwtol.KitID, dye.Position)
	if !valid {
		err = responses.InvalidKitIDError
		return
	}

	opticalResult, err := deps.Plc.CalculateOpticalResult(dye, dwtol.KitID, knownValue, cycleCount)
	if err != nil {
		return
	}
	err = deps.Store.UpsertDyeWellTolerance(context.Background(), opticalResult)
	if err != nil {
		return
	}

	deps.WsMsgCh <- "PROGRESS_OPTCALIB_" + fmt.Sprintf("%d", 100)
	deps.WsMsgCh <- fmt.Sprintf("SUCCESS_OPTCALIB_successfully caliberated rt-pcr for dye %s", dye.Name)

}

func validateandGetKitID(kitID string, dyePos int) (knownValue int64, valid bool) {

	if len(kitID) != 8 {
		logger.Errorln("length of kit id is invalid")
		return
	}

	year, err := strconv.ParseInt(kitID[0:2], 10, 64)
	if err != nil {
		logger.Errorln("year of kit id is invalid")
		return
	}
	// we need the last two digit hence sustracting 2000
	prevInvalidYear := time.Now().Year() - 3 - 2000
	if year <= int64(prevInvalidYear) {
		logger.Errorln("year of kit id is outdated")

	}
	knownValue, err = strconv.ParseInt(kitID[2:6], 10, 64)
	if err != nil {
		logger.Errorln("known value of kit id is invalid")
		return
	}

	month, err := strconv.ParseInt(kitID[6:7], 10, 64)
	if err != nil {
		monthStr := strings.ToUpper(kitID[6:7])
		if !([]rune(monthStr)[0] >= 'A' && []rune(monthStr)[0] <= 'C') {
			logger.Errorln("month of kit id is invalid")
			return
		}
	}
	if !(month > 1 && month < 9) {
		logger.Errorln("month of kit id is invalid")
		return
	}
	dyePosition, err := strconv.ParseInt(kitID[7:], 10, 64)
	if dyePosition != int64(dyePos) || err != nil {
		logger.Errorln("dye position of kit id is invalid")
		return
	}
	valid = true
	return
}
