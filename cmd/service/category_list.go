package service

import (
	todo "github.com/LineCoran/go-api"
	"github.com/LineCoran/go-api/cmd/repository"
)

type CategoryListService struct {
	repo repository.CategoryList
}


func NewCategoryListService(repo repository.CategoryList) *CategoryListService {
	return &CategoryListService{repo: repo}
}

func (s *CategoryListService) CreateOne(userId int, expense todo.ExpenseCategory) (int, error) {

	isExist, err := s.repo.IsExists(userId, expense.Name)

	if (err != nil) {
		return 0, err
	}

	if (isExist) {
		return 0, nil
	}
	
	return s.repo.CreateOne(userId, expense)
}

func (s *CategoryListService) GetByName(userId int, name string) (int, error) {
	return s.repo.GetByName(userId, name)
}

func (s *CategoryListService) GetAllByUserId(id int) ([]todo.ExpenseCategory, error) {
	return s.repo.GetAllByUserId(id)
}

func (s *CategoryListService) DeleteCategory(userId int, categoryId int) (int, error) {
	return s.repo.DeleteCategory(userId, categoryId)
}