package main

type Unit struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	Id   uint32 `json:"id" gorm:"column:id"`
	Pid  uint32 `json:"pid" gorm:"column:pid"`
	Deep uint8  `json:"deep" gorm:"column:deep"`
	Name string `json:"name" gorm:"column:name"`
}

func (Area) TableName() string {
	return "area"
}
