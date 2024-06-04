package repository

import (
	"NotificationOfBirthdays"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user NotificationOfBirthdays.User) (int, error)
	GetUser(password, email string) (NotificationOfBirthdays.Author, error)
}

type Subscription interface {
}

type Repository struct {
	Authorization
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
