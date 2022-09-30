package models

import "time"

//go:generate easytags $GOFILE json:snake
type Project struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Data      *Data     `json:"data,omitempty"`
	Revision  int       `json:"revision"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// this is used to mark if error in reading project from disk
	// no real use in system
	Error error `json:"error,omitempty"`

	_ struct{} `type:"structure"`
}

type ProjectShort struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Revision  int       `json:"revision"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Data struct {
	BuildingLimits GeoJSONFeatureCollection `json:"building_limits"`
	HeighPlateaus  GeoJSONFeatureCollection `json:"heigh_plateaus"`
}
