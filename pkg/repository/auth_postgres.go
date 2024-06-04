package repository

import (
	"NotificationOfBirthdays"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user NotificationOfBirthdays.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email)"+
		" values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(password, email string) (NotificationOfBirthdays.Author, error) {
	var user NotificationOfBirthdays.Author
	query := fmt.Sprintf("SELECT id, name, email FROM %s"+
		" WHERE password_hash=$1 AND email=$2", userTable)
	err := r.db.Get(&user, query, password, email)

	return user, err
}
