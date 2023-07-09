package tasks

import (
	"strconv"
	"time"
)

const layout string = "2006-01-02 15:04"

type Task struct {
	TaskId      int32
	Title       string
	Description string
	DueDate     string
	Status      TaskStatus
}

func Empty() *Task {
	return &Task{-1, "", "", "2023-06-01 12:00", NotStarted}
}

func NewTask(id int32, title string, desc string, duedateStr string, status TaskStatus) (*Task, error) {
	_, err := time.Parse(layout, duedateStr)
	if err != nil {
		return nil, err
	}
	return &Task{id, title, desc, duedateStr, status}, nil
}

func (t Task) String() string {
	return "TaskId: " + strconv.Itoa(int(t.TaskId)) + "\n" +
		"Title: " + t.Title + "\n" +
		"Description: " + t.Description + "\n" +
		"Due date: " + t.DueDate + "\n" +
		"Status: " + t.Status.String()
}
