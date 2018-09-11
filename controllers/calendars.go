package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetCalendar(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetCalendar()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	yearMonth := c.Param("year-month")

	cal := services.NewCalendarEvents()

	statusCode, resp, err := cal.GetCalendarMonthly(rotationId, yearMonth)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, err))
}
