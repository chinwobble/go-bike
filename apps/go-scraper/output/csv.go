package output

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chinwobble/web-scraper/scrapers"
)

func WriteToCSV(results []scrapers.ScrapeResult, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	// Initialize the CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{
		"Id",
		"CurrentPrice",
		"OriginalPrice",
		"Source",
		"Sku",
		"Brand",
		"Category",
		"Title",
		"Model",
		"Options",
		"Size",
		"Groupset",
	})

	// Write all records to the CSV file
	for _, record := range results {

		if err := writer.Write([]string{
			record.Id,
			strconv.FormatFloat(record.CurrentPrice, 'f', -1, 64),
			strconv.FormatFloat(record.OriginalPrice, 'f', -1, 64),
			record.Source,
			record.Sku,
			record.Brand,
			record.Category,
			record.Title,
			record.GetModel(),
			strings.Join(record.Options, ", "),
			record.Properties["Size"],
			record.Properties["Groupset"],
		}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
}
