package service

import (
	"NotificationOfBirthdays"
	"NotificationOfBirthdays/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(userId int, subscription NotificationOfBirthdays.Subscription) (int, error) {
	return s.repo.CreateSubscription(userId, subscription)
}

func (s *SubscriptionService) DeleteSubscription(userId, birthdayUserId int) error {
	return s.repo.DeleteSubscription(userId, birthdayUserId)
}

func (s *SubscriptionService) GetAllSubscriptions() ([]NotificationOfBirthdays.UserAndHisSubscriptions, error) {
	return s.repo.GetAllSubscriptions()
}
