package router

import (
	"net/http"
	"test_api/api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
	}))

	t := e.Group("/tasks")
	tc := controller.NewTaskController()
	t.GET("/", tc.GetAllTask)
	t.GET("/:taskId", tc.GetTaskById)
	return e
}
