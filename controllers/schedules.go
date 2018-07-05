package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

func GetScheduleNow(c echo.Context) error {
	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	u, err := models.GetScheduleNow(rotationId)
	return c.JSON(http.StatusOK, response(u, err))
}
