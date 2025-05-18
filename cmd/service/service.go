package service

import (
	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
	IsExist(username string) (int, error)
}

type ExpenseItem interface {
}

type CategoryList interface {
	CreateOne(userId int, category todo.ExpenseCategory) (int, error)
	GetByName(userId int, name string) (int, error)
	GetAllByUserId(userId int) ([]todo.ExpenseCategory, error)
	DeleteCategory(userId int, categoryId int) (int, error)
}

type ExpenseList interface {
	Create(userId int, expese todo.Expense) (int, error)
	Delete(id string) (string, error)
	GetById(id int) (todo.Expense, error)
	GetAllByUserId(id int) ([]todo.UserExpense, error)
	Update(expenseId int, expense todo.Expense) (todo.Expense, error)
}

type Service struct {
	Authorization
	ExpenseItem
	ExpenseList
	CategoryList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ExpenseList:   NewExpenseListService(repos.ExpenseList),
		Authorization: NewAuthService(repos.Authorization),
		CategoryList: NewCategoryListService(repos.CategoryList),
	}
}
