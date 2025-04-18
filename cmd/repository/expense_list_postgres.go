package repository

import (
	"fmt"

	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

type ExpenseListPostgres struct {
	db *sqlx.DB
}

func NewExpenseListPostgres(db *sqlx.DB) *ExpenseListPostgres {
	return &ExpenseListPostgres{db: db}
}

func (r *ExpenseListPostgres) Create(userId int, expense todo.Expense) (int, error) {
	var id int
	createExpenseQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id, amount, description) values ($1, $2, $3, $4) RETURNING id", expenseTable)
	row := r.db.QueryRow(createExpenseQuery, userId, expense.CategoryId, expense.Amount, expense.Description)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (r *ExpenseListPostgres) Delete(id string) (string, error) {
	deleteExpenseQuery := fmt.Sprintf("DELETE from %s WHERE id = $1", expenseTable)
	result, err := r.db.Exec(deleteExpenseQuery, id)
	if err != nil {
		fmt.Printf("Error deleting expense: %v\n", err)
		return "error", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error", err
	}

	if rowsAffected == 0 {
		return "error", fmt.Errorf("no rows affected - record with id %s not found", id)
	}

	return id, nil
}

func (r *ExpenseListPostgres) GetById(id int) (todo.Expense, error) {
	var expense todo.Expense
	query := fmt.Sprintf("SELECT id, category_id, amount, description FROM %s WHERE id = $1", expenseTable)
	err := r.db.Get(&expense, query, id)
	if err != nil {
		return todo.Expense{}, fmt.Errorf("failed to get expense by id: %w", err)
	}
	return expense, nil
}

func (r *ExpenseListPostgres) GetAllByUserId(userId int) ([]todo.Expense, error) {
	var expenses []todo.Expense
	query := fmt.Sprintf("SELECT id, category_id, amount, created_at, description FROM %s WHERE user_id = $1", expenseTable)
	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return []todo.Expense{}, fmt.Errorf("failed to get expense by id: %w", err)
	}

	if len(expenses) == 0 {
		return []todo.Expense{}, nil
	}

	return expenses, nil
}
