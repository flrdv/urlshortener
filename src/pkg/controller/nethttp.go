package controller

import (
	"fmt"
	"io"
	"net/http"
	"urlshortener/src/pkg/model"
	"urlshortener/src/pkg/service"
)

type netHTTPController struct {
	shortenerService service.URLShortenerService
	linkService      service.LinkService
}

func NewHTTPController(
	shortenerService service.URLShortenerService,
	linkService service.LinkService) NetHTTPController {
	return netHTTPController{
		shortenerService: shortenerService,
		linkService:      linkService,
	}
}

func (n netHTTPController) CreateRedirect(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	to, err := n.linkService.Generate()
	if err != nil {
		http.Error(w, "failed to generate a shortened link", http.StatusInternalServerError)
		return
	}

	redirect := model.CreateRedirect{
		From: string(body),
		To:   to,
	}

	if err = n.shortenerService.CreateRedirect(redirect); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if _, err = fmt.Fprint(w, redirect.To); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (n netHTTPController) DoRedirect(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	redirectTo, err := n.shortenerService.GetRedirect(model.GetRedirect{
		From: req.URL.Path,
	})
	// TODO: add switch-case on error, maybe it's really internal server error
	if err != nil {
		http.Error(w, "redirect not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, redirectTo, http.StatusFound)
}
