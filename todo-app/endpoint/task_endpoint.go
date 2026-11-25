package endpoint

import (
	"context"
	"todo-app/model"

	"todo-app/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetTaskRequest struct {
	ID string `json:"id"`
}

func MakeCreateTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTaskRequest)

		task, err := svc.CreateTask(req.Title, req.Description, req.Status)
		if err != nil {
			return nil, err
		}
		return task, nil
	}
}

func MakeGetTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTaskRequest)
		task, err := svc.GetTask(req.ID)
		if err != nil {
			return nil, err
		}
		return task, nil
	}
}

func MakeGetAllTasksEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		tasks, err := svc.GetAllTasks()
		if err != nil {
			return nil, err
		}
		return tasks, nil
	}
}

func MakeUpdateTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Task)
		task, err := svc.UpdateTask(req)
		if err != nil {
			return nil, err
		}
		return task, nil
	}
}

func MakeDeleteTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTaskRequest)
		err := svc.DeleteTask(req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
