package service

import "NotificationOfBirthdays/pkg/repository"

type Authorization interface {
}

type Subscription interface {
}

type Service struct {
	Authorization
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
