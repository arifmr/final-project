package controllers

import (
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	repo repository.TodosRepo
}

func NewTodoController(repo repository.TodosRepo) *controller {
	return &controller{repo}
}

// CreateTodo godoc
// @Summary Create Todo
// @Description create todo
// @ID create-todo
// @Accept  json
// @Produce  json
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} mode.ErrorResponse
// @Failure 500 {object} mode.ErrorResponse
// @Router /todos [post]
func (ctrl *controller) CreateTodo(c *gin.Context) {
	var param model.Todo
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ctrl.repo.Create(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

// GetAllTodo godoc
// @Summary Get All Todo
// @Description get all todo
// @ID get-all-todo
// @Accept  json
// @Produce  json
// @Success 200 {object} model.SuccessResponse
// @Failure 500 {object} mode.ErrorResponse
// @Router /todos [get]
func (ctrl *controller) GetAllTodo(c *gin.Context) {
	result, err := ctrl.repo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

// GetTodoById godoc
// @Summary Get Todo By Id
// @Description get todo by id
// @ID get-todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400, 404 {object} mode.ErrorResponse
// @Router /todos/{id} [get]
func (ctrl *controller) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, errGet := ctrl.repo.GetById(int64(idint))
	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

// UpdateTodo godoc
// @Summary Update Todo
// @Description update todo
// @ID update-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} mode.ErrorResponse
// @Failure 500 {object} mode.ErrorResponse
// @Router /todos/{id} [put]
func (ctrl *controller) UpdateTodo(c *gin.Context) {
	var param model.Todo
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ctrl.repo.Update(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

// DeleteTodo godoc
// @Summary Delete Todo
// @Description delete todo
// @ID delete-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} mode.ErrorResponse
// @Failure 500 {object} mode.ErrorResponse
// @Router /todos/{id} [delete]
func (ctrl *controller) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = ctrl.repo.Delete(int64(idint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
