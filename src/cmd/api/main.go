package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"urlshortener/src/internal/config"
	"urlshortener/src/pkg/controller"
	"urlshortener/src/pkg/repository"
	"urlshortener/src/pkg/service"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(RunURLShortener(cfg))
}

func RunURLShortener(cfg config.Config) error {
	db, err := sqlx.Connect("postgres", renderConnectionURL(cfg))
	if err != nil {
		return err
	}

	dbRepo := repository.NewDBRepo(db)
	shortenerService := service.NewURLShortenerService(dbRepo)
	linkRepo := repository.NewLinkRepository()
	linkService := service.NewLinkService(linkRepo)
	netHTTPController := controller.NewHTTPController(shortenerService, linkService)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/shorten" {
			http.Error(w, "not found", http.StatusNotFound)
		}
		if req.Method != http.MethodPost {
			http.Error(w, "disallowed method", http.StatusMethodNotAllowed)
			return
		}

		netHTTPController.CreateRedirect(w, req)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "disallowed method", http.StatusMethodNotAllowed)
			return
		}

		netHTTPController.DoRedirect(w, req)
	})

	return http.ListenAndServe(cfg.Addr, mux)
}

func renderConnectionURL(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBAddr, cfg.DBName,
	)
}
