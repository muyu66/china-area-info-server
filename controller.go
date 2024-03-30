package main

import (
	"github.com/labstack/echo/v4"
	"gonum.org/v1/gonum/graph/simple"
	"net/http"
	"strconv"
)

type (
	controller struct {
		dg *simple.DirectedGraph
	}
)

func (ctl *controller) getProvinces(c echo.Context) error {
	var nodes = getRootNodes(ctl.dg)

	return c.JSON(http.StatusOK, nodes)
}

func (ctl *controller) getCities(c echo.Context) error {
	provinceId, err := strconv.ParseInt(c.Param("province_id"), 10, 64)
	if err != nil || provinceId <= 0 {
		return echo.ErrBadRequest
	}

	var nodes = getNodesById(ctl.dg, provinceId)

	return c.JSON(http.StatusOK, nodes)
}

func (ctl *controller) getDistricts(c echo.Context) error {
	var res []Unit

	return c.JSON(http.StatusOK, res)
}

func (ctl *controller) validate(c echo.Context) error {
	var body AreaDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, body)
}

func (ctl *controller) validateDistrict(c echo.Context) error {
	var body DistrictDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, body)
}

func (ctl *controller) parse(c echo.Context) error {
	var body AreaDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, body)
}

func (ctl *controller) parseDistrict(c echo.Context) error {
	var body DistrictDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, body)
}
