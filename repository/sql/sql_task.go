package sql

import (
	"gorm.io/gorm"
	"todo/tasks"
)

type Task struct {
	gorm.Model
	// Names have to be public for automigrate for working
	Title       string
	Description string
	DueDate     string
	Status      string
}

func (task *Task) fromSql() (*tasks.Task, error) {
	return tasks.NewTask(int32(task.ID), task.Title, task.Description, task.DueDate, tasks.FromString(task.DueDate))
}
