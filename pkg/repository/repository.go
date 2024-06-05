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
	CreateSubscription(userId int, subscription NotificationOfBirthdays.Subscription) (int, error)
	DeleteSubscription(userId, birthdayUserId int) error
	GetAllSubscriptions() ([]NotificationOfBirthdays.UserAndHisSubscriptions, error)
}

type Profile interface {
	GetUsers() ([]NotificationOfBirthdays.Author, error)
}

type Repository struct {
	Authorization
	Subscription
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Profile:       NewProfilePostgres(db),
		Subscription:  NewSubscriptionPostgres(db),
	}
}
