package routes

import (
	"final-project/docs"
	"final-project/http/controllers"
	"final-project/repository"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(repo repository.TodosRepo) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "Todos API Swagger"
	docs.SwaggerInfo.Description = "Todos API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}

	ctrl := controllers.NewTodoController(repo)

	todos := router.Group("/todos")
	{
		todos.POST("/", ctrl.CreateTodo)
		todos.GET("/", ctrl.GetAllTodo)
		todos.GET("/:id", ctrl.GetTodoById)
		todos.PUT("/", ctrl.UpdateTodo)
		todos.DELETE("/", ctrl.DeleteTodo)
	}

	return router
}
