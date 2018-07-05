package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetShift(c echo.Context) error {
	m := models.NewShift()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	user, err := m.GetShift(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(user, nil))
}

func PostShift(c echo.Context) error {
	m := models.NewShift()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.PostShift(m.Name, m.UserId); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusCreated, response("created", nil))
}

func PatchShift(c echo.Context) error {
	m := models.NewShift()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.PatchShift(shiftId); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response("updated", nil))
}

func DeleteShift(c echo.Context) error {
	m := models.NewShift()
	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.DeleteShift(shiftId); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response("deleted", nil))
}

func GetShifts(c echo.Context) error {
	m := models.NewShifts()
	users, err := m.GetShifts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, err))
}
