package repository

import (
	"NotificationOfBirthdays"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) GetUsers() ([]NotificationOfBirthdays.Author, error) {
	var users []NotificationOfBirthdays.Author
	query := fmt.Sprintf("SELECT id, name, email FROM %s", userTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, err
}
