package models

//go:generate easytags $GOFILE json:camel
type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data Data
}

type Data struct {
	BuildingLimits GeoJSONFeature
	HeighPlateaus  GeoJSONFeature
}
