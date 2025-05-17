package repository

import (
	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	SignUp(user todo.User) (int, error)
	GetUser(username string, password string) (int, error)
	IsExist(username string) (int, error)
}

type ExpenseItem interface {
}

type ExpenseList interface {
	Create(userId int, expense todo.Expense) (int, error)
	Delete(id string) (string, error)
	GetById(id int) (todo.Expense, error)
	GetAllByUserId(id int) ([]todo.Expense, error)
	Update(id int, expense todo.Expense) (todo.Expense, error)
}

type CategoryList interface {
	CreateOne(userId int, category todo.ExpenseCategory) (int, error)
	GetByName(userId int, name string) (int, error)
	IsExists(userId int, name string) (bool, error)
	GetAllByUserId(userId int) ([]todo.ExpenseCategory, error)
	DeleteCategory(userId int, categoryId int) (int, error)
}

type Repository struct {
	Authorization
	ExpenseItem
	ExpenseList
	CategoryList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ExpenseList:   NewExpenseListPostgres(db),
		Authorization: NewAuthRepository(db),
		CategoryList: NewCategoryListPostgres(db),
	}
}
