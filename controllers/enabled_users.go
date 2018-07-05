package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

/*
func GetEnabledUser(c echo.Context) error {
	m := models.NewEnabledUser()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	user, err := m.GetEnabledUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(user, nil))
}
*/

func GetEnabledUsers(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetEnabledUsers()")
	m := models.NewEnabledUsers()
	users, err := m.GetEnabledUsers()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, nil))
}
