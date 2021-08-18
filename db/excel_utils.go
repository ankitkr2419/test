package db

import (
	"fmt"
	"mylab/cpagent/responses"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

var CurrentExpExcelFile *excelize.File

var dbSheets = []string{WellSample, ExperimentSheet, StepsStageSheet, TemplateSheet, TargetSheet}

const (
	TECSheet        = "tec"
	RTPCRSheet      = "rtpcr"
	TempLogs        = "temperature logs"
	WellSample      = "well_sample"
	ExperimentSheet = "experiment"
	StepsStageSheet = "steps_stages"
	TemplateSheet   = "template"
	TargetSheet     = "target_details"
)

func GetExcelFile(path, fileName string) (f *excelize.File) {

	f = excelize.NewFile()

	index := f.NewSheet(TECSheet)
	f.NewSheet(RTPCRSheet)
	f.NewSheet(TempLogs)
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	f.NewStyle(`{"alignment":{"horizontal":"center"}]}`)
	f.SetSheetFormatPr(RTPCRSheet, excelize.DefaultColWidth(25))
	f.SetSheetFormatPr(TempLogs, excelize.DefaultColWidth(40))
	f.SetSheetFormatPr(TECSheet, excelize.DefaultColWidth(40))

	f.Path = fmt.Sprintf("%v/%s.xlsx", path, fileName)
	logger.Infoln("file saved in path---------->", f.Path)

	return
}

func AddRowToExcel(file *excelize.File, sheet string, values []interface{}) (err error) {

	styleID, _ := file.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	rows, err := file.Rows(sheet)
	if err != nil {
		logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		return
	}
	rowCount := 1
	for rows.Next() {
		rowCount = rowCount + 1
	}

	for i, v := range values {
		cell, err := excelize.CoordinatesToCellName(i+1, rowCount)
		if err != nil {
			logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		}
		file.SetCellStyle(sheet, cell, cell, styleID)
		file.SetCellValue(sheet, cell, v)

	}

	if err = file.SaveAs(file.Path); err != nil {
		logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		return
	}

	return
}

func AddMergeRowToExcel(file *excelize.File, sheet string, values []interface{}, space int) {

	styleID, _ := file.NewStyle(`{"alignment":{"horizontal":"center"}}`)

	rows, err := file.Rows(sheet)
	if err != nil {
		logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		return
	}
	rowCount := 1
	for rows.Next() {
		rowCount = rowCount + 1
	}
	//first cell is always the start cell
	startCell, err := excelize.CoordinatesToCellName(1, rowCount)
	if err != nil {
		logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
	}
	file.SetCellValue(sheet, startCell, values[0])
	j := 1
	for i, v := range values {
		if i == 0 {
			continue
		}
		startCell, err := excelize.CoordinatesToCellName(j+1, rowCount)
		if err != nil {
			logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		}
		logger.Println("cell, value---------------->", startCell, v)
		file.SetCellStyle(sheet, startCell, startCell, styleID)
		file.SetCellValue(sheet, startCell, v)

		j = j + space + 1

		endCell, err := excelize.CoordinatesToCellName(j, rowCount)
		if err != nil {
			logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		}
		file.MergeCell(sheet, startCell, endCell)

	}

	if err = file.SaveAs(file.Path); err != nil {
		logger.Errorln(responses.ExcelSheetAddRowError, err.Error())
		return
	}
}

func SetExperimentExcelFile(file *excelize.File) {
	CurrentExpExcelFile = file
	return
}

func (s *pgStore) SetExcelHeadings(file *excelize.File, experimentID uuid.UUID) (err error) {
	for _, v := range dbSheets {
		file.NewSheet(v)
		file.SetSheetFormatPr(v, excelize.DefaultColWidth(40))
	}
	err = s.addWellSampleHeadings(file, experimentID)
	err = s.addExperimentHeadings(file, experimentID)
	err = s.addStagesAndStepsHeadings(file, experimentID)
	err = s.addTemplateHeadings(file, experimentID)
	err = s.addTargetDetailsHeadings(file, experimentID)
	return
}

func (s *pgStore) addWellSampleHeadings(file *excelize.File, expID uuid.UUID) (err error) {
	var heading []interface{}

	//add headings for sheet wells and samples
	rows, err := s.db.Query(getWellsListQuery, expID)
	if err != nil {
		logger.Errorln("Failed to fetch wells and samples", err.Error())
		return
	}
	columnNames, err := rows.Columns() // []string{"id", "name"}
	if err != nil {
		logger.Errorln("Failed to fetch columns for wells and samples", err.Error())
		return
	}
	for _, v := range columnNames {
		heading = append(heading, v)
	}
	err = AddRowToExcel(file, WellSample, heading)
	if err != nil {
		logger.Errorln("Failed to add row for wells and samples", err.Error())
		return
	}
	return

}

func (s *pgStore) addExperimentHeadings(file *excelize.File, expID uuid.UUID) (err error) {
	var heading []interface{}

	//add headings for sheet wells and samples
	rows, err := s.db.Query(getExperimentQuery, expID)
	if err != nil {
		logger.Errorln("Failed to fetch experiment details", err.Error())
		return
	}
	columnNames, err := rows.Columns() // []string{"id", "name"}
	if err != nil {
		logger.Errorln("Failed to fetch columns for experiments", err.Error())
		return
	}
	for _, v := range columnNames {
		heading = append(heading, v)
	}
	err = AddRowToExcel(file, ExperimentSheet, heading)
	if err != nil {
		logger.Errorln("Failed to add row for experiments", err.Error())
		return
	}
	return

}

func (s *pgStore) addStagesAndStepsHeadings(file *excelize.File, expID uuid.UUID) (err error) {
	var heading []interface{}

	//add headings for sheet wells and samples
	rows, err := s.db.Query(getStageStepQuery, expID)
	if err != nil {
		logger.Errorln("Failed to fetch stages and steps", err.Error())
		return
	}
	columnNames, err := rows.Columns() // []string{"id", "name"}
	if err != nil {
		logger.Errorln("Failed to fetch columns for stages and steps", err.Error())
		return
	}
	for _, v := range columnNames {
		heading = append(heading, v)
	}
	err = AddRowToExcel(file, StepsStageSheet, heading)
	if err != nil {
		logger.Errorln("Failed to add row for stages and steps", err.Error())
		return
	}
	return
}

func (s *pgStore) addTargetDetailsHeadings(file *excelize.File, expID uuid.UUID) (err error) {
	var heading []interface{}

	//add headings for sheet wells and samples
	rows, err := s.db.Query(getwellsConfigured, expID)
	if err != nil {
		logger.Errorln("Failed to fetch target details", err.Error())
		return
	}
	columnNames, err := rows.Columns() // []string{"id", "name"}
	if err != nil {
		logger.Errorln("Failed to fetch columns for target details", err.Error())
		return
	}
	for _, v := range columnNames {
		heading = append(heading, v)
	}
	err = AddRowToExcel(file, TargetSheet, heading)
	if err != nil {
		logger.Errorln("Failed to add row for target details", err.Error())
		return
	}
	return

}

func (s *pgStore) addTemplateHeadings(file *excelize.File, expID uuid.UUID) (err error) {
	var heading []interface{}
	var createdTemp Experiment

	err = s.db.Get(&createdTemp, getExperimentQuery, expID)
	if err != nil {
		logger.Errorln("Failed to experiment details", err.Error())
		return
	}
	//add headings for sheet wells and samples
	rows, err := s.db.Query(getTemplateQuery, createdTemp.TemplateID)
	if err != nil {
		logger.Errorln("Failed to fetch template details", err.Error())
		return
	}
	columnNames, err := rows.Columns() // []string{"id", "name"}
	if err != nil {
		logger.Errorln("Failed to fetch columns for template", err.Error())
		return
	}
	for _, v := range columnNames {
		heading = append(heading, v)
	}
	err = AddRowToExcel(file, TemplateSheet, heading)
	if err != nil {
		logger.Errorln("Failed to add row for template", err.Error())
		return
	}
	return

}
