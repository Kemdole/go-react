package todo

import "github.com/gin-gonic/gin"

func initService(svc *Service) func() {
	return func() {
		//
	}
}

func createTodoHandler(svc *Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		svc.CreateTodo()
	}
}

func getTodoListHandler(svc *Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		svc.GetTodoList()
	}
}

func getTodoHandler(svc *Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		svc.GetTodo()
	}
}

func updateTodoHandler(svc *Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		svc.UpdateTodo()
	}
}

func deleteTodoHandler(svc *Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		svc.DeleteTodo()
	}
}
