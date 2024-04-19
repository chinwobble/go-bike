package main

import (
	"github.com/chinwobble/web-scraper/output"
	"github.com/chinwobble/web-scraper/scrapers"
)

func main() {
	result := scrapers.Scrape(
		"https://www.bikeexchange.com.au/s/road-bikes",
		"road-bike",
	)
	output.WriteToCSV(result, "bex-road-bikes.csv")
	// resultFrames := scrapers.Scrape("https://www.bikeexchange.com.au/s/road-frames")
	// writeToCSV(resultFrames, "bex-road-frames.csv")
}
