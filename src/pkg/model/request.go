package model

import "urlshortener/src/pkg/types"

type CreateRedirect struct {
	From types.URL `json:"from"`
	To   types.URL `json:"to"`
}

type GetRedirect struct {
	From types.URL `json:"from"`
}
