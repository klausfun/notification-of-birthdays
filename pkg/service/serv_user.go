package service

import (
	"NotificationOfBirthdays"
	"NotificationOfBirthdays/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetUsers() ([]NotificationOfBirthdays.Author, error) {
	return s.repo.GetUsers()
}
