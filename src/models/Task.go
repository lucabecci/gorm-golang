package models

import (
	"gorm.io/gorm"
)

var database *gorm.DB

//Task is for the logic and the ORM migration
type Task struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Priority    string
}

//InstanceDatabase is a function for save the db instance in the var db of the Task model
func InstanceDatabase(db *gorm.DB) {
	database = db
}

//InsertTask is a function for insert a one task in the db
func InsertTask(title string, description string, priority string) (Task, bool) {
	task := Task{
		Title:       title,
		Description: description,
		Priority:    priority,
	}
	result := database.Create(&task)
	if result.RowsAffected < 1 {
		return Task{}, false
	}
	return Task{
		task.ID,
		task.Title,
		task.Description,
		task.Priority,
	}, true
}

func GetTask(id string) (Task, bool) {
	var task Task

	result := database.Find(&task, id)
	result.Scan(&task)

	if task.ID == 0 {
		return Task{}, false
	}

	return task, true
}
