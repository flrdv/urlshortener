package repository

import "github.com/teris-io/shortid"

type LinkRepository interface {
	Generate() (string, error)
}

type linkRepository struct {
}

func NewLinkRepository() LinkRepository {
	return linkRepository{}
}

func (s linkRepository) Generate() (string, error) {
	return shortid.Generate()
}
