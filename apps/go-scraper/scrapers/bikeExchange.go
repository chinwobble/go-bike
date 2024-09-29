package scrapers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type bikeExchangeAdvertItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type ProductAddToCartProps struct {
	DisplayFavouriteStore bool `json:"displayFavouriteStore"`
	DefaultEmailID        int  `json:"defaultEmailId"`
	Variants              []struct {
		VariantID        int    `json:"variant_id"`
		Option           string `json:"option"`
		Description      string `json:"description"`
		CountOnHand      int    `json:"count_on_hand"`
		InfiniteQuantity bool   `json:"infinite_quantity"`
		// Sku                      any    `json:"sku"`
		URL                      string `json:"url"`
		ImageURL                 string `json:"image_url"`
		CheckoutURL              string `json:"checkout_url"`
		Price                    string `json:"price"`
		PriceInSessionCurrency   string `json:"price_in_session_currency"`
		PriceHTMLWithMetadata    string `json:"price_html_with_metadata"`
		AfterpayEligible         bool   `json:"afterpay_eligible"`
		AfterpayInstalmentAmount string `json:"afterpay_instalment_amount"`
		ZipCoEligible            bool   `json:"zip_co_eligible"`
		// ZipCoMerchantPublicKey   any    `json:"zip_co_merchant_public_key"`
		// ZipCoMerchantEnvironment any    `json:"zip_co_merchant_environment"`
		YotpoEnabled   bool   `json:"yotpo_enabled"`
		YotpoProductID string `json:"yotpo_product_id"`
		// YotpoAppKey              any    `json:"yotpo_app_key"`
	} `json:"variants"`
	// Inventories           []any `json:"inventories"`
	InfiniteQuantityLimit int `json:"infiniteQuantityLimit"`
	// PreferredCollectionStoreID any      `json:"preferred_collection_store_id"`
	// SessionLocation            any      `json:"sessionLocation"`
	AvailableSaleTypes  []string `json:"availableSaleTypes"`
	InitialAvailability string   `json:"initialAvailability"`
	// IsStoreListVisibleWithoutInput any      `json:"isStoreListVisibleWithoutInput"`
	// OnlyShowServiceAreaMatches     any      `json:"onlyShowServiceAreaMatches"`
	// ShowParentInfo                 any      `json:"showParentInfo"`
	AdvertType              string `json:"advert_type"`
	AdvertTitle             string `json:"advert_title"`
	AdvertURL               string `json:"advert_url"`
	AdvertCanHaveCustomForm bool   `json:"advert_can_have_custom_form"`
	LowItemLimit            int    `json:"low_item_limit"`
	Prompt                  string `json:"prompt"`
	AnalyticsData           struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Brand    string `json:"brand"`
	} `json:"analytics_data"`
	// Questions            []any `json:"questions"`
	HasVariantSelection  bool `json:"has_variant_selection"`
	HasQuantitySelection bool `json:"has_quantity_selection"`
	IsMStore             bool `json:"isMStore"`
	QuickViewData        struct {
		AdvertTitle         string `json:"advert_title"`
		SellerLogoURL       string `json:"seller_logo_url"`
		SellerShopURL       string `json:"seller_shop_url"`
		BusinessName        string `json:"business_name"`
		SellerCity          string `json:"seller_city"`
		PriceLine           string `json:"price_line"`
		LowestOriginalPrice string `json:"lowest_original_price"`
		PriceLineClass      string `json:"price_line_class"`
	} `json:"quick_view_data"`
	// SizeChart             any `json:"size_chart"`
	// SizeMatrix            any `json:"size_matrix"`
	// PreferredDeliveryType any `json:"preferredDeliveryType"`
}

type ProductPageSceneProps struct {
	AdvertAttributesData []struct {
		Label string `json:"label"`
		Value any    `json:"value"`
		URL   string `json:"url,omitempty"`
	} `json:"advertAttributesData"`
	AdvertDetailsData struct {
		Description string `json:"description"`
		Sections    []struct {
			Name    string `json:"name"`
			Heading string `json:"heading"`
			Content string `json:"content"`
		} `json:"sections"`
		RequiredSections  []any `json:"requiredSections"`
		VariantPromotions []any `json:"variantPromotions"`
		ShowGenericTerms  bool  `json:"showGenericTerms"`
		Documents         []any `json:"documents"`
	} `json:"advertDetailsData"`
}

