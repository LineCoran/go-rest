package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}
