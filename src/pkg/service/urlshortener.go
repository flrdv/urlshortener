package service

import (
	"urlshortener/src/pkg/model"
	"urlshortener/src/pkg/repository"
)

type URLShortenerService interface {
	CreateRedirect(redirect model.CreateRedirect) error
	GetRedirect(redirect model.GetRedirect) (string, error)
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
	return u.dbRepo.CreateRedirect(redirect)
}

func (u urlShortenerService) GetRedirect(redirect model.GetRedirect) (string, error) {
	return u.dbRepo.GetRedirect(redirect)
}
