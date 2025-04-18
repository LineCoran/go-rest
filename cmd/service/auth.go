package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "secret"
	signingKey = "secret"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	hashed := generateHashedPassword(user.Password)
	user.Password = hashed
	id, err := s.repo.SignUp(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

type tokensClaims struct {
	jwt.StandardClaims
	UserID int `json:"id"`
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	hashed := generateHashedPassword(password)
	id, err := s.repo.GetUser(username, hashed)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokensClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokensClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokensClaims)

	if !ok {
		return 0, errors.New("token claims are not type *tokensClaims")
	}

	return claims.UserID, nil
}

func generateHashedPassword(password string) string {
	hash := sha1.New()

	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
