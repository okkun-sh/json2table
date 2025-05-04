package json2table

import (
	"fmt"
	"strings"
)

func PrintTable(data []map[string]interface{}) {
	if len(data) == 0 {
		fmt.Println("No data")
		return
	}

	columns := collectAllKeys(data)
	widths := calcColumnWidths(columns, data)

	for _, col := range columns {
		fmt.Printf("| %-*s ", widths[col], col)
	}
	fmt.Println("|")

	for _, col := range columns {
		fmt.Print("| ")
		fmt.Print(strings.Repeat("-", widths[col]))
		fmt.Print(" ")
	}
	fmt.Println("|")

	for _, row := range data {
		for _, col := range columns {
			if row[col] == nil {
				fmt.Printf("| %-*s ", widths[col], "")
			} else {
				val := fmt.Sprintf("%v", row[col])
				fmt.Printf("| %-*s ", widths[col], val)
			}
		}
		fmt.Println("|")
	}
}

func collectAllKeys(data []map[string]interface{}) []string {
	keysSet := map[string]bool{}
	var keys []string
	for _, row := range data {
		for k := range row {
			if !keysSet[k] {
				keysSet[k] = true
				keys = append(keys, k)
			}
		}
	}
	return keys
}

func calcColumnWidths(columns []string, rows []map[string]interface{}) map[string]int {
	widths := make(map[string]int)
	for _, col := range columns {
		widths[col] = len(col)
		for _, row := range rows {
			val := fmt.Sprintf("%v", row[col])
			if len(val) > widths[col] {
				widths[col] = len(val)
			}
		}
	}
	return widths
}
