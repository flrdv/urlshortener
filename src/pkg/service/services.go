package service

import (
	"urlshortener/src/pkg/model"
	"urlshortener/src/pkg/repository"
	"urlshortener/src/pkg/types"
)

type URLShortenerService interface {
	CreateRedirect(redirect model.CreateRedirect) error
	GetRedirect(redirect model.GetRedirect) (types.URL, error)
}

type urlShortenerService struct {
	dbRepo repository.DBRepo
}

func NewURLShortenerService(dbRepo repository.DBRepo) URLShortenerService {
	return urlShortenerService{
		dbRepo: dbRepo,
	}
}

func (u urlShortenerService) CreateRedirect(redirect model.CreateRedirect) error {
	return u.dbRepo.CreateRedirect(redirect.From, redirect.To)
}

func (u urlShortenerService) GetRedirect(redirect model.GetRedirect) (types.URL, error) {
	return u.dbRepo.GetRedirect(redirect.From)
}
