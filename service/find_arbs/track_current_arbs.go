package find_arbs

import (
	"OISA_2x_sistem/arbitrage"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"time"
)

var OldArbsBySport map[string][]arbitrage.Arb
var TodaysArbs []arbitrage.Arb

var fileName = "arbs_log.xlsx"
var sheetName = "log_all"

func openExcelFile() *excelize.File {

	file, err := excelize.OpenFile(fileName)

	if err != nil {
		return createNewExcelFile()
	}

	index := file.GetSheetIndex(sheetName)
	file.SetActiveSheet(index)

	return file
}

func createNewExcelFile() *excelize.File {

	file := excelize.NewFile()
	index := file.NewSheet(sheetName)
	file.SetActiveSheet(index)
	setColumnNames(file)

	return file
}

func setColumnNames(file *excelize.File) {

	t := reflect.TypeOf(arbitrage.Arb{})

	file.SetCellValue(sheetName, fmt.Sprintf("%c1", rune(65)), "Date")
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		file.SetCellValue(sheetName, fmt.Sprintf("%c1", rune(65+i+1)), field.Name)
	}
}

func LogArbToExcelFile(arb *arbitrage.Arb) {

	file := openExcelFile()

	rows := file.GetRows(sheetName)
	row := len(rows) + 1

	runeA := 65
	file.SetCellValue(sheetName, fmt.Sprintf("%c%d", rune(runeA), row), time.Now().Format("2006-01-02"))

	v := reflect.ValueOf(*arb)
	runeB := 66
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		structFieldName := v.Type().Field(i).Name
		excelColName := file.GetCellValue(sheetName, fmt.Sprintf("%c1", rune(runeB+i)))

		if structFieldName != excelColName {
			fmt.Println("ERROR: struct field name and excel column name are not the same")
			continue
		}

		file.SetCellValue(sheetName, fmt.Sprintf("%c%d", rune(runeB+i), row), field.Interface())
	}

	// Save the file
	err := file.SaveAs(fileName)
	if err != nil {
		fmt.Println(err)
	}
}
