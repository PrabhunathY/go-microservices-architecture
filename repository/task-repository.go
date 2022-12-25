package repository

import "task/model"

type TaskRepository interface {
	PostTask(task *model.Task) (*model.Task, error)
	GetTaskList() ([]model.Task, error)
}
