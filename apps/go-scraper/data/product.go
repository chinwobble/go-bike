package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Product struct {
	Id            string
	CurrentPrice  string
	OriginalPrice string
	Source        string
	Sku           string
	Brand         string
	Category      string
	Title         string
	Model         string
	Options       string
	Size          string
	Groupset      string
}

func newProduct(record []string) Product {
	return Product{
		Id:            record[0],
		CurrentPrice:  record[1],
		OriginalPrice: record[2],
		Source:        record[3],
		Sku:           record[4],
		Brand:         record[5],
		Category:      record[6],
		Title:         record[7],
		Model:         record[8],
		Options:       record[9],
		Size:          record[10],
		Groupset:      record[11],
	}

}

func ReadCSVFile(file string) ([]Product, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil, err
	}
	products := make([]Product, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}
		product := newProduct(record)

		products = append(products, product)
	}

	return products, nil
}
