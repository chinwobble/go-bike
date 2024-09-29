package handlers

import (
	"net/http"

	"github.com/chinwobble/web-scraper/components"
	"github.com/chinwobble/web-scraper/utils"
)

func GetScrapes(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithPageContext(r.Context(), utils.PageContextValue{
		AreaName: "scrapes",
	})
	components.ScrapesPage(
		"scrapes",
	).Render(ctx, w)
}

func RegisterScrapeRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /scrapes", GetScrapes)
	// router.HandleFunc("GET /scrape", GetBikes)
}
