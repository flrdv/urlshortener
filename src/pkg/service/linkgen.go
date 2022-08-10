package service

import (
	"urlshortener/src/pkg/repository"
)

type LinkService interface {
	Generate() (string, error)
}

type linkService struct {
	linkRepository repository.LinkRepository
}

func NewLinkService(linkRepo repository.LinkRepository) LinkService {
	return linkService{
		linkRepository: linkRepo,
	}
}

func (s linkService) Generate() (string, error) {
	return s.Generate()
}
