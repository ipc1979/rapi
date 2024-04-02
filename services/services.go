package services

import (
	"vemo/model"
	"vemo/repositories"
)

func GetTasks(repository repositories.ITaskRepository) []model.Task {
	return repository.FindAll()
}

func CreateTask(repository repositories.ITaskRepository, task model.Task) model.Task {
	return repository.Create(task)
}

func GetTask(repository repositories.ITaskRepository, id int64) (model.Task, error) {
	return repository.FindOne(id)
}

func UpdateTask(repository repositories.ITaskRepository, id int64, task model.Task) (model.Task, error) {
	return repository.Update(id, task)
}

func DeleteTask(repository repositories.ITaskRepository, id int64) (int64, error) {
	return repository.Delete(id)
}
