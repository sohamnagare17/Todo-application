package repository

import (
	"fmt"
	"todo-app/model"
)

type TaskRepository interface {
	Create(task model.Task) (model.Task, error)
	GetByID(id string) (model.Task, error)
	GetAll() ([]model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id string) error
}

type taskRepository struct {
	tasks map[string]model.Task
}

func NewTaskRepository() *taskRepository {
	return &taskRepository{
		tasks: make(map[string]model.Task),
	}
}

func (r *taskRepository) Create(task model.Task) (model.Task, error) {
	r.tasks[task.ID] = task
	return task, nil
}

func (r *taskRepository) GetByID(id string) (model.Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return model.Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}

func (r *taskRepository) GetAll() ([]model.Task, error) {
	tasks := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) Update(task model.Task) (model.Task, error) {
	_, exists := r.tasks[task.ID]
	if !exists {
		return model.Task{}, fmt.Errorf("task not found")
	}
	r.tasks[task.ID] = task
	return task, nil
}

func (r *taskRepository) Delete(id string) error {
	delete(r.tasks, id)
	return nil
}