package db

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

var sequenceNumber int64 = 0
var createdRecipe Recipe

// done will help us clean up
var done bool

func ImportCSV(recipeName, csvPath string) (err error) {

	// Create DB conn
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

	// Add the recipe entry into the database for the given recipe name here
	r := Recipe{
		Name:               recipeName,
		Description:        "Covid Recipe",
		Position1:          1,
		Position2:          2,
		Position3:          3,
		Position4:          4,
		Position5:          5,
		Cartridge1Position: 1,
		Position7:          6,
		Cartridge2Position: 2,
		Position9:          7,
	}

	// Create Recipe
	createdRecipe, err = store.CreateRecipe(context.Background(), r)
	if err != nil {
		logger.Errorln("Couldn't insert recipe entry", err)
		return
	}
	logger.Info("Created Recipe-> ", createdRecipe)
	defer clearFailedRecipe(store)

	// Parse the csv file
	csvReader := csv.NewReader(csvfile)
	csvReader.FieldsPerRecord = -1
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
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

		if record[0] != "DUMMY" {
			if len(record) < 2 || record[1] == "" {
				err = fmt.Errorf("record has unexpected length or empty process name, maybe CSV is over.")
				logger.Warnln(err, record)
				break
			}
			logger.Infoln("Record-> ", record)
			err = createProcesses(record[1:], store)
			if err != nil {
				err = fmt.Errorf("Couldn't insert process entry.")
				logger.Errorln(err)
				return err
			}

		}
	}

	done = true
	return nil
}

// NOTE: Passing db connection as function parameter isn't the best approach
// But this avoids populating Storer interface with CSV Methods
func createProcesses(record []string, store Storer) (err error) {
	sequenceNumber++

	p := Process{
		Name:           "Process",
		Type:           record[0],
		SequenceNumber: sequenceNumber,
		RecipeID:       createdRecipe.ID,
	}
	// Insert into processes, note created processID
	createdProcess, err := store.CreateProcess(context.Background(), p)
	if err != nil {
		logger.Errorln("Error creating Process: ", p)
		return
	}
	// Create database entry for individual process here
	// based on the name in record[0]
	switch record[0] {
	case "AspireDispense":
		err = createAspireDispenseProcess(record[1:], createdProcess.ID, store)
	case "AttachDetach":
		err = createAttachDetachProcess(record[1:], createdProcess.ID, store)
	case "Delay":
		err = createDelayProcess(record[1:], createdProcess.ID, store)
	case "Piercing":
		err = createPiercingProcess(record[1:], createdProcess.ID, store)
	case "TipOperation":
		err = createTipOperationProcess(record[1:], createdProcess.ID, store)
	case "TipDocking":
		err = createTipDockingProcess(record[1:], createdProcess.ID, store)
	case "Shaking":
		err = createShakingProcess(record[1:], createdProcess.ID, store)
	case "Heating":
		err = createHeatingProcess(record[1:], createdProcess.ID, store)
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
func createAspireDispenseProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside aspire dispense create Process. Record: ", record, ". ProcessID:", processID)

	//  record[0] is Category
	if len(record[0]) != 2 {
		err = fmt.Errorf("Category is supposed to be only 2 characters. Category: %v", record[0])
		logger.Errorln(err)
		return
	}

	a := AspireDispense{
		ProcessID: processID,
	}

	switch {
	case strings.EqualFold(record[0], "WS"):
		a.Category = WS
	case strings.EqualFold(record[0], "SW"):
		a.Category = SW
	case strings.EqualFold(record[0], "WW"):
		a.Category = WW
	case strings.EqualFold(record[0], "WD"):
		a.Category = WD
	case strings.EqualFold(record[0], "DW"):
		a.Category = DW
	case strings.EqualFold(record[0], "DD"):
		a.Category = DD
	case strings.EqualFold(record[0], "SD"):
		a.Category = SD
	case strings.EqualFold(record[0], "DS"):
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
		err = fmt.Errorf("CartridgeType is supposed to be only from these [1,2]. Avoid any spaces. Current Category: %v", record[1])
		logger.Errorln(err)
		return
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
	if a.DispenseVolume, err = strconv.ParseFloat(record[11], 64); err != nil {
		logger.Errorln(err, record[11])
		return
	}
	if a.DispenseBlowVolume, err = strconv.ParseFloat(record[12], 64); err != nil {
		logger.Errorln(err, record[12])
		return
	}
	if a.DestinationPosition, err = strconv.ParseInt(record[13], 10, 64); err != nil {
		logger.Errorln(err, record[13])
		return
	}

	createdProcess, err := store.CreateAspireDispense(context.Background(), a)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AspireDispense Record Inserted->", createdProcess)

	return nil
}

func createAttachDetachProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside attach detach create Process. Record: ", record, ". ProcessID:", processID)
	a := AttachDetach{
		Operation: record[0],
		ProcessID: processID,
	}

	createdProcess, err := store.CreateAttachDetach(context.Background(), a)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("AttachDetach Record Inserted->", createdProcess)

	return nil
}

func createDelayProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside delay create Process. Record: ", record, ". ProcessID:", processID)

	d := Delay{
		ProcessID: processID,
	}
	if delay, err := strconv.ParseInt(record[0], 10, 64); err != nil {
		logger.Errorln(err, record[0])
		return err
	} else {
		d.DelayTime = time.Duration(delay)
	}

	createdProcess, err := store.CreateDelay(context.Background(), d)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Delay Record Inserted->", createdProcess)
	return nil
}

func createPiercingProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside piercing create Process. Record: ", record, ". ProcessID:", processID)

	p := Piercing{
		ProcessID: processID,
		Discard:   at_discard_box,
	}

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

	createdProcess, err := store.CreatePiercing(context.Background(), p)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Piercing Record Inserted->", createdProcess)

	return nil
}

func createTipOperationProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside tip operation create Process. Record: ", record, ". ProcessID:", processID)

	t := TipOperation{
		ProcessID: processID,
	}

	t.Type = TipOps(record[0])
	if t.Type == PickupTip {
		if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
			logger.Errorln(err, record[1])
			return err
		}
	}

	createdProcess, err := store.CreateTipOperation(context.Background(), t)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Operation Record Inserted->", createdProcess)

	return nil
}

func createTipDockingProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside tip docking create Process. Record: ", record, ". ProcessID:", processID)

	t := TipDock{
		ProcessID: processID,
	}

	t.Type = record[0]
	if t.Position, err = strconv.ParseInt(record[1], 10, 64); err != nil {
		logger.Errorln(err, record[1])
		return err
	}

	createdProcess, err := store.CreateTipDocking(context.Background(), t)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Tip Docking Record Inserted->", createdProcess)

	return nil
}

func createShakingProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside shaking create Process. Record: ", record, ". ProcessID:", processID)

	s := Shaking{
		ProcessID: processID,
	}

	if s.WithTemp, err = strconv.ParseBool(record[0]); err != nil {
		logger.Errorln(err, record[0])
		return err
	}

	if s.FollowTemp, err = strconv.ParseBool(record[1]); err != nil {
		logger.Errorln(err, record[1])
		return err
	}

	// Current Temperature is accurate only to 1 decimal point
	// In db we only store its multiplication by 10
	// As PLC can't handle decimals
	if temperature, err := strconv.ParseFloat(record[2], 64); err != nil {
		logger.Errorln(err, record[2])
		return err
	} else {
		s.Temperature = int64(temperature * 10)
	}

	if s.RPM1, err = strconv.ParseInt(record[3], 10, 64); err != nil {
		logger.Errorln(err, record[3])
		return err
	}

	if time1, err := strconv.ParseInt(record[4], 10, 64); err != nil {
		logger.Errorln(err, record[4])
		return err
	} else {
		s.Time1 = time.Duration(time1)
	}

	if s.RPM2, err = strconv.ParseInt(record[5], 10, 64); err != nil {
		logger.Errorln(err, record[5])
		return err
	}

	if time2, err := strconv.ParseInt(record[6], 10, 64); err != nil {
		logger.Errorln(err, record[6])
		return err
	} else {
		s.Time2 = time.Duration(time2)
	}

	createdProcess, err := store.CreateShaking(context.Background(), s)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Info("Shaking Record Inserted->", createdProcess)

	return nil
}

func createHeatingProcess(record []string, processID uuid.UUID, store Storer) (err error) {
	logger.Info("Inside heating create Process. Record: ", record, ". ProcessID:", processID)
	return nil
}

func clearFailedRecipe(store Storer) {
	if !done {
		err := store.DeleteRecipe(context.Background(), createdRecipe.ID)
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
