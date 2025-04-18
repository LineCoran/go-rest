package repository

import (
	"fmt"

	todo "github.com/LineCoran/go-api"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SignUp(user todo.User) (int, error) {
	var id int
	createExpenseQuery := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(createExpenseQuery, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUser(username string, password string) (int, error) {
	var id int
	createExpenseQuery := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	row := r.db.QueryRow(createExpenseQuery, username, password)
	if err := row.Scan(&id); err != nil {
		fmt.Printf("Error scanning id: %v\n", err)
		return 0, err
	}
	return id, nil
}
