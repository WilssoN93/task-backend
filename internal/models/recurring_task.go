package models

import (
	"time"
)

type RecurringTask struct {
	ID                int       `gorm:"primaryKey"`
	TaskID            int       `gorm:"column:task_id"`
	Task              Task      `gorm:"foreignKey:TaskID"`
	RecurringInterval time.Time `gorm:"type:date;column:recurring_interval"`
	NextRefreshDate   time.Time `gorm:"type:date;column:next_refresh_date"`
	Users             []*User   `gorm:"many2many:user_recurring_tasks"`
}
