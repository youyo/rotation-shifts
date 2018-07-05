package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetDisabledUser(c echo.Context) error {
	m := models.NewDisabledUser()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	user, err := m.GetDisabledUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(user, nil))
}

func GetDisabledUsers(c echo.Context) error {
	m := models.NewDisabledUsers()
	users, err := m.GetDisabledUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, nil))
}
