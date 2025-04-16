package todo

type Expense struct {
	ID          int     `json:"chat_id" binding:"required" db:"id"`
	CategoryId  int     `json:"category_id" binding:"required" db:"category_id"`
	Description string  `json:"description" db:"description"`
	Amount      float64 `json:"amount" binding:"required" db:"amount"`
}
