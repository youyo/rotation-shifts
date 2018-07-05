package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetRotationDetail(c echo.Context) error {
	m := models.NewRotationDetail()
	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	rotation, err := m.GetRotationDetail(rotationId, orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(rotation, nil))
}

func PutRotationDetail(c echo.Context) error {
	m := models.NewRotationDetail()
	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	if err := m.PutRotationDetail(rotationId, orderId, m.ShiftId); err != nil {
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	return c.JSON(http.StatusCreated, response("created", nil))
}

func GetRotationsDetail(c echo.Context) error {
	m := models.NewRotationsDetail()
	users, err := m.GetRotationsDetail()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, nil))
}
