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

	// root
	e.GET("/", controllers.GetUsers)

	// schedules
	e.GET("/schedules/:id/now", dummy)
	e.GET("/schedules/:id/:datetime", dummy)

	// api
	api := e.Group("/api")

	// users
	api.GET("/users", controllers.GetUsers)
	api.POST("/users", controllers.PostUser)
	api.DELETE("/users/:id", controllers.DeleteUser) // unused in production service
	api.PUT("/users/:id/enable", dummy)
	api.PUT("/users/:id/disable", dummy)
	api.GET("/users/enabled", controllers.GetEnabledUsers)
	api.GET("/users/disabled", controllers.GetDisabledUsers)
	api.GET("/users/detail/:id", dummy)
	api.PATCH("/users/detail/:id", dummy)

	// shifts
	api.GET("/shifts", controllers.GetShifts)
	api.POST("/shifts", controllers.PostShift)
	api.GET("/shifts/:id", controllers.GetShift)
	api.PATCH("/shifts/:id", controllers.PatchShift)
	api.DELETE("/shifts/:id", controllers.DeleteShift)

	// rotations
	api.GET("/rotations", controllers.GetRotations)
	api.POST("/rotations", controllers.PostRotation)
	api.GET("/rotations/:rotation_id", controllers.GetRotation)
	api.PATCH("/rotations/:rotation_id", dummy)
	//api.DELETE("/rotations/:rotation_id", controllers.DeleteRotation)
	api.GET("/rotations/detail/:rotation_id/:order_id", controllers.GetRotationDetail)
	api.PUT("/rotations/detail/:rotation_id/:order_id", controllers.PutRotationDetail)

	// schedules
	api.GET("/schedules/:rotation_id/now", controllers.GetScheduleNow)
	//api.GET("/schedules/:rotation_id/:datetime", dummy)

	e.Logger.Fatal(e.Start(":" + port))
}

func dummy(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
