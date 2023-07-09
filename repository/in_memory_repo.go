package repository

import (
	"fmt"
	"sort"
	"sync/atomic"
	"todo/tasks"
)

type InMemoryRepo struct {
	allTasks   map[int32]*tasks.Task
	lastTaskId atomic.Int32
}

func NewRepo() TaskRepo {
	return &InMemoryRepo{allTasks: make(map[int32]*tasks.Task)}
}

func (repo *InMemoryRepo) GetTask(taskId int32) (*tasks.Task, error) {
	return repo.allTasks[taskId], nil
}

func (repo *InMemoryRepo) GetAllTasks() ([]*tasks.Task, error) {
	result := make([]*tasks.Task, 0)
	for _, task := range repo.allTasks {
		result = append(result, task)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TaskId < result[j].TaskId
	})
	return result, nil
}

func (repo *InMemoryRepo) AddTask(title string, description string, dueDate string) (*tasks.Task, error) {
	taskId := repo.lastTaskId.Add(1)
	task := tasks.Task{TaskId: taskId, Title: title, Description: description, DueDate: dueDate, Status: tasks.NotStarted}
	repo.allTasks[taskId] = &task
	repo.lastTaskId.Store(taskId)

	allTasks, _ := repo.GetAllTasks()
	fmt.Println(allTasks)
	return &task, nil
}

func (repo InMemoryRepo) UpdateTask(
	taskId int32,
	newTitle string,
	newDesc string,
	newDueDate string,
	newStatus tasks.TaskStatus) (*tasks.Task, error) {
	taskToUpdate := repo.allTasks[taskId]
	if newTitle != "" {
		taskToUpdate.Title = newTitle
	}
	if newDesc != "" {
		taskToUpdate.Description = newDesc
	}
	if newDueDate != "" {
		taskToUpdate.DueDate = newDueDate
	}
	if newStatus != tasks.Unknown {
		taskToUpdate.Status = newStatus
	}
	return taskToUpdate, nil
}
