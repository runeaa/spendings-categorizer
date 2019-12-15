package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

const excelFileName = "path/to/your/excel"

func main() {
	categories := setUpGroups()
	sumCategory := make(map[string]float64)

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
	}
	for _, row := range xlFile.Sheets[1].Rows {
		text := row.Cells[3].String()
		amount := row.Cells[4].String()
		found := false

		for k, c := range categories {
			for _, vv := range c {
				if strings.Contains(text, vv) && vv != "" {
					value := sumCategory[k]
					parsedValue, err := strconv.ParseFloat(amount, 64)

					if err != nil {
					}
					newTotal := value + Abs(parsedValue)
					sumCategory[k] = newTotal
					found = true
				}
			}
		}
		if !found {
			fmt.Println("Expenses not found", text)
			found = false
		}
	}
	fmt.Println(sumCategory)
}

func setUpGroups() map[string][]string {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
	}

	categories := make(map[string][]string)
	nextIsCategory := true
	var category string

	for _, row := range xlFile.Sheets[0].Rows {
		for _, cell := range row.Cells {

			text := cell.String()

			if !nextIsCategory && text != "!" {
				categories[category] = append(categories[category], strings.Trim(text, " "))
			}

			if nextIsCategory {
				category = text
				nextIsCategory = false
			}

			if text == "!" {
				nextIsCategory = true
			}
		}
	}

	return categories
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
