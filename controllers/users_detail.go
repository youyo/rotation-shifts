package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetUserDetail(c echo.Context) error {
	m := models.NewUserDetail()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	user, err := m.GetUserDetail(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(user, nil))
}

func GetUsersDetail(c echo.Context) error {
	m := models.NewUsersDetail()
	users, err := m.GetUsersDetail()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, nil))
}
