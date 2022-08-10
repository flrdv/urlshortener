package repository

import (
	"urlshortener/src/pkg/model"
	"urlshortener/src/pkg/types"
)

type DBRepo interface {
	Init()
	CreateRedirect(redirect model.CreateRedirect) error
	GetRedirect(redirect model.GetRedirect) (to types.URL, err error)
}
