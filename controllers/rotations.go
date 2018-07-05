package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetRotation(c echo.Context) error {
	m := models.NewRotation()
	id, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	rotation, err := m.GetRotation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(rotation, nil))
}

func PostRotation(c echo.Context) error {
	m := models.NewRotation()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.PostRotation(m.Name, m.StartDate); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusCreated, response("created", nil))
}

/*
func PatchRotation(c echo.Context) error {
	m := models.NewRotation()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	shiftId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.PatchRotation(shiftId); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response("updated", nil))
}
*/

func DeleteRotation(c echo.Context) error {
	m := models.NewRotation()
	id, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.DeleteRotation(id); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response("deleted", nil))
}

func GetRotations(c echo.Context) error {
	m := models.NewRotations()
	rotations, err := m.GetRotations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(rotations, err))
}
