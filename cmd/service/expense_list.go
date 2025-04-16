package service

import (
	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

type ExpenseListService struct {
	repo repository.ExpenseList
}

func NewExpenseListService(repo repository.ExpenseList) *ExpenseListService {
	return &ExpenseListService{repo: repo}
}

func (s *ExpenseListService) Create(userId int, expense todo.Expense) (int, error) {
	return s.repo.Create(userId, expense)
}

func (s *ExpenseListService) Delete(id string) (string, error) {
	return s.repo.Delete(id)
}
