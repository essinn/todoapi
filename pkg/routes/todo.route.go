package routes

import (
	"todoapi/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine) {
	router.POST("/todos", controllers.NewTodo)
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todo/:todoId", controllers.GetTodo)
	router.PATCH("/todo/:todoId", controllers.UpdateTodo)
}
