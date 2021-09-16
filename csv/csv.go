package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"mylab/cpagent/service"

	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

const (
	csv_version = "1.4.0"
	version     = "VERSION"
	csv_type    = "TYPE"
	position    = "POSITION"
	recipe      = "RECIPE"
	template    = "TEMPLATE"
	dummy       = "DUMMY"
	blank       = ""
	rtpcr       = "RTPCR"
	extraction  = "EXTRACTION"

	target = "TARGET"
)

var sequenceNumber int64 = 0
var cycleCount uint16
var createdRecipe db.Recipe
var createdTemplate db.Template
var createdStages []db.Stage
var csvCtx context.Context = context.WithValue(context.Background(), db.ContextKeyUsername, "main")
var hStage, cStage db.Stage
var step db.Step

// done will help us clean up
var done, dataCapture, cycleSeen, templateCreated, createdHoldStage, createdCycleStage bool

func ImportCSV(csvPath string) (err error) {

	var store db.Storer
	store, err = db.Init()
	if err != nil {
		logger.Errorln("err", err.Error())
		return
	}

	service.LoadAllSetups(store)

	// open csvPath file for reading
	csvfile, err := os.Open(csvPath)
	if err != nil {
		logger.Errorln("Couldn't open the csv file", err)
		return
	}

	// Parse the csv file
	csvReader := csv.NewReader(csvfile)
	csvReader.FieldsPerRecord = -1

	record, err := csvReader.Read()
	if err != nil {
		logger.Errorln("error while reading a record from csvReader:", err)
		return err
	}
	if !strings.EqualFold(record[0], version) {
		logger.Errorln("No version found for csv:", record[0])
		return err
	}

	// 1.3.0 is the only currently supported version
	if strings.TrimSpace(record[1]) != csv_version {
		err = fmt.Errorf("%v version isn't currently supported for csv import. Please try version %v", record[1], version)
		logger.Errorln(err)
		return err
	}

	// Check for Type
	record, err = csvReader.Read()
	if !strings.EqualFold(record[0], csv_type) {
		logger.Errorln("No type found for csv:", record[0])
		return err
	}

	// rtpcr and extraction are the only currently supported types
	switch strings.TrimSpace(strings.ToUpper(record[1])) {
	case extraction:
		err = importExtraction(store, csvReader)
	case rtpcr:
		err = importRTPCR(store, csvReader)
	default:
		err = fmt.Errorf("%v version isn't currently supported for csv import. Please try version %v", record[1], version)
	}
	if err != nil {
		logger.Errorln(err)
	}
	return err
}

func importRTPCR(store db.Storer, csvReader *csv.Reader) (err error) {
	// clean up failed recipe
	defer clearFailedTemplate(store)

	// Iterate through the records

iterateCSV:
	for {
		// Read each record from csv
		record, err := csvReader.Read()
		if err == io.EOF {
			logger.Infoln("Reached end of csv file")
			break
		}
		if err != nil {
			logger.Errorln("error while reading a record from csvReader:", err)
			return err
		}

		switch strings.TrimSpace(strings.ToUpper(record[0])) {
		case dummy:
			continue
		case template:
			logger.Infoln("Record-> ", record)
			err = addTemplate(store, record[1:])
			if err != nil {
				err = fmt.Errorf("Couldn't add template details.")
				logger.Errorln(err)
				return err
			}
			templateCreated = true
			AddTargets(store, csvReader)
			AddCycles(store, csvReader)

		case blank:
			logger.Infoln("Record-> ", record)
			if len(record) < 2 || record[1] == "" {
				err = fmt.Errorf("record has unexpected length or empty process name, maybe CSV is over.")
				logger.Warnln(err, record)
				break iterateCSV
			}
		default:
			logger.Infoln("Record-> ", record)
			return responses.CSVBadContentError
		}
	}

	cStage.RepeatCount = cycleCount
	err = store.UpdateStage(csvCtx, cStage)
	if err != nil {
		return
	}

	err = store.UpdateStepCount(csvCtx)
	if err != nil {
		return
	}

	done = true
	return nil
}

