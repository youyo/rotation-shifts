package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetRotation(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetRotation()")
	rotationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetRotation(rotationId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func GetRotations(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetRotations()")

	statusCode, resp, err := services.GetRotations()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func PostRotation(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostRotation()")

	m := queries.NewRotation()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.PostRotation(m.Name, m.StartDate)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func DeleteRotation(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteRotation()")
	rotationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteRotation(rotationId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}
