package repo

import (
	"golang-boilerplate/main/models"
	"golang-boilerplate/main/service"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := service.DB.Exec(query, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1`
	err := service.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
