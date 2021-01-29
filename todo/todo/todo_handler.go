package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitService(svc *Service) func() {
	return func() {
		//
	}
}

func createTodoHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		resp, err := svc.CreateTodo(req)
		if err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}

		c.Header("Location", resp.ID)
		c.JSON(http.StatusCreated, resp)
	}
}

func getTodoListHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := svc.GetTodoList()
		if err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func getTodoHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := svc.GetTodo()
		if err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func updateTodoHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		resp, err := svc.UpdateTodo(req)
		if err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func deleteTodoHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := svc.DeleteTodo(); err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}

		c.Status(http.StatusOK)
	}
}