func AddTargets(store db.Storer, csvReader *csv.Reader) (err error) {

	subRecord, err := csvReader.Read()
	if err == io.EOF {
		logger.Infoln("Reached end of csv file")
		return
	}
	if err != nil {
		logger.Errorln("error while reading a record from csvReader:", err)
		return err
	}
	if subRecord[0] == target {

		targetDetails, err := store.GetTargetByName(csvCtx, subRecord[1])
		if err != nil {
			logger.Errorln("error while fetching target details:", err)
			return err
		}
		tempTarget := []db.TemplateTarget{db.TemplateTarget{
			TemplateID: createdTemplate.ID,
			TargetID:   targetDetails.ID,
		}}
		if threshold, err := strconv.ParseFloat(subRecord[2], 64); err != nil {
			logger.Errorln(err, subRecord[2])
			return err
		} else {
			tempTarget[0].Threshold = float64(threshold)
		}

		targetTemp, err := store.UpsertTemplateTarget(csvCtx, tempTarget, createdTemplate.ID)
		logger.Info("Created template target-> ", targetTemp)

	}
	return nil

}

func AddCycles(store db.Storer, csvReader *csv.Reader) (err error) {

	for {

		subRecord, err := csvReader.Read()
		if err == io.EOF {
			logger.Infoln("Reached end of csv file")
			break
		}
		if err != nil {
			logger.Errorln("error while reading a record from csvReader:", err)
			return err
		}
		startPoint := subRecord[1]
		logger.Infoln("sub Record-> ", subRecord)

	startPoint:
		if startPoint == "hold/cycle" {
			stageRecord, err := csvReader.Read()
			if err == io.EOF {
				logger.Infoln("Reached end of csv file")
				break
			}
			if err != nil {
				logger.Errorln("error while reading a record from csvReader:", err)
				return err
			}
			if stageRecord[0] != dummy {
				addStage(store, stageRecord)
			}
			goto startPoint
		}
	}
	return
}

