package repositories

import "tasks-backend/internal/models"

type ITaskRepository interface {
	CreateTask(task *models.Task) error
	DeleteTask(id uint) error
	UpdateTask(task *models.Task) error
	GetAllTasks() ([]*models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
}
