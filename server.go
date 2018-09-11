package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/youyo/rotation-shifts/controllers"
)

func main() {
	port := os.Getenv("PORT")
	e := echo.New()
	if os.Getenv("DEBUG") != "" {
		e.Debug = true
	}

	//e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// e.Use(middleware.CSRF())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH, echo.PUT},
	}))
	if os.Getenv("BASIC_AUTH") == "true" {
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == os.Getenv("BASIC_AUTH_USER") && password == os.Getenv("BASIC_AUTH_PASS") {
				return true, nil
			}
			return false, nil
		}))
	}

	// root
	e.GET("/", dummy)

	// api
	api := e.Group("/api")

	// users
	api.GET("/users", controllers.GetUsers)
	api.GET("/users/:id", controllers.GetUser)
	api.POST("/users", controllers.PostUser)
	api.DELETE("/users/:id", controllers.DeleteUser)
	api.PATCH("/users/:id", dummy)

	// shifts
	api.GET("/shifts", controllers.GetShifts)
	api.GET("/shifts/:id", controllers.GetShift)
	api.POST("/shifts", controllers.PostShift)
	api.PATCH("/shifts/:id", controllers.PatchShift)
	api.DELETE("/shifts/:id", controllers.DeleteShift)

	// rotations
	api.GET("/rotations", controllers.GetRotations)
	api.POST("/rotations", controllers.PostRotation)
	api.GET("/rotations/:id", controllers.GetRotation)
	api.PATCH("/rotations/:id", dummy)
	api.DELETE("/rotations/:id", controllers.DeleteRotation)

	api.GET("/rotation/details/:rotation_id", controllers.GetRotationDetails)
	api.GET("/rotation/details/:rotation_id/:order_id", controllers.GetRotationDetail)
	api.PUT("/rotation/details/:rotation_id", controllers.PutRotationDetail)
	api.DELETE("/rotation/details/:rotation_id", controllers.DeleteRotationDetail)
	api.DELETE("/rotation/details/:rotation_id/:order_id", controllers.DeleteRotationDetailWithOrderId)
	api.DELETE("/rotation/details/:rotation_id/:order_id/:shift_id", controllers.DeleteRotationDetailWithOrderIdAndShiftId)

	// overrides
	api.GET("/overrides", controllers.GetOverrides)
	api.GET("/overrides/:id", controllers.GetOverride)
	api.POST("/overrides", controllers.PostOverride)
	api.DELETE("/overrides/:id", controllers.DeleteOverride)

	/*
		// additional
		api.GET("/additionals", controllers.GetAdditionals)
		api.GET("/additionals/:id", controllers.GetAdditional)
		api.POST("/additionals", controllers.PostAdditional)
		api.DELETE("/additionals/:id", controllers.DeleteAdditional)

		// reduces
		api.GET("/reduces", controllers.GetReduces)
		api.GET("/reduces/:id", controllers.GetReduce)
		api.POST("/reduces", controllers.PostReduce)
		api.DELETE("/reduces/:id", controllers.DeleteReduce)
	*/

	// schedules
	api.GET("/schedules/:rotation_id/:datetime", controllers.GetSchedule)

	// calendars
	api.GET("/calendars/:rotation_id/:year-month", controllers.GetCalendar)

	e.Logger.Fatal(e.Start(":" + port))
}

func dummy(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