func addStage(store db.Storer, record []string) (err error) {
	logger.Infoln("Record-> ", record)
	if !templateCreated {
		err = fmt.Errorf("template doesn't exist, first add its entry")
		logger.Errorln(err)
		return err
	}

	switch strings.TrimSpace(strings.ToLower(record[1])) {
	case db.Hold:
		if cycleSeen {
			err = fmt.Errorf("Couldn't create Holding step entry as Cycle entry is alreday present.")
			logger.Errorln(err)
			return err
		}
		if createdHoldStage {
			err = addHoldStep(store, record[1:])
			break
		}
		err = addHoldStage(store, record[1:])
		if err != nil {
			err = fmt.Errorf("Couldn't create Holding step entry.")
			logger.Errorln(err)
			return err
		}

	case db.Cycle:
		if createdCycleStage {
			err = addCycleStep(store, record[1:])
			break
		}
		err = addCycleStage(store, record[1:])
		if err != nil {
			err = fmt.Errorf("Couldn't create Cycling step entry.")
			logger.Errorln(err)
			return err
		}
		cycleSeen = true
	default:
		err = fmt.Errorf("unknown stage type found!")
		logger.Errorln(err)
		return err
	}
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func addTemplate(store db.Storer, templateDetails []string) (err error) {
	// template name and description are NOT allowed to be empty/blank
	for _, rd := range templateDetails[:2] {
		if rd == blank {
			return responses.BlankDetailsError
		}
	}

	db.CurrentTemp = config.GetRoomTemp()
	defer func() {
		// set current Temp to Room Temp
		db.CurrentTemp = config.GetRoomTemp()
	}()

	createdTemplate.Name = templateDetails[0]
	createdTemplate.Description = templateDetails[1]

	if temp, err := strconv.ParseInt(templateDetails[2], 10, 64); err != nil {
		logger.Errorln(err, templateDetails[2])
		return err
	} else {
		cycleCount = uint16(temp)
	}

	if dataCapture, err = strconv.ParseBool(templateDetails[3]); err != nil {
		logger.Warnln(err, templateDetails[3])
		return err
	}

	if createdTemplate.Publish, err = strconv.ParseBool(templateDetails[4]); err != nil {
		logger.Warnln(err, templateDetails[4])
		return err
	}

	if createdTemplate.LidTemp, err = strconv.ParseInt(templateDetails[5], 10, 64); err != nil {
		logger.Errorln(err, templateDetails[5])
		return err
	}

	if createdTemplate.Volume, err = strconv.ParseInt(templateDetails[6], 10, 64); err != nil {
		logger.Errorln(err, templateDetails[5])
		return err
	}

	// Create Template
	createdTemplate, err = store.CreateTemplate(csvCtx, createdTemplate)
	if err != nil {
		logger.Errorln("Couldn't insert template entry", err)
		return
	}
	logger.Info("Created template-> ", createdTemplate)

	go db.UpdateEstimatedTimeByTemplateID(csvCtx, store, createdTemplate.ID)
	return nil

}

func addHoldStage(store db.Storer, record []string) (err error) {

	hStage.Type = db.Hold
	hStage.TemplateID = createdTemplate.ID

	cStage.Type = db.Cycle
	cStage.TemplateID = createdTemplate.ID
	cStage.RepeatCount = cycleCount

	if cycleCount < db.RepeatCountDefault {
		logger.Warnln("Repeat Count for cycle stage is below threshold")
	}

	// Create both Stages
	createdStages, err = store.CreateStages(csvCtx, []db.Stage{hStage, cStage})
	if err != nil {
		logger.Errorln("Couldn't insert Stage entries", err)
		return
	}

	for _, st := range createdStages {
		if st.Type == db.Hold {
			hStage = st
		} else {
			cStage = st
		}
	}

	createdHoldStage = true
	logger.Info("Created Stages-> ", createdStages)

	err = addHoldStep(store, record)
	if err != nil {
		return
	}

	return nil
}

func addHoldStep(store db.Storer, record []string) (err error) {
	logger.Infoln("hold step record----------", record)
	step.StageID = hStage.ID
	step.DataCapture = dataCapture

	if temp, err := strconv.ParseFloat(record[1], 64); err != nil {
		logger.Errorln(err, record[1])
		return err
	} else {
		step.TargetTemperature = float32(temp)
	}

	if temp, err := strconv.ParseFloat(record[2], 64); err != nil {
		logger.Errorln(err, record[4])
		return err
	} else {
		step.RampRate = float32(temp)
	}

	if temp, err := strconv.ParseInt(record[3], 10, 32); err != nil {
		logger.Errorln(err, record[2])
		return err
	} else {
		step.HoldTime = int32(temp)
	}

	createdStep, err := store.CreateStep(csvCtx, step)
	if err != nil {
		logger.Errorln("Couldn't insert Step entry", err)
		return err
	}

	logger.Infoln("Step Created for Hold: ", createdStep)
	return nil
}

func addCycleStep(store db.Storer, record []string) (err error) {

	step.StageID = cStage.ID
	step.DataCapture = dataCapture

	if temp, err := strconv.ParseFloat(record[1], 64); err != nil {
		logger.Errorln(err, record[1])
		return err
	} else {
		step.TargetTemperature = float32(temp)
	}

	if temp, err := strconv.ParseFloat(record[2], 64); err != nil {
		logger.Errorln(err, record[4])
		return err
	} else {
		step.RampRate = float32(temp)
	}

	if temp, err := strconv.ParseInt(record[3], 10, 32); err != nil {
		logger.Errorln(err, record[2])
		return err
	} else {
		step.HoldTime = int32(temp)
	}

	createdStep, err := store.CreateStep(csvCtx, step)
	if err != nil {
		logger.Errorln("Couldn't insert Step entry", err)
		return err
	}

	logger.Infoln("Step Created for Cycle: ", createdStep)
	return nil
}

func addCycleStage(store db.Storer, record []string) (err error) {

	if !createdHoldStage {

		cStage.Type = db.Cycle
		cStage.TemplateID = createdTemplate.ID
		cStage.RepeatCount = cycleCount

		createdStages, err = store.CreateStages(csvCtx, []db.Stage{cStage})
		if err != nil {
			logger.Errorln("Couldn't insert Cycle Stage entry", err)
			return
		}
		cStage = createdStages[0]

		createdCycleStage = true
		logger.Info("Created Stages-> ", createdStages)
	}

	err = addCycleStep(store, record)
	if err != nil {
		return
	}

	return nil
}

func importExtraction(store db.Storer, csvReader *csv.Reader) (err error) {
	// clean up failed recipe
	defer clearFailedRecipe(store)

	// Iterate through the records

iterateCSV:
	for {
		// Read each record from csv
		record, err := csvReader.Read()
		if err == io.EOF {
			logger.Infoln("Reached end of csv file")
			break
		}
		if err != nil {
			logger.Errorln("error while reading a record from csvReader:", err)
			return err
		}

		switch strings.TrimSpace(strings.ToUpper(record[0])) {
		case dummy:
			continue
		case recipe:
			logger.Infoln("Record-> ", record)
			err = addRecipeDetails(record[1:])
			if err != nil {
				err = fmt.Errorf("Couldn't add recipe details.")
				logger.Errorln(err)
				return err
			}
		case position:
			logger.Infoln("Record-> ", record)
			err = createRecipe(record, store)
			if err != nil {
				err = fmt.Errorf("Couldn't create recipe entry.")
				logger.Errorln(err)
				return err
			}
		case blank:
			logger.Infoln("Record-> ", record)
			if len(record) < 2 || record[1] == "" {
				err = fmt.Errorf("record has unexpected length or empty process name, maybe CSV is over.")
				logger.Warnln(err, record)
				break iterateCSV
			} else {
				err = createProcesses(record[1:], store)
				if err != nil {
					err = fmt.Errorf("Couldn't create process entry.")
					logger.Errorln(err)
					return err
				}
			}
		default:
			logger.Infoln("Record-> ", record)
			return responses.CSVBadContentError
		}
	}

	done = true
	return nil
}

func addRecipeDetails(recipeDetails []string) (err error) {
	// recipe time and is_published are allowed to be empty/blank
	for _, rd := range recipeDetails[:2] {
		if rd == blank {
			return responses.BlankDetailsError
		}
	}

	createdRecipe.Name = recipeDetails[0]
	createdRecipe.Description = recipeDetails[1]

	if createdRecipe.TotalTime, err = db.CalculateTimeInSeconds(recipeDetails[2]); err != nil {
		logger.Warnln(err, recipeDetails[2])
	}

	if createdRecipe.IsPublished, err = strconv.ParseBool(recipeDetails[3]); err != nil {
		logger.Warnln(err, recipeDetails[3])
		return err
	}

	return nil
}

func createRecipe(record []string, store db.Storer) (err error) {

	for i, rec := range record {
		record[i] = strings.TrimSpace(rec)
	}

	// extra record just to make creation easy
	var positions [12]int64

	// NOTE: Error during parsing at here means ignore that cell
	if positions[1], err = strconv.ParseInt(record[1], 10, 64); err != nil {
		logger.Warnln(err, record[1])
	} else {
		createdRecipe.Position1 = &positions[1]
	}

	if positions[2], err = strconv.ParseInt(record[2], 10, 64); err != nil {
		logger.Warnln(err, record[2])
	} else {
		createdRecipe.Position2 = &positions[2]
	}

	if positions[3], err = strconv.ParseInt(record[3], 10, 64); err != nil {
		logger.Warnln(err, record[3])
	} else {
		createdRecipe.Position3 = &positions[3]
	}

	if positions[4], err = strconv.ParseInt(record[4], 10, 64); err != nil {
		logger.Warnln(err, record[4])
	} else {
		createdRecipe.Position4 = &positions[4]
	}

	if positions[5], err = strconv.ParseInt(record[5], 10, 64); err != nil {
		logger.Warnln(err, record[5])
	} else {
		createdRecipe.Position5 = &positions[5]
	}

	if positions[6], err = strconv.ParseInt(record[6], 10, 64); err != nil {
		logger.Warnln(err, record[6])
	} else {
		createdRecipe.Position6 = &positions[6]
	}

	if positions[7], err = strconv.ParseInt(record[7], 10, 64); err != nil {
		logger.Warnln(err, record[7])
	} else {
		createdRecipe.Position7 = &positions[7]
	}

	if positions[8], err = strconv.ParseInt(record[8], 10, 64); err != nil {
		logger.Warnln(err, record[8])
	} else {
		createdRecipe.Cartridge1Position = &positions[8]
	}

	if positions[9], err = strconv.ParseInt(record[9], 10, 64); err != nil {
		logger.Warnln(err, record[9])
	} else {
		createdRecipe.Position9 = &positions[9]
	}

	if positions[10], err = strconv.ParseInt(record[10], 10, 64); err != nil {
		logger.Warnln(err, record[10])
	} else {
		createdRecipe.Cartridge2Position = &positions[10]
	}

	if positions[11], err = strconv.ParseInt(record[11], 10, 64); err != nil {
		logger.Warnln(err, record[11])
	} else {
		createdRecipe.Position11 = &positions[11]
	}

	logger.Infoln("Record that will be created--> ", createdRecipe)

	// Create Recipe
	createdRecipe, err = store.CreateRecipe(csvCtx, createdRecipe)
	if err != nil {
		logger.Errorln("Couldn't insert recipe entry", err)
		return
	}
	logger.Info("Created Recipe-> ", createdRecipe)
	return nil
}

// NOTE: Passing db connection as function parameter isn't the best approach
// But this avoids populating db.Storer interface with CSV Methods
func createProcesses(record []string, store db.Storer) (err error) {

	// Create database entry for individual process here
	// based on the name in record[0]
	switch strings.TrimSpace(record[0]) {
	case "AspireDispense":
		err = createAspireDispenseProcess(record[1:], store)
	case "AttachDetach":
		err = createAttachDetachProcess(record[1:], store)
	case "Delay":
		err = createDelayProcess(record[1:], store)
	case "Piercing":
		err = createPiercingProcess(record[1:], store)
	case "TipOperation":
		err = createTipOperationProcess(record[1:], store)
	case "TipDocking":
		err = createTipDockingProcess(record[1:], store)
	case "Shaking":
		err = createShakingProcess(record[1:], store)
	case "Heating":
		err = createHeatingProcess(record[1:], store)
	default:
		err = fmt.Errorf("unknown process found in csv!: %v ", record[0])
	}
	if err != nil {
		logger.Errorln(err)
		return
	}
	return nil
}

// WARN: DB changes will also need to be reflected in below functions!
func createAspireDispenseProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside aspire dispense create Process. Record: ", record)

	//  record[0] is Category
	if len(record[0]) != 2 {
		err = fmt.Errorf("Category is supposed to be only 2 characters. Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	a := db.AspireDispense{}
	switch {
	case strings.EqualFold(strings.TrimSpace(record[0]), "WS"):
		a.Category = db.WS
	case strings.EqualFold(strings.TrimSpace(record[0]), "SW"):
		a.Category = db.SW
	case strings.EqualFold(strings.TrimSpace(record[0]), "WW"):
		a.Category = db.WW
	case strings.EqualFold(strings.TrimSpace(record[0]), "WD"):
		a.Category = db.WD
	case strings.EqualFold(strings.TrimSpace(record[0]), "DW"):
		a.Category = db.DW
	case strings.EqualFold(strings.TrimSpace(record[0]), "DD"):
		a.Category = db.DD
	case strings.EqualFold(strings.TrimSpace(record[0]), "SD"):
		a.Category = db.SD
	case strings.EqualFold(strings.TrimSpace(record[0]), "DS"):
		a.Category = db.DS
	default:
		err = fmt.Errorf("Category is supposed to be only from these [WW, WS,SW,DD,DS,SD,DW,WD].Current Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	// record[1] is CartridgeType
	switch record[1] {
	case "1":
		a.CartridgeType = db.Cartridge1
	case "2":
		a.CartridgeType = db.Cartridge2
	default:
		err = fmt.Errorf("CartridgeType is supposed to be only from these [1,2]. Avoid any spaces. Current Category: %v. Setting Cartridge Type to 1", record[1])
		a.CartridgeType = db.Cartridge1
		logger.Warnln(err)
	}

	if a.SourcePosition, err = strconv.ParseInt(record[2], 10, 64); err != nil {
		logger.Errorln(err, record[2])
		return
	}
	if a.AspireHeight, err = strconv.ParseFloat(record[3], 64); err != nil {
		logger.Errorln(err, record[3])
		return
	}
	if a.AspireMixingVolume, err = strconv.ParseFloat(record[4], 64); err != nil {
		logger.Errorln(err, record[4])
		return
	}
	if a.AspireNoOfCycles, err = strconv.ParseInt(record[5], 10, 64); err != nil {
		logger.Errorln(err, record[5])
		return
	}
	if a.AspireVolume, err = strconv.ParseFloat(record[6], 64); err != nil {
		logger.Errorln(err, record[6])
		return
	}
	if a.AspireAirVolume, err = strconv.ParseFloat(record[7], 64); err != nil {
		logger.Errorln(err, record[7])
		return
	}
	if a.DispenseHeight, err = strconv.ParseFloat(record[8], 64); err != nil {
		logger.Errorln(err, record[8])
		return
	}
	if a.DispenseMixingVolume, err = strconv.ParseFloat(record[9], 64); err != nil {
		logger.Errorln(err, record[9])
		return
	}
	if a.DispenseNoOfCycles, err = strconv.ParseInt(record[10], 10, 64); err != nil {
		logger.Errorln(err, record[10])
		return
	}

	// NOTE: Since version 1.2.1 we have deprecated CSV support for
	// dispense volume and dispense air volume
	if a.DestinationPosition, err = strconv.ParseInt(record[11], 10, 64); err != nil {
		logger.Errorln(err, record[11])
		return
	}
	valid, _ := service.Validate(a)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.AspireDispenseValidationError)
		return
	}
	err = service.ValidateAspireDispenceObject(csvCtx, service.Dependencies{
		Store: store,
	}, a, createdRecipe.ID)
	if err != nil {
		logger.WithField("ERROR", "ASPIRE DISPENSE").Errorln(err)
		return
	}

	createdProcess, err := store.CreateAspireDispense(csvCtx, a, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AspireDispense Record Inserted->", createdProcess)

	return nil
}

func createAttachDetachProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside attach detach create Process. Record: ", record)

	var height int64

	if strings.EqualFold(record[0], "attach") {
		if height, err = strconv.ParseInt(record[1], 10, 64); err != nil {
			logger.Errorln(err, record[1])
			return err
		}
	}

	a := db.AttachDetach{
		Operation: record[0],
		Height:    height,
	}

	valid, _ := service.Validate(a)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.AttachDetachValidationError)
		return
	}

	createdProcess, err := store.CreateAttachDetach(csvCtx, a, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AttachDetach Record Inserted->", createdProcess)

	return nil
}

func createDelayProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside delay create Process. Record: ", record)

	d := db.Delay{}
	if delay, err := strconv.ParseInt(record[0], 10, 64); err != nil {
		logger.Errorln(err, record[0])
		return err
	} else {
		d.DelayTime = delay
	}
	valid, _ := service.Validate(d)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.DelayValidationError)
		return
	}

	createdProcess, err := store.CreateDelay(csvCtx, d, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Delay Record Inserted->", createdProcess)
	return nil
}

func createPiercingProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside piercing create Process. Record: ", record)

	p := db.Piercing{}

	// record[0] is CartridgeType
	switch record[0] {
	case "1":
		p.Type = db.Cartridge1
	case "2":
		p.Type = db.Cartridge2
	default:
		err = fmt.Errorf("CartridgeType is supposed to be only from these [1,2]. Avoid any spaces. Current Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	wells := strings.Split(record[1], ",")
	heights := strings.Split(record[2], ",")

	for _, well := range wells {
		if wellInteger, err := strconv.ParseInt(well, 10, 64); err != nil {
			logger.Errorln(err, well)
			return err
		} else {
			p.CartridgeWells = append(p.CartridgeWells, wellInteger)
		}
	}

	for _, height := range heights {
		if heightInteger, err := strconv.ParseInt(height, 10, 64); err != nil {
			logger.Errorln(err, height)
			return err
		} else {
			p.Heights = append(p.Heights, heightInteger)
		}
	}

	logger.Debugln("After Trimming wells-> ", record[1], ".After splitting->", wells, ".Integer Wells-> ", p.CartridgeWells)
	logger.Debugln("After Trimming heights-> ", record[1], ".After splitting->", heights, ".Integer heights-> ", p.Heights)

	valid, _ := service.Validate(p)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.PiercingValidationError)
		return
	}
	err = service.ValidatePiercingObject(csvCtx, service.Dependencies{
		Store: store,
	}, &p, createdRecipe.ID)
	if err != nil {
		logger.WithField("ERROR", "PIERCING").Errorln(err)
		return
	}

	createdProcess, err := store.CreatePiercing(csvCtx, p, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Piercing Record Inserted->", createdProcess)

	return nil
}

