package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetSchedule(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetSchedule()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	dateTime := c.Param("datetime")

	s, err := services.NewSchedule(rotationId, dateTime)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}

	statusCode, resp, err := s.GetSchedule()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, err))
}
