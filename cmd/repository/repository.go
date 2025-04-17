package repository

import (
	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	SignUp(user todo.User) (int, error)
	GetUser(username string, password string) (int, error)
}

type ExpenseItem interface {
}

type ExpenseList interface {
	Create(userId int, expense todo.Expense) (int, error)
	Delete(id string) (string, error)
	GetById(id int) (todo.Expense, error)
}

type Repository struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ExpenseList:   NewExpenseListPostgres(db),
		Authorization: NewAuthRepository(db),
	}
}
