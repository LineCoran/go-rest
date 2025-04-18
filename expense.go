package todo

import "time"

type Expense struct {
	Id          int       `json:"id" db:"id"`
	CategoryId  int       `json:"category_id" binding:"required" db:"category_id"`
	Description string    `json:"description" db:"description"`
	Amount      float64   `json:"amount" binding:"required" db:"amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
