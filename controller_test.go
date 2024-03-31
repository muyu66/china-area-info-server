package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Suite struct {
	suite.Suite
	e *echo.Echo
	h *controller
}

func (s *Suite) SetupSuite() {
	var areas = loadData()
	var dg = makeGraph(&areas)
	s.h = &controller{dg: dg}
}

func (s *Suite) SetupTest() {
	s.e = echo.New()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetProvinces() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	if assert.NoError(s.T(), s.h.getProvinces(c)) {
		var res []UnitNode
		assert.NoError(s.T(), json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), int64(11), res[0].ID())
	}
}

func (s *Suite) TestGetCities() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	c.SetPath("/provinces/:province_id/cities")
	c.SetParamNames("province_id")
	c.SetParamValues("34")

	if assert.NoError(s.T(), s.h.getCities(c)) {
		var res []UnitNode
		assert.NoError(s.T(), json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "合肥", res[0].Name)
		assert.Equal(s.T(), int64(3401), res[0].ID())
		assert.Equal(s.T(), uint8(1), res[0].Deep)
	}
}

func (s *Suite) TestGetDistricts() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	c.SetPath("/cities/:city_id/districts")
	c.SetParamNames("city_id")
	c.SetParamValues("3401")

	if assert.NoError(s.T(), s.h.getDistricts(c)) {
		var res []UnitNode
		assert.NoError(s.T(), json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "瑶海", res[0].Name)
		assert.Equal(s.T(), int64(340102), res[0].ID())
	}
}

func (s *Suite) TestGetCitiesWrongId() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	c.SetPath("/provinces/:province_id/cities")
	c.SetParamNames("province_id")
	c.SetParamValues("3401")

	assert.ErrorIs(s.T(), s.h.getCities(c), echo.ErrBadRequest)
}

func (s *Suite) TestGetCitiesNotFoundId() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	c.SetPath("/provinces/:province_id/cities")
	c.SetParamNames("province_id")
	c.SetParamValues("1")

	assert.ErrorIs(s.T(), s.h.getCities(c), echo.ErrNotFound)
}

func (s *Suite) TestValidate() {
	const body = `{"province_id":34,"city_id":3401,"district_id":340102}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	if assert.NoError(s.T(), s.h.validate(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "{\"ok\":true}\n", rec.Body.String())
	}
}

func (s *Suite) TestValidateError() {
	const body = `{"province_id":34,"city_id":3402,"district_id":340102}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	if assert.NoError(s.T(), s.h.validate(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "{\"ok\":false}\n", rec.Body.String())
	}
}

func (s *Suite) TestParse() {
	const body = `{"district_id":340102}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	if assert.NoError(s.T(), s.h.parse(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.Equal(s.T(), "{\"province_name\":\"安徽\",\"city_name\":\"合肥\",\"district_name\":\"瑶海\"}\n", rec.Body.String())
	}
}

func (s *Suite) TestParseWrongId() {
	const body = `{"district_id":34}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)

	assert.ErrorIs(s.T(), s.h.parse(c), echo.ErrBadRequest)
}
