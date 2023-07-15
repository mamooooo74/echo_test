package router

import (
	"test_api/api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	t := e.Group("/tasks")
	tc := controller.NewTaskController()
	t.GET("/", tc.GetAllTask)
	t.GET("/:taskId", tc.GetTaskById)
	return e
}
