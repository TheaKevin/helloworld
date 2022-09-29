package api

import (
	"github.com/TheaKevin/helloworld/todos"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
	}))

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	s.Router.DELETE("/delete/:taskId", todosHandler.DeleteTodo)
	s.Router.PUT("/changeDone/:taskId", todosHandler.ChangeDoneTodo)
}
