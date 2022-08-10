package repository

import (
	"urlshortener/src/pkg/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBRepo interface {
	Init()
	CreateRedirect(redirect model.CreateRedirect) error
	GetRedirect(redirect model.GetRedirect) (to string, err error)
}

var schema = `
CREATE TABLE redirects (
	from text,
	to text
)
`

type Redirect struct {
	From string `db:"from"`
	To   string `db:"to"`
}

type postgreRepository struct {
	db PostgreDB
}

func NewDBRepo(db *sqlx.DB) DBRepo {
	return postgreRepository{
		db: NewPostgreDB(db),
	}
}

func (p postgreRepository) Init() {
	p.db.Init(schema)
}

func (p postgreRepository) CreateRedirect(createRedirect model.CreateRedirect) error {
	return p.db.Create("INSERT INTO redirects VALUES ($1, $2)", createRedirect)
}

func (p postgreRepository) GetRedirect(getRedirect model.GetRedirect) (to string, err error) {
	redirect, err := p.db.Get("SELECT from, to FROM redirects WHERE from=$1", getRedirect)
	return redirect.To, err
}
