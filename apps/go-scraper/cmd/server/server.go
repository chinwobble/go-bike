package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chinwobble/web-scraper/data"
	"github.com/chinwobble/web-scraper/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	err := data.InitDB()
	if err != nil {
		panic(err)
	}
	router := http.NewServeMux()

	handlers.RegisterProductRoutes(router)
	handlers.RegisterHomeRoutes(router)
	handlers.RegisterAlertRoutes(router)
	handlers.RegisterScrapeRoutes(router)
	server := &http.Server{
		Addr:         "localhost:9000",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on %v\n", server.Addr)
	server.ListenAndServe()
}