// TODO: Implement Discard at_pickup_passing for tip operation whenever feature added
func createTipOperationProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside tip operation create Process. Record: ", record)

	t := db.TipOperation{}

	t.Type = db.TipOps(record[0])
	if t.Type == db.PickupTip {
		if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
			logger.Errorln(err, record[1])
			return err
		}
		err = service.ValidateTipPickupObject(csvCtx, service.Dependencies{
			Store: store,
		}, t, createdRecipe.ID)
		if err != nil {
			logger.WithField("ERROR", "TIP OPERATION").Errorln(err)
			return
		}
	} else if t.Type != db.DiscardTip {
		err = responses.TipOperationTypeInvalid
		return err
	}
	valid, _ := service.Validate(t)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.TipOperationValidationError)
		return
	}

	createdProcess, err := store.CreateTipOperation(csvCtx, t, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Operation Record Inserted->", createdProcess)

	return nil
}

func createTipDockingProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside tip docking create Process. Record: ", record)

	t := db.TipDock{}

	t.Type = record[0]
	if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
		logger.Errorln(err, record[1])
		return err
	}
	valid, _ := service.Validate(t)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.TipDockingValidationError)
		return
	}
	err = service.ValidateTipDockObject(csvCtx, service.Dependencies{
		Store: store,
	}, t, createdRecipe.ID)
	if err != nil {
		logger.WithField("ERROR", "ASPIRE DISPENSE").Errorln(err)
		return
	}
	createdProcess, err := store.CreateTipDocking(csvCtx, t, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Docking Record Inserted->", createdProcess)

	return nil
}

func createShakingProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside shaking create Process. Record: ", record)

	s := db.Shaker{}

	if s.WithTemp, err = strconv.ParseBool(record[0]); err != nil {
		logger.Errorln(err, record[0])
		return err
	}

	if s.FollowTemp, err = strconv.ParseBool(record[1]); err != nil {
		logger.Errorln(err, record[1])
		return err
	}

	// Current Temperature is accurate only to 1 decimal point
	// While sending it to PLC  we need to multiply by 10
	// As PLC can't handle decimals
	if s.Temperature, err = strconv.ParseFloat(record[2], 64); err != nil {
		logger.Errorln(err, record[2])
		return err
	}

	if s.RPM1, err = strconv.ParseInt(record[3], 10, 64); err != nil {
		logger.Errorln(err, record[3])
		return err
	}

	if time1, err := strconv.ParseInt(record[4], 10, 64); err != nil {
		logger.Errorln(err, record[4])
		return err
	} else {
		s.Time1 = time1
	}

	if s.RPM2, err = strconv.ParseInt(record[5], 10, 64); err != nil {
		logger.Errorln(err, record[5])
		return err
	}

	if time2, err := strconv.ParseInt(record[6], 10, 64); err != nil {
		logger.Errorln(err, record[6])
		return err
	} else {
		s.Time2 = time2
	}

	valid, _ := service.Validate(s)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.ShakingValidationError)
		return
	}

	createdProcess, err := store.CreateShaking(csvCtx, s, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Shaking Record Inserted->", createdProcess)

	return nil
}

