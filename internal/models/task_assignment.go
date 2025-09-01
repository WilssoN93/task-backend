package models

import "time"

type TaskAssignment struct {
	ID          int       `gorm:"primaryKey"`
	CompletedAt time.Time `gorm:"type:date;column:completed_at"`
	IsCompleted bool      `gorm:"column:is_completed"`
	Deadline    time.Time `gorm:"type:date;column:deadline"`
	TaskID      uint      `gorm:"column:task_id" json:"-"`
	Task        Task      `gorm:"foreignKey:TaskID"`
	Users       []User    `gorm:"many2many:user_task_assignments"`
}
