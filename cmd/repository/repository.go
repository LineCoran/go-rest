package repository

type Authorization interface {
}

type ExpenseItem interface {
}

type ExpenseList interface {
}

type Repository struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewRepository() *Repository {
	return &Repository{}
}
