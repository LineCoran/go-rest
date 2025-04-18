package service

import (
	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ExpenseItem interface {
}

type ExpenseList interface {
	Create(userId int, expese todo.Expense) (int, error)
	Delete(id string) (string, error)
	GetById(id int) (todo.Expense, error)
	GetAllByUserId(id int) ([]todo.Expense, error)
	Update(expenseId int, expense todo.Expense) (todo.Expense, error)
}

type Service struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ExpenseList:   NewExpenseListService(repos.ExpenseList),
		Authorization: NewAuthService(repos.Authorization),
	}
}
