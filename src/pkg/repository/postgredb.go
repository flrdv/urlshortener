package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"urlshortener/src/pkg/model"
)

type PostgreDB interface {
	Init(schema string)
	Create(query string, redirect model.CreateRedirect) error
	Get(query string, from model.GetRedirect) (Redirect, error)
}

type postgreDB struct {
	db *sqlx.DB
}

func NewPostgreDB(db *sqlx.DB) PostgreDB {
	return postgreDB{
		db: db,
	}
}

func (p postgreDB) Init(schema string) {
	p.db.MustExec(schema)
}

func (p postgreDB) Get(query string, getRedirect model.GetRedirect) (Redirect, error) {
	redirect := new(Redirect)
	err := p.db.Get(redirect, query, string(getRedirect.From))

	return *redirect, err
}

func (p postgreDB) Create(query string, redirect model.CreateRedirect) error {
	tx, err := p.db.Begin()
	if err != nil {
		return tryRollback(err, tx.Rollback)
	}

	if _, err = tx.Query(query, string(redirect.From), string(redirect.To)); err != nil {
		return tryRollback(err, tx.Rollback)
	}

	return tx.Commit()
}

func tryRollback(srcErr error, rollback func() error) error {
	if err := rollback(); err != nil {
		return errors.Wrap(srcErr, err.Error())
	}

	return srcErr
}
