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
	// TODO Вынести название таблицы в константу
	createExpenseQuery := fmt.Sprintf("INSERT INTO expense (chat_id, category_id, amount, description) values ($1, $2, $3, $4) RETURNING id")
	row := r.db.QueryRow(createExpenseQuery, expense.ID, expense.CategoryId, expense.Amount, expense.Description)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (r *ExpenseListPostgres) Delete(id string) (string, error) {
	deleteExpenseQuery := fmt.Sprintf("DELETE from expense WHERE id = $1")
	row := r.db.QueryRow(deleteExpenseQuery, id)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return "error", err
	}
	return id, nil
}

func (r *ExpenseListPostgres) GetById(id int) (todo.Expense, error) {
	var expense todo.Expense
	query := fmt.Sprintf("SELECT id, category_id, amount, description FROM expense WHERE id = $1")
	err := r.db.Get(&expense, query, id)
	if err != nil {

		return todo.Expense{}, fmt.Errorf("failed to get expense by id: %w", err)
	}
	return expense, nil
}
