package db

import (
	"fmt"
	"mylab/cpagent/responses"
	"os"
	"strconv"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
)

const ContextKeyUsername = "username"

const (
	TECSheet   = "tec"
	RTPCRSheet = "rtpcr"
	TempLogs   = "temperature logs"
)

type ProcessSequence struct {
	ID             uuid.UUID `db:"id" json:"process_id" validate:"required"`
	SequenceNumber int64     `db:"sequence_num" json:"sequence_num" validate:"required,gte=1"`
}

func CalculateTimeInSeconds(t string) (totalTime int64, err error) {

	var hours, minutes, seconds int64
	timeArr := strings.Split(t, ":")
	if len(timeArr) != 3 {
		err = fmt.Errorf("time format isn't of the form HH:MM:SS")
		return 0, err
	}

	hours, err = parseIntRange(timeArr[0], "hours", 0, 24)
	if err != nil {
		return 0, err
	}

	minutes, err = parseIntRange(timeArr[1], "minutes", 0, 59)
	if err != nil {
		return 0, err
	}

	seconds, err = parseIntRange(timeArr[2], "seconds", 0, 59)
	if err != nil {
		return 0, err
	}

	totalTime = hours*60*60 + minutes*60 + seconds

	return
}

func parseIntRange(timeString, unit string, min, max int64) (value int64, err error) {
	value, err = strconv.ParseInt(timeString, 10, 64)
	if err != nil || value > max || value < min {
		err = fmt.Errorf("please check %v format, valid range: [%d,%d]", unit, min, max)
		return 0, err
	}
	return
}

func GetExcelFile(path, fileName string) (f *excelize.File) {
	// logging output to file and console
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		// ignore error and try creating log output file
	}

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

	f.Path = fmt.Sprintf("%v/%s_%v.xlsx", path, fileName, time.Now().Unix())
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

		j = j + space - 1

		endCell, err := excelize.CoordinatesToCellName(j+1, rowCount)
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
