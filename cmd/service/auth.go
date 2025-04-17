package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

const salt = "'asdfaeweasdf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	hashed := generateHashedPassword(user.Password)
	fmt.Println(hashed)
	user.Password = hashed
	id, err := s.repo.SignUp(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func generateHashedPassword(password string) string {
	hash := sha1.New()

	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
