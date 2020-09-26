package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Configs struct {
	BCFLSTRSpreadsheetID  string `json:"BCFLSTRSpreadsheetID"`
	BCFLGPSpreadsheetID   string `json:"BCFLGPSpreadsheetID"`
	BCFLLOCSpreadsheetID  string `json:"BCFLLOCSpreadsheetID"`
	BCFLFGTSpreadsheetID  string `json:"BCFLFGTSpreadsheetID"`
	VRLSTRSpreadsheetID   string `json:"VRLSTRSpreadsheetID"`
	VRLGPSpreadsheetID    string `json:"VRLGPSpreadsheetID"`
	VRLLOCSpreadsheetID   string `json:"VRLLOCSpreadsheetID"`
	VRLFGTSpreadsheetID   string `json:"VRLFGTSpreadsheetID"`
	SheetNameWithoutRange string `json:"SheetNameWithoutRange"`
	MSSQLHost             string `json:"MSSQLHost"`
	BCFLDatabaseName      string `json:"BCFLDatabaseName"`
	VRLDatabaseName       string `json:"VRLDatabaseName"`
	UserName              string `json:"UserName"`
	Password              string `json:"Password"`
	Query                 string `json:"Query"`
	BCFLOldDateInHours    int    `json:"BCFLOldDateInHours"`
	VRLOldDateInHours     int    `json:"VRLOldDateInHours"`
}

var (
	Configurations = Configs{}
)

func SetConfig() {
	input, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	error := json.Unmarshal(input, &Configurations)
	if error != nil {
		fmt.Println("Config file is missing in root directory")
		panic(error)
	} else {
		fmt.Println("Follwing values has been picked from config values:")
		fmt.Println(Configurations)
	}
}
