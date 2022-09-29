package models

//go:generate easytags $GOFILE json:snake
type GeoJSONFeatureCollection struct {
	Type     string         `json:"type"`
	Features GeoJSONFeature `json:"features"`
}

type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
}
