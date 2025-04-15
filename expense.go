package todo

type Expense struct {
	ID          int    `json:"chat_id" binding:"required"`
	CategoryId  int    `json:"category_id" binding:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" binding:"required"`
}
