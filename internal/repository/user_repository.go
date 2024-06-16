package repository

import (
	"auth-service/internal/model"
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
