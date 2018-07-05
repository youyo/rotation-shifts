package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/youyo/rotation-shifts/models"
)

/*
func GetUser(c echo.Context) error {
	m := models.NewUser()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	user, err := m.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(user, nil))
}
*/

func PostUser(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.PostUser()")
	c.Echo().Logger.Debug("Called models.NewUser()")
	m := models.NewUser()
	c.Echo().Logger.Debug("Called models.PostUser()")
	userId, err := m.PostUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	c.Echo().Logger.Debug("Called models.NewUserDetail()")
	mud := models.NewUserDetail()
	if err := c.Bind(mud); err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	c.Echo().Logger.Debug("Called models.PostUserDetail()")
	if err := mud.PostUserDetail(userId, mud.Name); err != nil {
		c.Echo().Logger.Debug("Called models.DeleteUser()")
		m.DeleteUser(userId)
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}

	c.Echo().Logger.Debug("Called models.NewEnabledUser()")
	eu := models.NewEnabledUser()
	c.Echo().Logger.Debug("Called models.PostEnabledUser()")
	if err := eu.PostEnabledUser(userId); err != nil {
		c.Echo().Logger.Error("Failed to call func: models.PostEnabledUser(). error: %v", err)
		mud.DeleteUserDetail(userId)
		m.DeleteUser(userId)
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	if mud.Email != "" {
		c.Echo().Logger.Debug("Called models.UpdateUserDetailEmail()")
		if err := mud.UpdateUserDetailEmail(userId, mud.Email); err != nil {
			c.Echo().Logger.Error("Failed to call func: models.UpdateUserDetailEmail(). error: %v", err)
		}
	}
	if mud.PhoneNumber != "" {
		c.Echo().Logger.Debug("Called models.UpdateUserDetailPhoneNumber()")
		if err := mud.UpdateUserDetailPhoneNumber(userId, mud.PhoneNumber); err != nil {
			c.Echo().Logger.Error("Failed to call func: models.UpdateUserDetailPhoneNumber(). error: %v", err)
		}
	}
	return c.JSON(http.StatusOK, response(userId, nil))
}

func DeleteUser(c echo.Context) error {
	c.Echo().Logger.Debug("Called controllers.DeleteUser()")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	meu := models.NewEnabledUser()
	if err := meu.DeleteEnabledUser(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	mdu := models.NewDisabledUser()
	if err := mdu.DeleteDisabledUser(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	mud := models.NewUserDetail()
	if err := mud.DeleteUserDetail(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, response(nil, err))
	}
	m := models.NewUser()
	if err := m.DeleteUser(userId); err != nil {

	}
	return nil
}

func GetUsers(c echo.Context) error {
	m := models.NewUsers()
	users, err := m.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response(nil, err))
	}
	return c.JSON(http.StatusOK, response(users, err))
}
