package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetRotationDetail(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetRotationDetail()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetRotationDetail(rotationId, orderId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func GetRotationDetails(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetRotationDetails()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetRotationDetails(rotationId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func PutRotationDetail(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PutRotationDetail()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	m := queries.NewRotationDetail()
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.PutRotationDetail(rotationId, m.OrderId, m.ShiftId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func DeleteRotationDetail(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteRotationDetail()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteRotationDetail(rotationId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func DeleteRotationDetailWithOrderId(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteRotationDetailWithOrderId()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteRotationDetailWithOrderId(rotationId, orderId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}

func DeleteRotationDetailWithOrderIdAndShiftId(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteRotationDetailWithOrderIdAndShiftId()")

	rotationId, err := strconv.Atoi(c.Param("rotation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	shiftId, err := strconv.Atoi(c.Param("shift_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteRotationDetailWithOrderIdAndShiftId(rotationId, orderId, shiftId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}

	return c.JSON(statusCode, response(resp, nil))
}
