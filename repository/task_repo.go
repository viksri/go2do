package repository

import (
	"todo/tasks"
)

type TaskRepo interface {
	GetTask(taskId int32) (*tasks.Task, error)
	GetAllTasks() ([]*tasks.Task, error)
	AddTask(title string, description string, dueDate string) (*tasks.Task, error)
	UpdateTask(
		taskId int32,
		newTitle string,
		newDesc string,
		newDueDate string,
		newStatus tasks.TaskStatus) (*tasks.Task, error)
}
