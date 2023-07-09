package service

import (
	"todo/tasks"
)

func FromGrpcStatus(status TaskStatus) tasks.TaskStatus {
	switch status {
	case TaskStatus_NOT_STARTED:
		return tasks.NotStarted
	case TaskStatus_IN_PROGRESS:
		return tasks.InProgress
	case TaskStatus_BLOCKED:
		return tasks.Blocked
	case TaskStatus_COMPLETE:
		return tasks.Complete
	default:
		return tasks.NotStarted
	}
}

func ToGrpcStatus(status tasks.TaskStatus) TaskStatus {
	switch status {
	case tasks.NotStarted:
		return TaskStatus_NOT_STARTED
	case tasks.InProgress:
		return TaskStatus_IN_PROGRESS
	case tasks.Blocked:
		return TaskStatus_BLOCKED
	case tasks.Complete:
		return TaskStatus_COMPLETE
	default:
		return TaskStatus_NOT_STARTED
	}
}

func FromGrpcTask(task *Task) (*tasks.Task, error) {
	return tasks.NewTask(task.TaskId, task.Title, task.Description, task.DueDate, FromGrpcStatus(task.Status))
}

func ToGrpcTask(task *tasks.Task) (*Task, error) {
	grpcTask := Task{
		TaskId:      task.TaskId,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      ToGrpcStatus(task.Status)}
	return &grpcTask, nil
}
