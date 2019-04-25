package rhubarb

// A Manifest tells the scraper how to find the collection on the page
type Manifest struct {
	URL                       string
	SingleItemSelector        string
	AttributeSelectors        map[string]string
	PaginationOffsetParameter string
	PaginationPageSize        int
	PaginationNumPages        int
	SkipMissingFeatures       bool
}
