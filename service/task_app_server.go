package service

import (
	"context"
	"fmt"
	"todo/repository"
)

type TaskAppServer struct{}

func (server *TaskAppServer) mustEmbedUnimplementedTaskProtoServer() {
	//TODO implement me
	panic("implement me")
}

var repo repository.TaskRepo = repository.NewRepo()

func (server *TaskAppServer) Details(ctx context.Context, request *DetailsTaskRequest) (*DetailsTaskResponse, error) {
	task, err := repo.GetTask(request.TaskId)
	if err != nil {
		return nil, err
	}
	grpcTask, gerr := ToGrpcTask(task)
	if gerr != nil {
		return nil, gerr
	}
	return &DetailsTaskResponse{Task: grpcTask}, nil
}

func (server *TaskAppServer) List(ctx context.Context, request *ListTaskRequest) (*ListTaskResponse, error) {
	tasks, err := repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	result := make([]*Task, 0)
	for _, task := range tasks {
		grpcTask, gerr := ToGrpcTask(task)
		if gerr != nil {
			return nil, gerr
		}
		result = append(result, grpcTask)
	}
	fmt.Println(result)
	return &ListTaskResponse{Tasks: result}, nil
}

func (server *TaskAppServer) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	task, err := repo.AddTask(req.Title, req.GetDescription(), req.GetDueDate())
	if err != nil {
		return nil, err
	}
	fmt.Println("New task: " + task.String())

	grpcTask, gerr := ToGrpcTask(task)
	if gerr != nil {
		return nil, gerr
	}
	return &CreateTaskResponse{Task: grpcTask, TaskId: task.TaskId}, nil
}

func (server *TaskAppServer) Update(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	task, err := repo.UpdateTask(
		req.TaskId, req.GetNewTitle(), req.GetNewDescription(), req.GetNewDueDate(), FromGrpcStatus(req.GetNewStatus()))
	if err != nil {
		return nil, err
	}
	fmt.Println("New task: " + task.String())

	grpcTask, gerr := ToGrpcTask(task)
	if gerr != nil {
		return nil, gerr
	}
	return &UpdateTaskResponse{Task: grpcTask}, nil
}
