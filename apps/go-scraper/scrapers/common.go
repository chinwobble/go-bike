package scrapers

import (
	"fmt"
	"strings"
)

type ScrapeResult struct {
	Id            string
	Title         string
	CurrentPrice  float64
	OriginalPrice float64
	Source        string
	Sku           string
	Size          string
	Brand         string
	Category      string
	Options       []string
	Properties    map[string]string
}

func CheckGiantVariant(
	title string,
	variant giantVariant) bool {
	title = strings.ReplaceAll(strings.ToLower(title), " disc", "")
	title = strings.ReplaceAll(title, " -", "")
	title = strings.ReplaceAll(title, "\"new year sale\" ", "")
	title = strings.ReplaceAll(title, "\"march madness sale\" ", "")
	candidate1 := fmt.Sprintf("%s %s %s", variant.Family, variant.SubFamily, variant.Number)
	candidate2 := strings.Replace(candidate1, "advanced", "adv", -1)
	candidate3 := strings.Replace(candidate1, "advanced", "", -1)
	if variant.ContainsAdvanced {
		return strings.Contains(title, candidate1) || strings.Contains(title, candidate2) || strings.Contains(title, candidate3)
	}
	return strings.Contains(title, candidate1)
}

type giantVariant struct {
	Family           string
	SubFamily        string
	Number           string
	ContainsAdvanced bool
}

func (result ScrapeResult) GetModel() string {
	if strings.ToLower(result.Brand) == "liv" && strings.ToLower(result.Category) == "road-bike" {
		models := map[string]giantVariant{
			"Avail Advanced SL 2":  {"avail", "advanced sl", "2", true},
			"Avail Advanced SL 1":  {"avail", "advanced sl", "1", true},
			"Avail Advanced SL 0":  {"avail", "advanced sl", "0", true},
			"Avail Advanced Pro 2": {"avail", "advanced pro", "2", true},
			"Avail Advanced Pro 1": {"avail", "advanced pro", "1", true},
			"Avail Advanced Pro 0": {"avail", "advanced pro", "0", true},
			"Avail Advanced 2":     {"avail", "advanced", "2", true},
			"Avail Advanced 1":     {"avail", "advanced", "1", true},
			"Avail Advanced 0":     {"avail", "advanced", "0", true},

			"Langma Advanced SL 2":  {"langma", "advanced sl", "2", true},
			"Langma Advanced SL 1":  {"langma", "advanced sl", "1", true},
			"Langma Advanced SL 0":  {"langma", "advanced sl", "0", true},
			"Langma Advanced Pro 2": {"langma", "advanced pro", "2", true},
			"Langma Advanced Pro 1": {"langma", "advanced pro", "1", true},
			"Langma Advanced Pro 0": {"langma", "advanced pro", "0", true},
			"Langma Advanced 2":     {"langma", "advanced", "2", true},
			"Langma Advanced 1":     {"langma", "advanced", "1", true},
			"Langma Advanced 0":     {"langma", "advanced", "0", true},

			"Brava Advanced SL 2":  {"brava", "advanced sl", "2", true},
			"Brava Advanced SL 1":  {"brava", "advanced sl", "1", true},
			"Brava Advanced SL 0":  {"brava", "advanced sl", "0", true},
			"Brava Advanced Pro 2": {"brava", "advanced pro", "2", true},
			"Brava Advanced Pro 1": {"brava", "advanced pro", "1", true},
			"Brava Advanced Pro 0": {"brava", "advanced pro", "0", true},
			"Brava Advanced 2":     {"brava", "advanced", "2", true},
			"Brava Advanced 1":     {"brava", "advanced", "1", true},
			"Brava Advanced 0":     {"brava", "advanced", "0", true},

			"Avail AR 5": {"avail", "ar", "5", false},
			"Avail AR 4": {"avail", "ar", "4", false},
			"Avail AR 3": {"avail", "ar", "3", false},
			"Avail AR 2": {"avail", "ar", "2", false},
			"Avail AR 1": {"avail", "ar", "1", false},
			"Avail AR 0": {"avail", "ar", "0", false},
		}
		for model, variant := range models {
			if CheckGiantVariant(result.Title, variant) {
				return model
			}
		}
	}
	if strings.ToLower(result.Brand) == "giant" && strings.ToLower(result.Category) == "road-bike" {
		models := map[string]giantVariant{
			"TCR Advanced SL 2":  {"tcr", "advanced sl", "2", true},
			"TCR Advanced SL 1":  {"tcr", "advanced sl", "1", true},
			"TCR Advanced SL 0":  {"tcr", "advanced sl", "0", true},
			"TCR Advanced Pro 2": {"tcr", "advanced pro", "2", true},
			"TCR Advanced Pro 1": {"tcr", "advanced pro", "1", true},
			"TCR Advanced Pro 0": {"tcr", "advanced pro", "0", true},
			"TCR Advanced 2":     {"tcr", "advanced", "2", true},
			"TCR Advanced 1":     {"tcr", "advanced", "1", true},
			"TCR Advanced 0":     {"tcr", "advanced", "0", true},

			"Propel Advanced SL 2":  {"propel", "advanced sl", "2", true},
			"Propel Advanced SL 1":  {"propel", "advanced sl", "1", true},
			"Propel Advanced SL 0":  {"propel", "advanced sl", "0", true},
			"Propel Advanced Pro 2": {"propel", "advanced pro", "2", true},
			"Propel Advanced Pro 1": {"propel", "advanced pro", "1", true},
			"Propel Advanced Pro 0": {"propel", "advanced pro", "0", true},
			"Propel Advanced 2":     {"propel", "advanced", "2", true},
			"Propel Advanced 1":     {"propel", "advanced", "1", true},
			"Propel Advanced 0":     {"propel", "advanced", "0", true},

			"Defy Advanced SL 2":  {"defy", "advanced sl", "2", true},
			"Defy Advanced SL 1":  {"defy", "advanced sl", "1", true},
			"Defy Advanced SL 0":  {"defy", "advanced sl", "0", true},
			"Defy Advanced Pro 2": {"defy", "advanced pro", "2", true},
			"Defy Advanced Pro 1": {"defy", "advanced pro", "1", true},
			"Defy Advanced Pro 0": {"defy", "advanced pro", "0", true},
			"Defy Advanced 2":     {"defy", "advanced", "2", true},
			"Defy Advanced 1":     {"defy", "advanced", "1", true},
			"Defy Advanced 0":     {"defy", "advanced", "0", true},

			"Contend AR 5": {"contend", "ar", "5", false},
			"Contend AR 4": {"contend", "ar", "4", false},
			"Contend AR 3": {"contend", "ar", "3", false},
			"Contend AR 2": {"contend", "ar", "2", false},
			"Contend AR 1": {"contend", "ar", "1", false},
			"Contend AR 0": {"contend", "ar", "0", false},
		}
		for model, variant := range models {
			if CheckGiantVariant(result.Title, variant) {
				return model
			}
		}
	}
	return ""
}
