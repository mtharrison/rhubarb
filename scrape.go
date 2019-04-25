package rhubarb

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

// Scrape actually scrapes the site using your manifest
func Scrape(manifest Manifest) []map[string]string {

	var results []map[string]string

	resultChan := make(chan (map[string]string))

	go func() {
		for m := range resultChan {
			results = append(results, m)
		}
	}()

	c := colly.NewCollector(colly.Async(true))

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	c.OnHTML(manifest.SingleItemSelector, func(e *colly.HTMLElement) {

		m := make(map[string]string, len(manifest.AttributeSelectors))

		for key, selector := range manifest.AttributeSelectors {
			text := strings.TrimSpace(e.ChildText(selector))
			if len(text) == 0 && manifest.SkipMissingFeatures {
				return
			}
			m[key] = e.ChildText(selector)
		}

		resultChan <- m
	})

	for i := 0; i < manifest.PaginationNumPages; i++ {
		url := fmt.Sprintf("%s&%s=%d", manifest.URL, manifest.PaginationOffsetParameter, manifest.PaginationPageSize*i)
		c.Visit(url)
	}

	c.Wait()
	close(resultChan)

	return results
}
