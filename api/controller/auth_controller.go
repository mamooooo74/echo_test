package controller

import (
	"fmt"
	"net/http"
	"test_api/api/model"
	"test_api/util"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type IAuthController interface {
	Login(c echo.Context) error
}

type authController struct{}

func NewAuthController() IAuthController {
	return authController{}
}

func (a authController) Login(c echo.Context) error {

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"statusCode": 1,
			"message":    "ログインできませんでした",
		})
	}
	pass := user.Password
	db := util.NewDB()
	defer util.CloseDB(db)
	db.Where("email = ?", user.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"statusCode": 1,
			"message":    "ログインできませんでした",
		})
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24,
		HttpOnly: true,
	}
	sess.Values["id"] = user.ID
	sess.Save(c.Request(), c.Response())
	fmt.Println("hello")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"statusCode": 1,
		"message":    "ログインに成功しました。",
	})
}
