package handlers

import (
	"net/http"
	"yandex_go/cmd/models"
	"yandex_go/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}