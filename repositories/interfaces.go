package repositories

import "vemo/model"

type ITaskRepository interface {
	FindAll() []model.Task
	Create(model.Task) model.Task
	FindOne(int64) (model.Task, error)
	Delete(int64) (int64, error)
	Update(int64, model.Task) (model.Task, error)
}