func Scrape(url string, category string) []ScrapeResult {
	c := colly.NewCollector(
		colly.AllowedDomains("www.shop.bikeexchange.com.au", "bikeexchange.com.au"),
		colly.CacheDir("./bikeexchange_cache"),
	)

	resultsMap := make(map[string]ScrapeResult, 2)
	results := make([]ScrapeResult, 0)
	count := 0
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// get list items
	c.OnHTML(".AdvertTile", func(e *colly.HTMLElement) {
		count++
		var advert bikeExchangeAdvertItem
		json.Unmarshal([]byte(e.Attr("data-advert-tracking-data")), &advert)
		priceRegular, err := convertPriceStringToFloat(e.ChildText(".Price-regular"))
		if err != nil {
			priceRegular, err = convertPriceStringToFloat(e.ChildText(".Price-was"))
		}
		priceSale, _ := convertPriceStringToFloat(e.ChildText(".Price-sale"))
		// not a sale
		if priceSale == 0 {
			priceSale = priceRegular
		}
		// fmt.Println(count, advert.Id, advert.Brand, priceSale, priceRegular, advert.Name)
		resultsMap[advert.Id] = ScrapeResult{
			Id:            advert.Id,
			Brand:         advert.Brand,
			Title:         advert.Name,
			CurrentPrice:  priceSale,
			OriginalPrice: priceRegular,
			Source:        "BikeExchange",
			Sku:           "???",
			Category:      category,
			Options:       make([]string, 0),
			Properties:    make(map[string]string),
		}

		// visit the details page
		e.Request.Visit(e.ChildAttr(".t-advertTileLink", "href"))
	})

	c.OnHTML("a.next_page.btn", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("[data-react-class=ProductPageScene]", func(h *colly.HTMLElement) {
		var props ProductPageSceneProps
		json.Unmarshal([]byte(h.Attr("data-react-props")), &props)
		var id string
		for idx, prop := range props.AdvertAttributesData {
			if prop.Label == "ID" {
				id = fmt.Sprintf("%.0f", prop.Value)
				fmt.Printf("match found %.0f\n", prop.Value)
				break
			}
			if idx == len(props.AdvertAttributesData)-1 {
				fmt.Println("no match found")
			}
		}
		if entry, exists := resultsMap[id]; exists {
			for _, prop := range props.AdvertAttributesData {
				if prop.Label == "ID" {
					continue
				}
				fmt.Printf("Adding entry %s \n", prop.Label)
				entry.Properties[prop.Label] = fmt.Sprintf("%v", prop.Value)
			}
		} else {
			fmt.Printf("Cannot find %s \n", id)
		}
	})

	// JSON.parse(document.querySelectorAll("[data-react-class=ProductAddToCart]")[0].attributes.getNamedItem("data-react-props").value)
	c.OnHTML("[data-react-class=ProductAddToCart]", func(h *colly.HTMLElement) {
		// for some reason this is duplicated per page

		var props ProductAddToCartProps
		json.Unmarshal([]byte(h.Attr("data-react-props")), &props)
		if entry, exists := resultsMap[props.AnalyticsData.ID]; exists {
			// fmt.Printf("variant count: %d\n", len(props.Variants))
			for _, variant := range props.Variants {
				if contains(entry.Options, variant.Option) {
					continue
				}
				entry.Options = append(entry.Options, variant.Option)
				resultsMap[props.AnalyticsData.ID] = entry
			}
		}
	})

	// get title
	c.OnHTML(".t-productHeaderHeading", func(e *colly.HTMLElement) {
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})
	// visit the initial index
	c.Visit(url)
	for _, v := range resultsMap {
		results = append(results, v)
	}

	return results
}

func convertPriceStringToFloat(input string) (float64, error) {
	// convert "Now AU$2,949" to 2949
	output := strings.ReplaceAll(input, "AU", "")
	output = strings.ReplaceAll(output, "$", "")
	output = strings.ReplaceAll(output, ",", "")
	output = strings.ReplaceAll(output, "Now ", "")
	return strconv.ParseFloat(output, 32)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
