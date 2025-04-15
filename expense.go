package todo

type Expense struct {
	ID          int    `json:"chat_id"`
	CategoryId  int    `json:"category_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
