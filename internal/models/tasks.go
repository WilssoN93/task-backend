package models

type Task struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}
