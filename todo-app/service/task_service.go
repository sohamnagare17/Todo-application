package service

import (
	"todo-app/model"
	"todo-app/repository"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(title, description, status string) (model.Task, error)
	GetTask(id string) (model.Task, error)
	GetAllTasks() ([]model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func (s *taskService) NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) CreateTask(title, description, status string) (model.Task, error) {
	task := model.Task{
		Title:       title,
		Description: description,
		Status:      status,
		ID:          uuid.NewString(),
	}
	return s.repo.Create(task)
}

func (s *taskService) GetTask(id string) (model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) UpdateTask(task model.Task) (model.Task, error) {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
