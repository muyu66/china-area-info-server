package main

import (
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo, ctl *controller) {
	e.POST("/parse", ctl.parse)

	e.POST("/validate", ctl.validate)

	e.GET("/provinces", ctl.getProvinces)
	e.GET("/provinces/:province_id/cities", ctl.getCities)
	e.GET("/cities/:city_id/districts", ctl.getDistricts)
}
