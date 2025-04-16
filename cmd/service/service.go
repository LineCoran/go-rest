package service

import (
	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

type Authorization interface {
}

type ExpenseItem interface {
}

type ExpenseList interface {
	Create(userId int, expese todo.Expense) (int, error)
	Delete(id string) (string, error)
	GetById(id int) (todo.Expense, error)
}

type Service struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{ExpenseList: NewExpenseListService(repos.ExpenseList)}
}
