package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/bhambri94/orders-to-sheets-app/configs"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetLatestDataFromSQL(fromDateTime string) [][]interface{} {
	var DBConnection *sql.DB
	var err error
	connectString := "sqlserver://" + configs.Configurations.UserName + ":" + configs.Configurations.Password + "@" + configs.Configurations.MSSQLHost + "?database=" + configs.Configurations.DatabaseName + "&connection+timeout=300"
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

	println("Running Query -> " + configs.Configurations.Query)
	Rows, err := DBConnection.Query(configs.Configurations.Query)

	// _, err := db.ExecContext(ctx, configs.Configurations.Query, sql.Named("Arg1", sql.Out{Dest: &outArg}))
	if err != nil {
		log.Fatal(err)
	}

	var finalValues [][]interface{}
	BlankRow := make([]interface{}, 3)
	finalValues = append(finalValues, BlankRow)
	fmt.Println("adding rows to finalValues")
	// var NumericToString []uint8
	// var QuotationDate string
	// var LostDate string
	Columns, _ := Rows.Columns()
	// singleRow := make([]interface{}, len(Columns))
	// for i, _ := range Columns {
	// 	singleRow[i] = new(sql.RawBytes)
	// 	//check column name, if it is id, and you know it is integer
	// 	//vals[i] = new(int)
	// }

	rawResult := make([][]byte, len(Columns))
	result := make([]interface{}, len(Columns))

	dest := make([]interface{}, len(Columns)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for Rows.Next() {
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

		// if err := Rows.Scan(singleRow...); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(singleRow[0].(string))

		// if err := Rows.Scan(&singleRow[0], &singleRow[1], &singleRow[2], &QuotationDate, &singleRow[4], &singleRow[5], &singleRow[6], &NumericToString, &singleRow[8], &singleRow[9], &singleRow[10], &LostDate); err != nil {
		// 	log.Fatal(err)
		// }
		// singleRow[7] = B2S(NumericToString)
		// singleRow[3] = QuotationDate
		// singleRow[11] = LostDate[:10]
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
