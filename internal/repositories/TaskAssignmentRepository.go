package repositories

import (
	"errors"
	"tasks-backend/internal/db"
	"tasks-backend/internal/models"
)

type taskAssignmentRepositry struct{}

func NewTaskAssignmentRepository() ITaskAssignmentRepository {
	return &taskAssignmentRepositry{}
}

func (r *taskAssignmentRepositry) GetAllAssignments() ([]*models.TaskAssignment, error) {
	var assignments []*models.TaskAssignment

	err := db.TaskDb.Find(assignments).Error
	return assignments, err
}

func (r *taskAssignmentRepositry) GetAssignmentsByTaskId(taskId uint) ([]*models.TaskAssignment, error) {
	var assignments []*models.TaskAssignment
	err := db.TaskDb.Find(assignments).Where("task_id = ? ", taskId).Error
	return assignments, err
}

func (r *taskAssignmentRepositry) CreateAssignment(assignment *models.TaskAssignment) error {

	if assignment.ID != 0 {
		return errors.New("assignment id must not be assigned when creating task assignment")
	}

	return db.TaskDb.Create(assignment).Error
}

func (r *taskAssignmentRepositry) UpdateAssignment(assignment *models.TaskAssignment) error {
	return db.TaskDb.Updates(assignment).Error
}

func (r *taskAssignmentRepositry) DeleteAssignment(id uint) error {
	return db.TaskDb.Delete(&models.TaskAssignment{}, id).Error
}
