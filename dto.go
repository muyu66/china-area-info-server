package main

type AreaDTO struct {
	ProvinceId int64 `json:"province_id"`
	CityId     int64 `json:"city_id"`
	DistrictId int64 `json:"district_id"`
}

type DistrictDTO struct {
	DistrictId int64 `json:"district_id"`
}

type BooleanDTO struct {
	Ok bool `json:"ok"`
}

type AreaNameDTO struct {
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
}
