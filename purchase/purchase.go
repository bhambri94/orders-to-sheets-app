package purchase

import (
	"strconv"
	"strings"
)

func GetFinalValuesFormatted(values [][]interface{}, LastIndex int, OrderType string) [][]interface{} {
	var finalValues [][]interface{}
	iterator := 0
	SiNo := LastIndex
	ItemCode := ""
	BlankRow := []interface{}{""}
	finalValues = append(finalValues, BlankRow)
	for iterator < len(values) {
		var row []interface{}
		if len(values[iterator]) < 5 {
			iterator++
			continue
		}

		if !strings.Contains(values[iterator][4].(string), OrderType) {
			iterator++
			continue
		}

		if len(values[iterator][3].(string)) > 10 {
			values[iterator][3] = values[iterator][3].(string)[5:10]
		}
		if len(values[iterator][5].(string)) > 10 {
			values[iterator][5] = values[iterator][5].(string)[5:10]
		}
		if len(values[iterator][12].(string)) > 10 {
			values[iterator][12] = values[iterator][12].(string)[5:10]
		}
		if len(values[iterator][21].(string)) > 10 {
			values[iterator][21] = values[iterator][21].(string)[5:10]
		}

		if ItemCode != values[iterator][6] {
			finalValues = append(finalValues, BlankRow)
			ItemCode = values[iterator][6].(string)
			SiNo++
		}
		i := 1
		firstElement := true
		for i < len(values[iterator]) {
			if firstElement {
				values[iterator][0] = strconv.Itoa(SiNo)
				row = append(row, values[iterator][0])
				firstElement = false
			}
			if i > 22 {
				i++
				continue
			}
			row = append(row, values[iterator][i])
			i++
		}
		finalValues = append(finalValues, row)
		iterator++
	}
	finalValues = append(finalValues, BlankRow)
	return finalValues
}
