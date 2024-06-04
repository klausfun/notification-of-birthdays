package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Subscription interface {
}

type Repository struct {
	Authorization
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
