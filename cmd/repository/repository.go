package repository

import (
	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type ExpenseItem interface {
}

type ExpenseList interface {
	Create(userId int, expense todo.Expense) (int, error)
}

type Repository struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{ExpenseList: NewExpenseListPostgres(db)}
}
