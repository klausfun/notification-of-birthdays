package service

import (
	"NotificationOfBirthdays"
	"NotificationOfBirthdays/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "rjdiokveo78f3kovnjo1oofviperfvjn"
	signingKey = "uheofjrvihj23injonvjn94nvjn3"
	tokenTTl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user NotificationOfBirthdays.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(password, email string) (NotificationOfBirthdays.Author, string, error) {
	user, err := s.repo.GetUser(generatePasswordHash(password), email)
	if err != nil {
		return NotificationOfBirthdays.Author{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	strToken, err := token.SignedString([]byte(signingKey))
	return user, strToken, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
