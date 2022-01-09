package api

import (
	"echo-gorm/db"
	"echo-gorm/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()

	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func AddUser(c echo.Context) error {
	db := db.DbManager()

	user := new(model.User)
	c.Bind(user)
	db.Create(user)

	return c.NoContent(http.StatusCreated)
}
