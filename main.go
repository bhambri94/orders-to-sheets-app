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
	HoursCount := configs.Configurations.OldDateInHours
	fromDateTime = currentTime.Add(time.Duration(-HoursCount) * time.Hour).Format("2006-01-02")
	toDateTime = currentTime.Format("2006-01-02")
	fmt.Println("Fetching results from Date: "+fromDateTime, " to "+toDateTime)
	Month := strings.ToUpper(time.Now().Month().String())
	dbValues := db.GetLatestDataFromSQL(fromDateTime, toDateTime)
	SheetName := configs.Configurations.SheetNameWithoutRange + Month
	sheets.CreateSheetIfNotPresent(SheetName)
	valuesFromSheet := sheets.BatchGet(configs.Configurations.SheetNameWithoutRange + Month + "!A2:Z5000")
	AppendIndex := len(valuesFromSheet) + 1
	LastIndex := 0
	if len(valuesFromSheet) > 2 {
		var err error
		LastIndex, err = strconv.Atoi(valuesFromSheet[len(valuesFromSheet)-1][0])
		if err != nil {
			LastIndex = 1
		}
	}
	finalvalues := purchase.GetFinalValuesFormatted(dbValues, LastIndex)
	fmt.Println("FinalValues from Sheet are:")
	fmt.Println(finalvalues)
	sheets.BatchAppend(SheetName+"!A"+strconv.Itoa(AppendIndex), finalvalues)
}
