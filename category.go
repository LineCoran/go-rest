package todo

import "time"

type ExpenseCategory struct {
	Id          int       `json:"id" db:"id"`
	Name  		string       `json:"category_name" binding:"required" db:"name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
