package rhubarb

import (
	"fmt"
	"testing"
)

func TestScrape(t *testing.T) {

	// First make a manifest

	manifest := Manifest{
		URL:                       "https://www.rightmove.co.uk/house-prices/detail.html?country=england&locationIdentifier=OUTCODE%5E1595&searchLocation=M41&referrer=landingPage",
		SingleItemSelector:        ".soldUnit",
		PaginationOffsetParameter: "index",
		PaginationPageSize:        25,
		PaginationNumPages:        1,
		SkipMissingFeatures:       true,
		AttributeSelectors: map[string]string{
			"address":  ".soldAddress",
			"price":    "table tr:first-child .soldPrice",
			"type":     "table tr:first-child .soldType",
			"date":     "table tr:first-child .soldDate",
			"bedrooms": "table tr:first-child .noBed",
		},
	}

	result := Scrape(manifest)

	fmt.Println(len(result))
}
