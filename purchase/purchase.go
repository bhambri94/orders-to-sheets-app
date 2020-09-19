package purchase

import "strconv"

func GetFinalValuesFormatted(values [][]interface{}) [][]interface{} {
	var finalValues [][]interface{}
	iterator := 1
	SiNo := 0
	ItemCode := ""
	BlankRow := []interface{}{""}
	for iterator < len(values) {
		var row []interface{}
		if len(values[iterator]) < 5 {
			iterator++
			continue
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
			row = append(row, values[iterator][i])
			i++
		}
		finalValues = append(finalValues, row)
		iterator++
	}
	finalValues = append(finalValues, BlankRow)
	return finalValues
}
