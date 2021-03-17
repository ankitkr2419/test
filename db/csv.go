package db

import(
	// "database/sql"
	"os"
	// "github.com/jmoiron/sqlx"
	"context"
	"fmt"
	"encoding/csv"
	"io"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"

)

var sequenceNumber int64 = 0
var createdRecipe Recipe
// done will help us clean up
var done bool 

func ImportCSV(recipeName, csvPath string) (err error){

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
		Name: recipeName,
		Description: "Covid Recipe",
		Position1: 1,
		Position2: 2,
		Position3: 3,
		Position4: 4,
		Position5: 5,
		Cartridge1Position: 1,
		Position7: 6,
		Cartridge2Position: 2,
		Position9: 7,
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
			if err != nil{
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
func createProcesses(record []string, store Storer) (err error){
	sequenceNumber++

	p := Process{
		Name: "Process",
		Type: record[0],
		SequenceNumber: sequenceNumber,
		RecipeID: createdRecipe.ID,
	}
	// Insert into processes, note created processID
	createdProcess, err := store.CreateProcess(context.Background(), p)
	if err != nil {
		logger.Errorln("Error creating Process: ", p)
		return
	}
	// Create database entry for individual process here 
	// based on the name in record[0]
	switch record[0]{
	case "AspireDispense":
		createAspireDispenseProcess(record[1:], createdProcess.ID, store)
	case "AttachDetach":
		createAttachDetachProcess(record[1:], createdProcess.ID, store)
	case "Delay":
		createDelayProcess(record[1:], createdProcess.ID, store)
	case "Piercing":
		createPiercingProcess(record[1:], createdProcess.ID, store)
	case "TipOperation":
		createTipOperationProcess(record[1:], createdProcess.ID, store)
	case "TipDocking":
		createTipDockingProcess(record[1:], createdProcess.ID, store)
	case "Shaking":
		createShakingProcess(record[1:], createdProcess.ID, store)
	case "Heating":
		createHeatingProcess(record[1:], createdProcess.ID, store)
	default:
		err = fmt.Errorf("unknown process found in csv!: %v ", record[0])
		logger.Errorln(err)
		return
	}
	return nil
}
	
func createAspireDispenseProcess(record []string, processID uuid.UUID, store Storer) (err error){
	 logger.Info("Inside aspire dispense create Process. Record: ", record,". ProcessID:" ,processID)
	 return nil
}

func createAttachDetachProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside attach detach create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createDelayProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside delay create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createPiercingProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside piercing create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createTipOperationProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside tip operation create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createTipDockingProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside tip docking create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createShakingProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside shaking create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}

func createHeatingProcess(record []string, processID uuid.UUID, store Storer) (err error){
	logger.Info("Inside heating create Process. Record: ", record,". ProcessID:" ,processID)
	return nil
}


func clearFailedRecipe(store Storer) {
	if !done {
		err := store.DeleteRecipe(context.Background(), createdRecipe.ID)
		if err != nil {
			logger.Warnln("Couldn't cleanUp the partial recipe with ID: ", createdRecipe.ID)
			return
		}
	}
	logger.Info("Partial recipe cleaned up")
}
