package controller

import (
	"encoding/json"
	"net/http"
	"task/errors"
	"task/model"
	"task/service"
)

type controller struct{}

var (
	taskService service.TaskService
)

type TaskController interface {
	GetTaskList(resp http.ResponseWriter, req *http.Request)
	AddTask(resp http.ResponseWriter, req *http.Request)
}

func NewTaskController(service service.TaskService) TaskController {
	taskService = service
	return &controller{}
}

func (*controller) GetTaskList(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	taskList, err := taskService.GetAllTask()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the task list"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(taskList)
}

func (*controller) AddTask(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var task model.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := taskService.ValidateTask(&task)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := taskService.CreateTask(&task)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the task"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
