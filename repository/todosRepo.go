package repository

import "final-project/model"

type TodosRepo interface {
	Create(param model.Todo) (model.Todo, error)
	Get() ([]model.Todo, error)
	GetById(id int64) (model.Todo, error)
	Update(param model.Todo) (model.Todo, error)
	Delete(id int64) error
}
