package repositories

import "tasks-backend/internal/models"

type IUserRepository interface {
	CreateUser(task *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(task *models.User) error
	Delete(id uint) error
}
