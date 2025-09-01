package repositories

import (
	"errors"
	"tasks-backend/internal/db"
	"tasks-backend/internal/models"
)

type taskRepository struct{}

func NewTaskRepository() ITaskRepository {
	return &taskRepository{}
}

func (r *taskRepository) CreateTask(task *models.Task) error {

	if task.ID != 0 {
		return errors.New("task id cannot be set when creating task")
	}

	return db.TaskDb.Create(task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	return db.TaskDb.Delete(&models.Task{}, id).Error
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	return db.TaskDb.Updates(task).Error
}

func (r *taskRepository) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task

	err := db.TaskDb.Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint) (*models.Task, error) {

	var task *models.Task
	err := db.TaskDb.Find(&task, id).Error
	return task, err
}
