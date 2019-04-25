package main

import (
	"encoding/csv"
	"github.com/mtharrison/rhubarb"
	"log"
	"os"
)

func main() {
	// First make a manifest

	manifest := rhubarb.Manifest{
		URL:                       "https://www.rightmove.co.uk/house-prices/detail.html?country=england&locationIdentifier=OUTCODE%5E1595&searchLocation=M41&referrer=landingPage",
		SingleItemSelector:        ".soldUnit",
		PaginationOffsetParameter: "index",
		PaginationPageSize:        25,
		PaginationNumPages:        40,
		SkipMissingFeatures:       true,
		AttributeSelectors: map[string]string{
			"address":  ".soldAddress",
			"price":    "table tr:first-child .soldPrice",
			"type":     "table tr:first-child .soldType",
			"date":     "table tr:first-child .soldDate",
			"bedrooms": "table tr:first-child .noBed",
		},
	}

	result := rhubarb.Scrape(manifest)

	t := tabularize(result)

	file, err := os.OpenFile("houses.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(t)
	if err != nil {
		log.Fatal(err)
	}

}

func tabularize(input []map[string]string) [][]string {
	var output [][]string
	var header []string

	for key := range input[0] {
		header = append(header, key)
	}

	output = append(output, header)

	for _, row := range input {
		var newRow []string
		for _, value := range header {
			newRow = append(newRow, row[value])
		}
		output = append(output, newRow)
	}

	return output
}
