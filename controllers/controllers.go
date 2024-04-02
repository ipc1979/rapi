package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"vemo/model"
	"vemo/repositories"
	"vemo/services"
)

func GetTasks(id int64, body []byte) (int, interface{}) {
	repository := repositories.GetTaskRepository()
	if id != 0 {
		task, err := services.GetTask(repository, id)
		if err != nil {
			return http.StatusNotFound, getErrorMessage(err)
		}
		return http.StatusOK, task
	} else {
		tasks := services.GetTasks(repository)
		return http.StatusOK, tasks
	}
}

func CreateTask(id int64, body []byte) (int, interface{}) {
	repository := repositories.GetTaskRepository()
	task, err := extractBody(body)
	if err == nil {
		return http.StatusCreated, services.CreateTask(repository, task)
	}
	return http.StatusBadRequest, getErrorMessage(err)
}

func UpdateTask(id int64, body []byte) (int, interface{}) {
	repository := repositories.GetTaskRepository()
	task, err := extractBody(body)
	if err != nil {
		return http.StatusBadRequest, getErrorMessage(err)
	}
	task, err = services.UpdateTask(repository, id, task)
	if err != nil {
		return http.StatusNotFound, getErrorMessage(err)
	}
	return http.StatusNoContent, task
}

func DeleteTask(id int64, body []byte) (int, interface{}) {
	repository := repositories.GetTaskRepository()
	id, err := services.DeleteTask(repository, id)
	if err != nil {
		return http.StatusNotFound, getErrorMessage(err)
	}
	return http.StatusNoContent, id
}

func getErrorMessage(err error) interface{} {
	message := make(map[string]string)
	message["message"] = err.Error()
	return message
}

func extractBody(body []byte) (model.Task, error) {
	task := model.Task{}
	err := json.Unmarshal(body, &task)
	var message string
	if task.Name == "" {
		message = message + "name is required, "
		err = errors.New(message)
	}
	if task.Description == "" {
		message = message + "description is required, "
		err = errors.New(message)
	}

	return task, err
}
