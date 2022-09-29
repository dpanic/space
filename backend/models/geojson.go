package models

//go:generate easytags $GOFILE json:camel
type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   GeoJSONGeometry        `json:"geometry"`
}

type GeoJSONGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}
