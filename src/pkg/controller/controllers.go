package controller

import (
	"net/http"
)

type NetHTTPController interface {
	CreateRedirect(w http.ResponseWriter, req *http.Request)
	DoRedirect(w http.ResponseWriter, req *http.Request)
}
