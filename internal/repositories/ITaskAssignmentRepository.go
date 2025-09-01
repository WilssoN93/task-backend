package repositories

import "tasks-backend/internal/models"

type ITaskAssignmentRepository interface {
	CreateAssignment(assignment *models.TaskAssignment) error
	UpdateAssignment(assignment *models.TaskAssignment) error
	GetAllAssignments() ([]*models.TaskAssignment, error)
	GetAssignmentsByTaskId(taskId uint) ([]*models.TaskAssignment, error)
	DeleteAssignment(id uint) error
}
