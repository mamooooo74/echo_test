package router

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"test_api/api/controller"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	ac := controller.NewAuthController()

	e.POST("/login", ac.Login)
	e.GET("/user", ac.GetUser)
	e.GET("logout", ac.Logout)

	t := e.Group("/tasks")
	tc := controller.NewTaskController()
	t.GET("/", tc.GetAllTask)
	t.GET("/:taskId", tc.GetTaskById)
	return e
}
