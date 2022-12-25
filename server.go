package main

import (
	"fmt"
	"net/http"

	"task/controller"
	"task/repository"
	router "task/router"
	"task/service"
)

var (
	taskRepository repository.TaskRepository = repository.NewFirestoreRepository()
	taskService    service.TaskService       = service.NewTaskService(taskRepository)
	taskControler  controller.TaskController = controller.NewTaskController(taskService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":9000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.GET("/task", taskControler.GetTaskList)
	httpRouter.POST("/task", taskControler.AddTask)
	httpRouter.SERVE(port)
}
