package repository

import (
	"fmt"
	"strings"

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
    
    fields := []string{"user_id", "category_id", "amount", "description"}
    values := []interface{}{userId, expense.CategoryId, expense.Amount, expense.Description}
    
    if !expense.CreatedAt.IsZero() {
        fields = append(fields, "created_at")
        values = append(values, expense.CreatedAt)
    }
    placeholders := make([]string, len(values))
    for i := range values {
        placeholders[i] = fmt.Sprintf("$%d", i+1)
    }
    
    query := fmt.Sprintf(
        "INSERT INTO %s (%s) VALUES (%s) RETURNING id",
        expenseTable,
        strings.Join(fields, ", "),
        strings.Join(placeholders, ", "),
    )
    
    row := r.db.QueryRow(query, values...)
    if err := row.Scan(&id); err != nil {
        return 0, fmt.Errorf("error creating expense: %w", err)
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

func (r *ExpenseListPostgres) GetAllByUserId(userId int) ([]todo.UserExpense, error) {
	var expenses []todo.UserExpense
	query := fmt.Sprintf("SELECT e.id, e.amount, c.name, e.created_at FROM %s e RIGHT JOIN %s c ON e.category_id = c.id WHERE e.user_id = $1", expenseTable, categoryTable)
	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return []todo.UserExpense{}, fmt.Errorf("failed to get expense by id: %w", err)
	}

	if len(expenses) == 0 {
		return []todo.UserExpense{}, nil
	}

	return expenses, nil
}

func (r *ExpenseListPostgres) Update(expenseId int, expense todo.Expense) (todo.Expense, error) {
	query := fmt.Sprintf(`
        UPDATE %s 
        SET category_id = $1, 
            amount = $2, 
            description = $3 
        WHERE id = $4
        RETURNING id, category_id, amount, description, created_at
    `, expenseTable)

	var updatedExpense todo.Expense
	err := r.db.QueryRowx(query,
		expense.CategoryId,
		expense.Amount,
		expense.Description,
		expenseId,
	).StructScan(&updatedExpense)

	if err != nil {
		return todo.Expense{}, fmt.Errorf("failed to update expense: %w", err)
	}

	return updatedExpense, nil
}
