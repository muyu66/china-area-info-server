package main

type AreaDTO struct {
	ProvinceId string `json:"province_id"`
	CityId     string `json:"city_id"`
	DistrictId string `json:"district_id"`
}

type DistrictDTO struct {
	DistrictId string `json:"district_id"`
}
