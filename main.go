package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bhambri94/orders-to-sheets-app/configs"
	"github.com/bhambri94/orders-to-sheets-app/db"
	"github.com/bhambri94/orders-to-sheets-app/purchase"
	"github.com/bhambri94/orders-to-sheets-app/sheets"
)

func main() {
	configs.SetConfig()
	var fromDateTime string
	var toDateTime string
	currentTime := time.Now()
	HoursCount := configs.Configurations.BCFLOldDateInHours
	fromDateTime = currentTime.Add(time.Duration(-HoursCount) * time.Hour).Format("2006-01-02")
	toDateTime = currentTime.Format("2006-01-02")
	fmt.Println("Fetching BCFL results from Date: "+fromDateTime, " to "+toDateTime)
	Month := strings.ToUpper(time.Now().Month().String())
	dbValues := db.GetLatestDataFromSQL(configs.Configurations.BCFLDatabaseName, fromDateTime, toDateTime)
	orderType := []string{"STR", "GP", "LOC", "FGT"}
	iterator := 0
	for iterator < len(orderType) {
		SheetName := configs.Configurations.SheetNameWithoutRange + orderType[iterator] + "_" + Month
		if iterator == 0 {
			sheets.SetSpreadSheetID(configs.Configurations.BCFLSTRSpreadsheetID)
		} else if iterator == 1 {
			sheets.SetSpreadSheetID(configs.Configurations.BCFLGPSpreadsheetID)
		} else if iterator == 2 {
			sheets.SetSpreadSheetID(configs.Configurations.BCFLLOCSpreadsheetID)
		} else if iterator == 3 {
			sheets.SetSpreadSheetID(configs.Configurations.BCFLFGTSpreadsheetID)
		}
		sheets.CreateSheetIfNotPresent(SheetName)
		valuesFromSheet := sheets.BatchGet(configs.Configurations.SheetNameWithoutRange + orderType[iterator] + "_" + Month + "!A2:W5000")
		AppendIndex := len(valuesFromSheet) + 1
		LastIndex := 0
		if len(valuesFromSheet) > 2 {
			var err error
			LastIndex, err = strconv.Atoi(valuesFromSheet[len(valuesFromSheet)-1][0])
			if err != nil {
				LastIndex = 1
			}
		}
		finalvalues := purchase.GetFinalValuesFormatted(dbValues, LastIndex, orderType[iterator])
		fmt.Println("FinalValues for BCFL Sheet are:")
		fmt.Println(finalvalues)
		sheets.BatchAppend(SheetName+"!A"+strconv.Itoa(AppendIndex), finalvalues)
		time.Sleep(1000 * time.Millisecond)
		iterator++
	}

	HoursCount = configs.Configurations.VRLOldDateInHours
	fromDateTime = currentTime.Add(time.Duration(-HoursCount) * time.Hour).Format("2006-01-02")
	toDateTime = currentTime.Format("2006-01-02")
	fmt.Println("Fetching VRL results from Date: "+fromDateTime, " to "+toDateTime)
	dbValues = db.GetLatestDataFromSQL(configs.Configurations.VRLDatabaseName, fromDateTime, toDateTime)
	iterator = 0
	for iterator < len(orderType) {
		SheetName := configs.Configurations.SheetNameWithoutRange + orderType[iterator] + "_" + Month
		if iterator == 0 {
			sheets.SetSpreadSheetID(configs.Configurations.VRLSTRSpreadsheetID)
		} else if iterator == 1 {
			sheets.SetSpreadSheetID(configs.Configurations.VRLGPSpreadsheetID)
		} else if iterator == 2 {
			sheets.SetSpreadSheetID(configs.Configurations.VRLLOCSpreadsheetID)
		} else if iterator == 3 {
			sheets.SetSpreadSheetID(configs.Configurations.VRLFGTSpreadsheetID)
		}
		sheets.CreateSheetIfNotPresent(SheetName)
		valuesFromSheet := sheets.BatchGet(configs.Configurations.SheetNameWithoutRange + orderType[iterator] + "_" + Month + "!A2:W5000")
		AppendIndex := len(valuesFromSheet) + 1
		LastIndex := 0
		if len(valuesFromSheet) > 2 {
			var err error
			LastIndex, err = strconv.Atoi(valuesFromSheet[len(valuesFromSheet)-1][0])
			if err != nil {
				LastIndex = 1
			}
		}
		finalvalues := purchase.GetFinalValuesFormatted(dbValues, LastIndex, orderType[iterator])
		fmt.Println("FinalValues for VRL Sheet are:")
		fmt.Println(finalvalues)
		sheets.BatchAppend(SheetName+"!A"+strconv.Itoa(AppendIndex), finalvalues)
		time.Sleep(1000 * time.Millisecond)
		iterator++
	}
}
