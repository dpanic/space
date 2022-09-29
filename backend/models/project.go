package models

//go:generate easytags $GOFILE json:snake
type Project struct {
	Id   string `json:"id"`
	Name string `json:"name" validate:"required"`
	Data Data   `json:"data"`

	_ struct{} `type:"structure"`
}

type Data struct {
	BuildingLimits GeoJSONFeatureCollection `json:"building_limits"`
	HeighPlateaus  GeoJSONFeatureCollection `json:"heigh_plateaus"`
}
