package controller

import (
	"net/http"
	"test_api/api/model"
	"test_api/util"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTask(c echo.Context) error
	GetTaskById(c echo.Context) error
}
type taskController struct {
}

func NewTaskController() ITaskController {
	return taskController{}
}

func (tc taskController) GetAllTask(c echo.Context) error {
	tasks := []model.Task{}
	db := util.NewDB()
	defer util.CloseDB(db)
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func (tc taskController) GetTaskById(c echo.Context) error {
	id := c.Param("taskId")
	task := model.Task{}
	db := util.NewDB()
	defer util.CloseDB(db)
	db.Where("id = ?", id).First(&task)
	return c.JSON(200, task)
}
