package main

import (
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo, ctl *controller) {
	e.POST("/parse", ctl.parse)
	e.POST("/parse_district", ctl.parseDistrict)

	e.POST("/validate", ctl.validate)
	e.POST("/validate_district", ctl.validateDistrict)

	e.GET("/provinces", ctl.getProvinces)
	e.GET("/provinces/:province_id/cities", ctl.getCities)
	e.GET("/citys/:city_id/districts", ctl.getDistricts)
}
