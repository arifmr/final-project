package postgres

import (
	"final-project/model"
	"final-project/repository"

	"gorm.io/gorm"
)

type todos struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repository.TodosRepo {
	return &todos{db}
}

func (todo *todos) Create(param model.Todo) (model.Todo, error) {
	err := todo.db.Create(&param).Error
	return param, err
}

func (todo *todos) Get() ([]model.Todo, error) {
	var todoList []model.Todo
	err := todo.db.Find(&todoList).Error
	return todoList, err
}

func (todo *todos) GetById(id int64) (model.Todo, error) {
	var todoById model.Todo
	err := todo.db.First(&todoById, id).Error
	return todoById, err
}

func (todo *todos) Update(param model.Todo) (model.Todo, error) {
	err := todo.db.Save(&param).Error
	return param, err
}

func (todo *todos) Delete(id int64) error {
	return todo.db.Delete(&model.Todo{}, id).Error
}
