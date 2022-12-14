package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"space/backend/models"
	"testing"
)

type test struct {
	BuildingLimits [][][]float64
	HeightPlateaus []map[string]interface{}
	Wants          wants
}
type wants struct {
	Matched   int
	Unmatched int
}

var tests = []test{
	{
		BuildingLimits: [][][]float64{
			{
				{0.0, 0.0}, {0.0, 60.0}, {60.0, 60.0}, {60.0, 0.0}, {0.0, 0.0},
			},
		},
		HeightPlateaus: []map[string]interface{}{
			{
				"elevation": float64(10),
				"coordinates": [][]float64{
					{0.0, 40.0}, {60.0, 40.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 40.0},
				},
			},
			{
				"elevation": float64(20),
				"coordinates": [][]float64{
					{0.0, 20.0}, {60.0, 20.0}, {60.0, 40.0}, {0.0, 40.0}, {0.0, 20.0},
				},
			},
			{
				"elevation": float64(30),
				"coordinates": [][]float64{
					{0.0, 0.0}, {60.0, 0.0}, {60.0, 20.0}, {0.0, 20.0}, {0.0, 0.0},
				},
			},
		},
		Wants: wants{
			Matched: 3,
		},
	},

	{
		BuildingLimits: [][][]float64{
			{
				{0.0, 40.0}, {60.0, 40.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 40.0},
			},
		},
		HeightPlateaus: []map[string]interface{}{
			{
				"elevation": float64(10),
				"coordinates": [][]float64{
					{0.0, 40.0}, {60.0, 40.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 40.0},
				},
			},
			{
				"elevation": float64(20),
				"coordinates": [][]float64{
					{0.0, 20.0}, {60.0, 20.0}, {60.0, 40.0}, {0.0, 40.0}, {0.0, 20.0},
				},
			},
			{
				"elevation": float64(30),
				"coordinates": [][]float64{
					{0.0, 0.0}, {60.0, 0.0}, {60.0, 20.0}, {0.0, 20.0}, {0.0, 0.0},
				},
			},
		},
		Wants: wants{
			Matched:   1,
			Unmatched: 2,
		},
	},

	{
		BuildingLimits: [][][]float64{
			{
				{0.0, 0.0}, {0.0, 60.0}, {60.0, 60.0}, {60.0, 0.0}, {0.0, 0.0},
			},
		},
		HeightPlateaus: []map[string]interface{}{
			{
				"elevation": float64(10),
				"coordinates": [][]float64{
					{0.0, 40.0}, {60.0, 40.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 40.0},
				},
			},
			// {
			// 	"elevation": float64(20),
			// 	"coordinates": [][]float64{
			// 		{0.0, 20.0}, {60.0, 20.0}, {60.0, 40.0}, {0.0, 40.0}, {0.0, 20.0},
			// 	},
			// },
			{
				"elevation": float64(30),
				"coordinates": [][]float64{
					{0.0, 0.0}, {60.0, 0.0}, {60.0, 20.0}, {0.0, 20.0}, {0.0, 0.0},
				},
			},
		},
		Wants: wants{
			Matched:   2,
			Unmatched: 1,
		},
	},

	{
		BuildingLimits: [][][]float64{
			{
				{0.0, 30.0}, {60.0, 30.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 30.0},
			},
			{
				{0.0, 0.0}, {60.0, 0.0}, {60.0, 30.0}, {0.0, 30.0}, {0.0, 0.0},
			},
		},
		HeightPlateaus: []map[string]interface{}{
			{
				"elevation": float64(10),
				"coordinates": [][]float64{
					{0.0, 40.0}, {60.0, 40.0}, {60.0, 60.0}, {0.0, 60.0}, {0.0, 40.0},
				},
			},
			{
				"elevation": float64(20),
				"coordinates": [][]float64{
					{0.0, 20.0}, {60.0, 20.0}, {60.0, 40.0}, {0.0, 40.0}, {0.0, 20.0},
				},
			},
			{
				"elevation": float64(30),
				"coordinates": [][]float64{
					{0.0, 0.0}, {60.0, 0.0}, {60.0, 20.0}, {0.0, 20.0}, {0.0, 0.0},
				},
			},
		},
		Wants: wants{
			Matched:   4,
			Unmatched: 0,
		},
	},
}

func TestSplits(t *testing.T) {
	totalTests := len(tests)
	for idx := range tests {
		project := getProject(idx)
		project.Data.Populate()

		results := Splits(&project)

		project.Data.BuildingSplits.Features = results
		project.Data.Populate()

		raw := project.Data.Draw()
		os.WriteFile(fmt.Sprintf("tmp/splits_%d.json", idx), raw, 0755)

		raw, _ = json.MarshalIndent(project, "", "\t")
		os.WriteFile(fmt.Sprintf("tmp/project_%d.json", idx), raw, 0755)

		var (
			matched   int
			unmatched int
		)
		for i := range project.Data.BuildingSplits.Features {
			status := project.Data.BuildingSplits.Features[i].Properties["status"]

			if status == "matched" {
				matched++
			} else {
				unmatched++
			}
		}

		if matched != tests[idx].Wants.Matched {
			t.Errorf("[ test %d / %d ] Wanted 'matched' == %d, got %d.\n", idx+1, totalTests, tests[idx].Wants.Matched, matched)
		}

		if unmatched != tests[idx].Wants.Unmatched {
			t.Errorf("[ test %d / %d ] Wanted 'unmatched' == %d, got %d.\n", idx+1, totalTests, tests[idx].Wants.Unmatched, unmatched)
		}
	}
}

func init() {
	os.MkdirAll("tmp", 0755)
}

func getProject(index int) (project models.Project) {
	BuildingLimits := tests[index].BuildingLimits
	HeightPlateaus := tests[index].HeightPlateaus

	project = models.Project{
		Data: &models.Data{
			BuildingLimits: &models.GeoJSONFeatureCollection{
				Features: []*models.GeoJSONFeature{},
			},
			HeightPlateaus: &models.GeoJSONFeatureCollection{
				Features: []*models.GeoJSONFeature{},
			},

			BuildingSplits: &models.GeoJSONFeatureCollection{
				Features: []*models.GeoJSONFeature{},
			},
		},
	}

	for _, bl := range BuildingLimits {
		var coordinates [][][]float64
		coordinates = append(coordinates, bl)

		obj := models.GeoJSONFeature{
			Properties: map[string]interface{}{},
			Geometry: models.GeoJSONGeometry{
				Coordinates: coordinates,
			},
		}

		project.Data.BuildingLimits.Features = append(project.Data.BuildingLimits.Features, &obj)
	}

	for _, hp := range HeightPlateaus {
		obj := models.GeoJSONFeature{
			Properties: map[string]interface{}{
				"elevation": hp["elevation"],
			},
			Geometry: models.GeoJSONGeometry{
				Coordinates: [][][]float64{
					hp["coordinates"].([][]float64),
				},
			},
		}

		project.Data.HeightPlateaus.Features = append(project.Data.HeightPlateaus.Features, &obj)
	}

	return
}
