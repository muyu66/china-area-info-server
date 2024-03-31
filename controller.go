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
	sort(&nodes)
	return c.JSON(http.StatusOK, nodes)
}

func (ctl *controller) getCities(c echo.Context) error {
	provinceId, err := strconv.ParseInt(c.Param("province_id"), 10, 64)
	if err != nil || provinceId <= 0 {
		return echo.ErrBadRequest
	}
	if !checkExist(ctl.dg, provinceId) {
		return echo.ErrNotFound
	}
	if !checkDeep(ctl.dg, provinceId, 0) {
		return echo.ErrBadRequest
	}

	var nodes = getNodesById(ctl.dg, provinceId)
	sort(&nodes)
	return c.JSON(http.StatusOK, nodes)
}

func (ctl *controller) getDistricts(c echo.Context) error {
	cityId, err := strconv.ParseInt(c.Param("city_id"), 10, 64)
	if err != nil || cityId <= 0 {
		return echo.ErrBadRequest
	}
	if !checkExist(ctl.dg, cityId) {
		return echo.ErrNotFound
	}
	if !checkDeep(ctl.dg, cityId, 1) {
		return echo.ErrBadRequest
	}

	var nodes = getNodesById(ctl.dg, cityId)
	sort(&nodes)
	return c.JSON(http.StatusOK, nodes)
}

func (ctl *controller) validate(c echo.Context) error {
	var body AreaDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}
	if body.DistrictId <= 0 {
		return echo.ErrBadRequest
	}
	if !checkExist(ctl.dg, body.DistrictId) {
		return echo.ErrNotFound
	}
	if !checkDeep(ctl.dg, body.DistrictId, 2) {
		return echo.ErrBadRequest
	}

	m, err := getAllParentNodesById(ctl.dg, body.DistrictId)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, BooleanDTO{
		Ok: body.ProvinceId == m[0].ID() && body.CityId == m[1].ID(),
	})
}

func (ctl *controller) parse(c echo.Context) error {
	var body DistrictDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}
	if body.DistrictId <= 0 {
		return echo.ErrBadRequest
	}
	if !checkExist(ctl.dg, body.DistrictId) {
		return echo.ErrNotFound
	}
	if !checkDeep(ctl.dg, body.DistrictId, 2) {
		return echo.ErrBadRequest
	}

	m, err := getAllParentNodesById(ctl.dg, body.DistrictId)
	if err != nil {
		return echo.ErrBadRequest
	}

	var res = AreaNameDTO{
		ProvinceName: m[0].Name,
		CityName:     m[1].Name,
		DistrictName: m[2].Name,
	}

	return c.JSON(http.StatusOK, res)
}
