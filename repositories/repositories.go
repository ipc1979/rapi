package repositories

import (
	"errors"
	"sync"
	"vemo/model"
)

var access = &sync.Mutex{}
var repository *TaskRepository

type TaskRepository struct {
	tasks  map[int64]model.Task
	nextId int64
}

func (t *TaskRepository) FindAll() []model.Task {
	access.Lock()
	tasks := make([]model.Task, 0)
	for _, task := range t.tasks {
		tasks = append(tasks, task)
	}
	defer access.Unlock()
	return tasks
}

func (t *TaskRepository) Create(task model.Task) model.Task {
	access.Lock()
	newTask := model.Task{
		Id:          t.nextId,
		Name:        task.Name,
		Description: task.Description,
	}
	t.tasks[t.nextId] = newTask
	t.nextId++
	access.Unlock()
	return newTask
}

func (t *TaskRepository) FindOne(id int64) (model.Task, error) {
	access.Lock()
	defer access.Unlock()
	return t.getTask(id)
}

func (t *TaskRepository) Update(id int64, task model.Task) (model.Task, error) {
	access.Lock()
	_, err := t.getTask(id)
	if err == nil {
		task.Id = id
		t.tasks[id] = task
	}
	access.Unlock()
	return task, err
}

func (t *TaskRepository) Delete(id int64) (int64, error) {
	access.Lock()
	_, err := t.getTask(id)
	if err == nil {
		delete(t.tasks, id)
	}
	access.Unlock()
	return id, err
}

func (t *TaskRepository) getTask(id int64) (model.Task, error) {
	var err error = nil
	task, exist := t.tasks[id]
	if !exist {
		err = errors.New("task not exist")
	}
	return task, err
}

func GetTaskRepository() *TaskRepository {
	access.Lock()
	if repository == nil {
		repository = &TaskRepository{
			tasks:  make(map[int64]model.Task, 0),
			nextId: 1,
		}
	}
	access.Unlock()
	return repository
}
