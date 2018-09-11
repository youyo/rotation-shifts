package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetReduce(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetReduce()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetReduce(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func GetReduces(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetReduces()")
	statusCode, resp, err := services.GetReduces()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func PostReduce(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostReduce()")

	m := queries.NewReduce()
	if err := c.Bind(m); err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.CreateReduce(m.Name, m.Date, m.Hour, m.RotationId, m.UserId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func DeleteReduce(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteReduce()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteReduce(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}
