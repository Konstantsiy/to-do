package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/Konstantsiy/todo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "qwerty12345"
	signingKey = "sdfgshfgh2635264^6235*"
	tokenTTL   = 2 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SigningString()
}
