package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func PostShift(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostShift()")
	m := queries.NewShift()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.PostShift(m.Name, m.UserId)
	if err != nil {
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func PatchShift(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PatchShift()")

	m := queries.NewShiftDetail()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.PatchShift(shiftId, m)
	if err != nil {
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func DeleteShift(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteShift()")
	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteShift(shiftId)
	if err != nil {
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func GetShifts(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetShifts()")
	statusCode, resp, err := services.GetShifts()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func GetShift(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetShift()")
	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetShift(shiftId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}
