package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetAdditional(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetAdditional()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetAdditional(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func GetAdditionals(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetAdditionals()")
	statusCode, resp, err := services.GetAdditionals()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func PostAdditional(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostAdditional()")

	m := queries.NewAdditional()
	if err := c.Bind(m); err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.CreateAdditional(m.Name, m.Date, m.Hour, m.RotationId, m.UserId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func DeleteAdditional(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteAdditional()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteAdditional(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}
