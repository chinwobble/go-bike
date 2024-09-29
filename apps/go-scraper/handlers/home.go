package handlers

import (
	"net/http"

	"github.com/chinwobble/web-scraper/components"
	"github.com/chinwobble/web-scraper/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithPageContext(r.Context(), utils.PageContextValue{
		AreaName: "home",
	})
	components.ScrapesPage(
		"home",
	).Render(ctx, w)
}

func RegisterHomeRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", GetHome)
	// router.HandleFunc("GET /scrape", GetBikes)
}
