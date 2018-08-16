package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models/queries"
	"github.com/youyo/rotation-shifts/models/services"
)

func GetOverride(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetOverride()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.GetOverride(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func GetOverrides(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.GetOverrides()")
	statusCode, resp, err := services.GetOverrides()
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func PostOverride(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostOverride()")

	m := queries.NewOverride()
	if err := c.Bind(m); err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.CreateOverride(m.Name, m.Date, m.RotationId, m.UserId)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func PatchOverrideDetail(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PatchOverrideDetail()")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	m := queries.NewOverrideDetail()
	if err := c.Bind(m); err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.PatchOverrideDetail(id, m)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}

func DeleteOverride(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteOverride()")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}

	statusCode, resp, err := services.DeleteOverride(id)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(statusCode, response(nil, err))
	}
	return c.JSON(statusCode, response(resp, nil))
}
