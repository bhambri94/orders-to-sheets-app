package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/bhambri94/orders-to-sheets-app/configs"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetLatestDataFromSQL(databaseName string, fromDateTime string, toDateTime string) [][]interface{} {
	var DBConnection *sql.DB
	var err error
	connectString := "sqlserver://" + configs.Configurations.UserName + ":" + configs.Configurations.Password + "@" + configs.Configurations.MSSQLHost + "?database=" + databaseName + "&connection+timeout=300"
	println("opening sql connection with connstring:" + connectString)

	RetryCounter := 0
	ConnectionSuccess := false
	for RetryCounter < 5 && !ConnectionSuccess {
		DBConnection, err = sql.Open("mssql", connectString)
		defer DBConnection.Close()
		if err != nil {
			RetryCounter++
			if strings.Contains(err.Error(), "Client.Timeout") {
			} else {
				println("Open Error:", err)
				log.Fatal(err)
			}
		} else {
			ConnectionSuccess = true
		}
	}

	println("Running Query -> " + configs.Configurations.Query + " '" + fromDateTime + "' , '" + toDateTime + "'")
	Rows, err := DBConnection.Query(configs.Configurations.Query + " '" + fromDateTime + "' , '" + toDateTime + "'")

	if err != nil {
		log.Fatal(err)
	}

	var finalValues [][]interface{}
	BlankRow := make([]interface{}, 3)
	finalValues = append(finalValues, BlankRow)
	Columns, _ := Rows.Columns()
	rawResult := make([][]byte, len(Columns))

	dest := make([]interface{}, len(Columns)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for Rows.Next() {
		result := make([]interface{}, len(Columns))
		err = Rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)

		fmt.Println("adding rows to finalValues")
		finalValues = append(finalValues, result)
	}

	println("closing connection")
	DBConnection.Close()
	return finalValues

}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
