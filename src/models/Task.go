package models

type Task struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Priority    string
}
