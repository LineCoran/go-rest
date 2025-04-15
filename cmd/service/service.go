package service

import "github.com/LineCoran/go-api/cmd/repository"

type Authorization interface {
}

type ExpenseItem interface {
}

type ExpenseList interface {
}

type Service struct {
	Authorization
	ExpenseItem
	ExpenseList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
