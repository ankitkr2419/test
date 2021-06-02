package db

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"mylab/cpagent/responses"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

const (
	csv_version = "1.3.0"
	version     = "VERSION"
	position    = "POSITION"
	recipe      = "RECIPE"
	dummy       = "DUMMY"
	blank       = ""
)

var sequenceNumber int64 = 0
var createdRecipe Recipe
var csvCtx context.Context = context.WithValue(context.Background(), ContextKeyUsername, "main")

// done will help us clean up
var done bool

func ImportCSV(csvPath string) (err error) {

	var store Storer
	store, err = Init()
	if err != nil {
		logger.Errorln("err", err.Error())
		return
	}

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
	for _, rd := range recipeDetails[:2] {
		if rd == blank {
			return responses.BlankDetailsError
		}
	}
	createdRecipe.Name = recipeDetails[0]
	createdRecipe.Description = recipeDetails[1]
	return nil
}

func createRecipe(record []string, store Storer) (err error) {

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

	if createdRecipe.TotalTime, err = CalculateTimeInSeconds(record[12]); err != nil {
		logger.Warnln(err, record[12])
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
// But this avoids populating Storer interface with CSV Methods
func createProcesses(record []string, store Storer) (err error) {

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
func createAspireDispenseProcess(record []string, store Storer) (err error) {
	logger.Info("Inside aspire dispense create Process. Record: ", record)

	//  record[0] is Category
	if len(record[0]) != 2 {
		err = fmt.Errorf("Category is supposed to be only 2 characters. Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	a := AspireDispense{}
	switch {
	case strings.EqualFold(strings.TrimSpace(record[0]), "WS"):
		a.Category = WS
	case strings.EqualFold(strings.TrimSpace(record[0]), "SW"):
		a.Category = SW
	case strings.EqualFold(strings.TrimSpace(record[0]), "WW"):
		a.Category = WW
	case strings.EqualFold(strings.TrimSpace(record[0]), "WD"):
		a.Category = WD
	case strings.EqualFold(strings.TrimSpace(record[0]), "DW"):
		a.Category = DW
	case strings.EqualFold(strings.TrimSpace(record[0]), "DD"):
		a.Category = DD
	case strings.EqualFold(strings.TrimSpace(record[0]), "SD"):
		a.Category = SD
	case strings.EqualFold(strings.TrimSpace(record[0]), "DS"):
		a.Category = DS
	default:
		err = fmt.Errorf("Category is supposed to be only from these [WW, WS,SW,DD,DS,SD,DW,WD].Current Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	// record[1] is CartridgeType
	switch record[1] {
	case "1":
		a.CartridgeType = Cartridge1
	case "2":
		a.CartridgeType = Cartridge2
	default:
		err = fmt.Errorf("CartridgeType is supposed to be only from these [1,2]. Avoid any spaces. Current Category: %v. Setting Cartridge Type to 1", record[1])
		a.CartridgeType = Cartridge1
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

	createdProcess, err := store.CreateAspireDispense(csvCtx, a, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AspireDispense Record Inserted->", createdProcess)

	return nil
}

func createAttachDetachProcess(record []string, store Storer) (err error) {
	logger.Info("Inside attach detach create Process. Record: ", record)
	a := AttachDetach{
		Operation: record[0],
		// TODO: Remove this hardcoding in future when magnet_operation_subtype will be used
		OperationType: "lysis",
	}

	createdProcess, err := store.CreateAttachDetach(csvCtx, a, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AttachDetach Record Inserted->", createdProcess)

	return nil
}

func createDelayProcess(record []string, store Storer) (err error) {
	logger.Info("Inside delay create Process. Record: ", record)

	d := Delay{}
	if delay, err := strconv.ParseInt(record[0], 10, 64); err != nil {
		logger.Errorln(err, record[0])
		return err
	} else {
		d.DelayTime = delay
	}

	createdProcess, err := store.CreateDelay(csvCtx, d, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Delay Record Inserted->", createdProcess)
	return nil
}

func createPiercingProcess(record []string, store Storer) (err error) {
	logger.Info("Inside piercing create Process. Record: ", record)

	p := Piercing{}

	// record[0] is CartridgeType
	switch record[0] {
	case "1":
		p.Type = Cartridge1
	case "2":
		p.Type = Cartridge2
	default:
		err = fmt.Errorf("CartridgeType is supposed to be only from these [1,2]. Avoid any spaces. Current Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	wells := strings.Split(record[1], ",")

	for _, well := range wells {
		if wellInteger, err := strconv.ParseInt(well, 10, 64); err != nil {
			logger.Errorln(err, well)
			return err
		} else {
			p.CartridgeWells = append(p.CartridgeWells, wellInteger)
		}
	}

	logger.Debugln("After Trimming wells-> ", record[1], ".After splitting->", wells, ".Integer Wells-> ", p.CartridgeWells)

	createdProcess, err := store.CreatePiercing(csvCtx, p, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Piercing Record Inserted->", createdProcess)

	return nil
}

// TODO: Implement Discard at_pickup_passing for tip operation whenever feature added
func createTipOperationProcess(record []string, store Storer) (err error) {
	logger.Info("Inside tip operation create Process. Record: ", record)

	t := TipOperation{}

	t.Type = TipOps(record[0])
	if t.Type == PickupTip {
		if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
			logger.Errorln(err, record[1])
			return err
		}
	} else if t.Type != DiscardTip {
		err = responses.TipOperationTypeInvalid
		return err
	}

	createdProcess, err := store.CreateTipOperation(csvCtx, t, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Operation Record Inserted->", createdProcess)

	return nil
}

func createTipDockingProcess(record []string, store Storer) (err error) {
	logger.Info("Inside tip docking create Process. Record: ", record)

	t := TipDock{}

	t.Type = record[0]
	if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
		logger.Errorln(err, record[1])
		return err
	}

	createdProcess, err := store.CreateTipDocking(csvCtx, t, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Docking Record Inserted->", createdProcess)

	return nil
}

func createShakingProcess(record []string, store Storer) (err error) {
	logger.Info("Inside shaking create Process. Record: ", record)

	s := Shaker{}

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

	createdProcess, err := store.CreateShaking(csvCtx, s, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Shaking Record Inserted->", createdProcess)

	return nil
}

func createHeatingProcess(record []string, store Storer) (err error) {
	logger.Info("Inside heating create Process. Record: ", record)

	h := Heating{}

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

	createdProcess, err := store.CreateHeating(csvCtx, h, createdRecipe.ID)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Heating Record Inserted->", createdProcess)

	return nil
}

func clearFailedRecipe(store Storer) {
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
