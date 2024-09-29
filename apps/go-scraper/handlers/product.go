package handlers

import (
	"log/slog"
	"net/http"

	"github.com/chinwobble/web-scraper/components"
	"github.com/chinwobble/web-scraper/data"
	"github.com/chinwobble/web-scraper/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	bikeType := r.PathValue("type")
	slog.Info("product bike types", "type", bikeType)
	ctx := utils.WithPageContext(r.Context(), utils.PageContextValue{
		AreaName: "products",
	})
	if bikeType == "index" || bikeType == "" {
		products := make([]data.Product, 0)
		components.ProductsPage("Road Bikes", products).Render(ctx, w)
		return
	}
	if bikeType == "road-bikes" {
		products, err := data.ReadCSVFile("bex-road-bikes.csv")
		if err != nil {
			panic("missing file")
		}
		products = utils.Filter(products, func(p data.Product) bool {
			return p.Brand == "Giant"
		})
		components.ProductsPage("Road Bikes", products).Render(ctx, w)
		return
	}
}

func RegisterProductRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", GetProducts)
	router.HandleFunc("GET /products/{type}", GetProducts)
}
