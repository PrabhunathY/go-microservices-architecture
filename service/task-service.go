package service

import (
	"errors"
	"math/rand"
	"task/model"
	"task/repository"
)

type TaskService interface {
	ValidateTask(task *model.Task) error
	CreateTask(task *model.Task) (*model.Task, error)
	GetAllTask() ([]model.Task, error)
}

type service struct{}

var (
	repo repository.TaskRepository
)

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	repo = taskRepository
	return &service{}
}

func (*service) ValidateTask(task *model.Task) error {
	if task == nil {
		err := errors.New("task is empy")
		return err
	}

	if task.Title == "" {
		err := errors.New("task title is empy")
		return err
	}
	return nil
}

func (*service) CreateTask(task *model.Task) (*model.Task, error) {
	task.ID = rand.Int63()
	return repo.PostTask(task)
}

func (*service) GetAllTask() ([]model.Task, error) {
	return repo.GetTaskList()
}
