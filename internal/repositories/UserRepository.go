package repositories

import (
	"errors"
	"tasks-backend/internal/db"
	"tasks-backend/internal/models"
)

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {

	if user.ID != 0 {
		return errors.New("user id cannot be set when creating user")
	}

	return db.TaskDb.Create(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return db.TaskDb.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.TaskDb.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := db.TaskDb.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return db.TaskDb.Updates(user).Error
}
