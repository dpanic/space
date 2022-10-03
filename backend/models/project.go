package models

import (
	"encoding/json"
	"fmt"
	"time"
)

//go:generate easytags $GOFILE json:snake
type Project struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Data      *Data     `json:"data,omitempty"`
	Revision  int       `json:"revision"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// this is used to mark if error in reading project from disk
	// no real use in system
	Error error `json:"error,omitempty"`

	_ struct{} `type:"structure"`
}

type Data struct {
	BuildingLimits *GeoJSONFeatureCollection `json:"building_limits"`
	HeighPlateaus  *GeoJSONFeatureCollection `json:"heigh_plateaus"`
	BuildingSplits *GeoJSONFeatureCollection `json:"building_splits"`
}

var (
	properties = map[string]map[string]interface{}{
		"BuildingLimits": {
			"name":           "BuildingLimits",
			"stroke":         "#000000",
			"stroke-width":   2,
			"stroke-opacity": 1,
			"fill":           "#555555",
			"fill-opacity":   0.9,
			"height":         10,
		},

		"HeighPlateaus": {
			"name":           "HeighPlateaus",
			"elevation":      0,
			"stroke":         "#000000",
			"stroke-width":   2,
			"stroke-opacity": 1,
			"fill":           "#00FF00",
			"fill-opacity":   0.9,
			"height":         10,
		},

		"BuildingSplits": {
			"name":           "BuildingSplits",
			"elevation":      0,
			"stroke":         "#000000",
			"stroke-width":   2,
			"stroke-opacity": 1,
			"fill":           "#0000FF",
			"fill-opacity":   0.9,
			"height":         10,
		},
	}
)

func GetProperty(name string) (out map[string]interface{}) {
	out = make(map[string]interface{})

	if value, ok := properties[name]; ok {
		for key, val := range value {
			out[key] = val
		}
	}

	return
}

func (data *Data) Populate() {
	data.BuildingLimits.Type = "FeatureCollection"

	if data.BuildingLimits != nil {
		for i := range data.BuildingLimits.Features {
			properties := GetProperty("BuildingLimits")
			properties["name"] = "BuildingLimits"

			data.BuildingLimits.Features[i] = &GeoJSONFeature{
				Type:       "Feature",
				Properties: properties,
				Geometry: GeoJSONGeometry{
					Type:        "Polygon",
					Coordinates: data.BuildingLimits.Features[i].Geometry.Coordinates,
				},
			}
		}
	}

	if data.HeighPlateaus != nil {
		data.HeighPlateaus.Type = "FeatureCollection"

		for i := range data.HeighPlateaus.Features {
			properties := GetProperty("HeighPlateaus")
			properties["elevation"] = data.HeighPlateaus.Features[i].Properties["elevation"]
			properties["name"] = "HeighPlateaus"

			data.HeighPlateaus.Features[i] = &GeoJSONFeature{
				Type:       "Feature",
				Properties: properties,
				Geometry: GeoJSONGeometry{
					Type:        "Polygon",
					Coordinates: data.HeighPlateaus.Features[i].Geometry.Coordinates,
				},
			}
		}
	}

	if data.BuildingSplits != nil {
		data.BuildingSplits.Type = "FeatureCollection"

		itMatched := 0
		itUnmatched := 0
		for i := range data.BuildingSplits.Features {
			properties := GetProperty("BuildingSplits")
			properties["status"] = data.BuildingSplits.Features[i].Properties["status"]

			var color string
			if properties["status"] == "matched" {
				color = genColor(2, 0, (itMatched+1)*40)
				itMatched++
			} else {
				color = genColor(0, 0, (itUnmatched+1)*40)
				itUnmatched++
			}

			properties["name"] = "BuildingSplits"
			properties["elevation"] = data.BuildingSplits.Features[i].Properties["elevation"]
			properties["type"] = data.BuildingSplits.Features[i].Properties["type"]
			properties["fill"] = color
			properties["stroke"] = color

			data.BuildingSplits.Features[i] = &GeoJSONFeature{
				Type:       "Feature",
				Properties: properties,
				Geometry: GeoJSONGeometry{
					Type:        "Polygon",
					Coordinates: data.BuildingSplits.Features[i].Geometry.Coordinates,
				},
			}
		}
	}
}

func genColor(index, start, step int) string {
	colors := []int{
		0,
		0,
		0,
	}

	colors[index] += start + step
	if colors[index] > 255 {
		colors[index] = 255
	}

	return fmt.Sprintf("#%02X%02X%02X", colors[0], colors[1], colors[2])
}

func (data *Data) Draw() (raw []byte) {
	raw, _ = json.MarshalIndent(data.BuildingSplits, "", "\t")
	return
}
