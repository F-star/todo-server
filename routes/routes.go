package routes

import (
	todo "todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todos", todo.GetAllTodos)
		v1.POST("todo", todo.CreateATodo)
		v1.GET("todo/:id", todo.GetATodo)
		v1.DELETE("todo/:id", todo.DeleteTodo)
		v1.PUT("todo", todo.UpdateATodo)
	}
	return r
}
