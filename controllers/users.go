package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetUser(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetUser()")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetUser(userId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func GetUsers(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetUsers()")
	statusCode, resp, err := services.GetUsers()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func PostUser(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostUser()")
	m := queries.NewUser()
	if err := c.Bind(m); err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.CreateUser(m.Name, m.PhoneNumber)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func DeleteUser(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteUser()")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteUser(userId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}
