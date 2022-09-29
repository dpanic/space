package models

//go:generate easytags $GOFILE json:snake
type GeoJSONFeatureCollection struct {
	Type     string
	Features GeoJSONFeature
}

type GeoJSONFeature struct {
	Type       string
	Properties map[string]interface{}
	Geometry   struct {
		Type        string
		Coordinates [][]float64
	}
}
