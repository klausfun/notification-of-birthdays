package service

import (
	"NotificationOfBirthdays"
	"NotificationOfBirthdays/pkg/repository"
)

type Authorization interface {
	CreateUser(user NotificationOfBirthdays.User) (int, error)
	GenerateToken(password, email string) (NotificationOfBirthdays.Author, string, error)
	ParseToken(token string) (int, error)
}

type Subscription interface {
}

type Profile interface {
	GetUsers() ([]NotificationOfBirthdays.Author, error)
}

type Service struct {
	Authorization
	Subscription
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile:       NewProfileService(repos.Profile),
	}
}