func createHeatingProcess(record []string, store db.Storer) (err error) {
	logger.Info("Inside heating create Process. Record: ", record)

	h := db.Heating{}

	// Current Temperature is accurate only to 1 decimal point
	// While sending it to PLC  we need to multiply by 10
	// As PLC can't handle decimals
	if h.Temperature, err = strconv.ParseFloat(record[0], 64); err != nil {
		logger.Errorln(err, record[0])
		return err
	}

	if h.FollowTemp, err = strconv.ParseBool(record[1]); err != nil {
		logger.Errorln(err, record[1])
		return err
	}

	if time1, err := strconv.ParseInt(record[2], 10, 64); err != nil {
		logger.Errorln(err, record[4])
		return err
	} else {
		h.Duration = time1
	}
	valid, _ := service.Validate(h)
	if !valid {
		logger.WithField("err", "Validation Error").Errorln(responses.HeatingValidationError)
		return
	}

	createdProcess, err := store.CreateHeating(csvCtx, h, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Heating Record Inserted->", createdProcess)

	return nil
}

func clearFailedRecipe(store db.Storer) {
	if !done {
		err := store.DeleteRecipe(csvCtx, createdRecipe.ID)
		if err != nil {
			logger.Warnln("Couldn't cleanUp the partial recipe with ID: ", createdRecipe.ID)
			return
		}
		logger.Info("Partial recipe cleaned up")
		return
	}
	logger.Info("complete recipe inserted successfully")
	return
}

func clearFailedTemplate(store db.Storer) {
	if !done {
		err := store.DeleteTemplate(csvCtx, createdTemplate.ID)
		if err != nil {
			logger.Warnln("Couldn't cleanUp the partial template with ID: ", createdTemplate.ID)
			return
		}
		logger.Info("Partial template cleaned up")
		return
	}
	logger.Info("complete template inserted successfully")
	return
}
