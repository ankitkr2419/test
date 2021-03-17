package db

import(
	// "database/sql"
	"os"
	// "github.com/jmoiron/sqlx"
	"context"
	// "fmt"
	"encoding/csv"
	"io"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"

)

var sequenceNumber int64 = 0
var createdRecipe Recipe

func ImportCSV(recipeName, csvPath string){

	// Create DB conn
	var store Storer
	store, err := Init() 
	if err != nil {
		logger.Fatalln("err", err.Error())
		return
	}
	
	// open csvPath file for reading
	csvfile, err := os.Open(csvPath)
	if err != nil {
		logger.Fatalln("Couldn't open the csv file", err)
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
		Position7: 7,
		Cartridge2Position: 2,
		Position9: 9,
	}

	// Create Recipe
	createdRecipe, err = store.CreateRecipe(context.Background(), r)
	if err != nil {
		logger.Fatalln("Couldn't insert recipe entry", err)
		return
	}
	logger.Info("Created Recipe-> ", createdRecipe)

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
			logger.Fatalln("error while reading a record from csvReader:", err)
		}
		if record[0] != "DUMMY" {
			logger.Infoln("Record-> ", record)
			err = store.createProcesses(record[1:])
			if err != nil{
				logger.Fatalln("Couldn't insert process entry", err)
				return
			}
			// Create database entry for individual process here 
			// based on the name in record[1]
		}
	}
	return
}

func (s *pgStore) createProcesses(record []string) (err error){
	sequenceNumber++
	// start transaction

	p := Process{
		Name: "Process",
		Type: record[0],
		SequenceNumber: sequenceNumber,
		RecipeID: createdRecipe.ID,
	}
	// Insert into processes, note created processID
	createdProcess, err := s.CreateProcess(context.Background(), p)
	if err != nil {
		logger.Fatalln("Couldn't insert process entry", err, p)
		return
	}
	switch record[0]{
			case "AspireDispense":
					s.createAspireDispenseProcess(record, createdProcess.ID)
			default:
			logger.Fatalln("unknown process found in csv!: ", record[0])
			return
	}
	//end transaction
	return nil
}
	
func (s *pgStore) createAspireDispenseProcess(record []string,processID uuid.UUID) (err error){
	 
	 return nil
}
