package sql

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"todo/repository"
	"todo/tasks"
)

type SqlRepo struct {
	db *gorm.DB
}

func NewRepo() repository.TaskRepo {
	db, err := gorm.Open(sqlite.Open("taskrepo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating table")
	merr := db.AutoMigrate(&Task{})
	if merr != nil {
		return nil
	}
	return SqlRepo{db: db}
}

func (repo SqlRepo) GetTask(taskId int32) (*tasks.Task, error) {
	var requestedTask Task
	repo.db.Where("ID = ?", taskId).First(&requestedTask)

	return requestedTask.fromSql()
}

func (repo SqlRepo) GetAllTasks() ([]*tasks.Task, error) {
	var allTasks []Task
	repo.db.Find(&allTasks)

	fmt.Printf("Number of tasks: %d\n", len(allTasks))
	result := make([]*tasks.Task, 0)
	for _, task := range allTasks {
		fmt.Printf("%d. %s %s %s\n", task.ID, task.Title, task.Description, task.DueDate)
		t, err := task.fromSql()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%d. %s\n", t.TaskId, t.Title)
		result = append(result, t)
	}
	return result, nil
}

func (repo SqlRepo) AddTask(title string, description string, dueDate string) (*tasks.Task, error) {
	newTask := Task{Title: title, Description: description, DueDate: dueDate, Status: ""}
	repo.db.Create(&newTask)
	return newTask.fromSql()
}

func (repo SqlRepo) UpdateTask(
	taskId int32,
	newTitle string,
	newDesc string,
	newDueDate string,
	newStatus tasks.TaskStatus) (*tasks.Task, error) {
	var taskToUpdate Task
	repo.db.Where("ID = ?", taskId).First(&taskToUpdate)
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
		taskToUpdate.Status = newStatus.String()
	}
	return taskToUpdate.fromSql()
}
