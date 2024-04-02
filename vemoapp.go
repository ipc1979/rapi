package main

import (
	"net/http"
	"vemo/controllers"
	"vemo/restapi"
)

func main() {
	rApi := restapi.NewRestApi()
	rApi.AddAction(http.MethodGet, "/tasks", controllers.GetTasks)
	rApi.AddAction(http.MethodPost, "/tasks", controllers.CreateTask)
	rApi.AddAction(http.MethodPut, "/tasks", controllers.UpdateTask)
	rApi.AddAction(http.MethodDelete, "/tasks", controllers.DeleteTask)
	rApi.Run()
}
