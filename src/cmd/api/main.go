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
	netHTTPController := controller.NewHTTPController(shortenerService)

	http.HandleFunc("/shorten", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			netHTTPController.DoRedirect(w, req)
		case http.MethodPost:
			netHTTPController.CreateRedirect(w, req)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return http.ListenAndServe(cfg.Addr, nil)
}

func renderConnectionURL(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBAddr, cfg.DBName,
	)
}
