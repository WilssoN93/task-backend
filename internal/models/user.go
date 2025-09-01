package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Image string `gorm:"column:image"`
	Tasks []Task `gorm:"many2many:user_task_assignments"`
}
