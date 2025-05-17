package repository

import (
	"database/sql"
	"errors"
	"fmt"

	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryListPostgres struct {
	db *sqlx.DB
}

func NewCategoryListPostgres(db *sqlx.DB) *CategoryListPostgres {
	return &CategoryListPostgres{db: db}
}

func (r *CategoryListPostgres) CreateOne(userId int, category todo.ExpenseCategory) (int, error) {
	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, name) values ($1, $2) RETURNING id", categoryTable)
	row := r.db.QueryRow(createCategoryQuery, userId, category.Name)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (r *CategoryListPostgres) GetByName(userId int, name string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = $1 AND name = $2", categoryTable)
	
	err := r.db.QueryRow(query, userId, name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrCategoryNotFound
		}
		return 0, fmt.Errorf("failed to get category by name: %w", err)
	}
	
	return id, nil
}

func (r *CategoryListPostgres) IsExists(userId int, name string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = $1 AND name = $2)", categoryTable)
	
	err := r.db.QueryRow(query, userId, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check category existence: %w", err)
	}
	
	return exists, nil
}

func (r *CategoryListPostgres) GetAllByUserId(userId int) ([]todo.ExpenseCategory, error) {
	var categories []todo.ExpenseCategory
	query := fmt.Sprintf("SELECT id, name, created_at FROM %s WHERE user_id = $1", categoryTable)
	err := r.db.Select(&categories, query, userId)
	if err != nil {
		return []todo.ExpenseCategory{}, fmt.Errorf("failed to get category by id: %w", err)
	}

	if len(categories) == 0 {
		return []todo.ExpenseCategory{}, nil
	}

	return categories, nil
}

func (r *CategoryListPostgres) DeleteCategory(userId int, categoryId int) (int, error) {

	query := fmt.Sprintf("DELETE from %s WHERE $1 = user_id AND $2 = id RETURNING id", categoryTable);

	result, err := r.db.Exec(query, userId, categoryId)
	if err != nil {
		fmt.Printf("Error deleting expense: %v\n", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("no rows affected - record with id %d not found", categoryId)
	}

	return categoryId, nil
}