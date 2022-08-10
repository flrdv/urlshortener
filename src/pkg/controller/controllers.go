package controller

import (
	"encoding/json"
	"net/http"
	"urlshortener/src/pkg/model"
	"urlshortener/src/pkg/service"
)

type NetHTTPController interface {
	CreateRedirect(w http.ResponseWriter, req *http.Request)
	DoRedirect(w http.ResponseWriter, req *http.Request)
}

type netHTTPController struct {
	shortenerService service.URLShortenerService
}

func NewHTTPController(shortenerService service.URLShortenerService) NetHTTPController {
	return netHTTPController{
		shortenerService: shortenerService,
	}
}

func (n netHTTPController) CreateRedirect(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	redirect := new(model.CreateRedirect)

	if err := json.NewDecoder(req.Body).Decode(redirect); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := n.shortenerService.CreateRedirect(*redirect); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (n netHTTPController) DoRedirect(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	redirect := new(model.GetRedirect)

	if err := json.NewDecoder(req.Body).Decode(redirect); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	redirectTo, err := n.shortenerService.GetRedirect(*redirect)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, string(redirectTo), http.StatusFound)
}
