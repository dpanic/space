package logic

import (
	"space/backend/models"
	"space/lib/logger"

	"github.com/engelsjk/polygol"
	"go.uber.org/zap"
)

func Splits(project *models.Project) (buildingSplits []*models.GeoJSONFeature) {
	buildingSplits = make([]*models.GeoJSONFeature, 0)

	buildingLimits := make([]*models.GeoJSONFeature, 0)
	buildingLimits = append(buildingLimits, project.Data.BuildingLimits.Features...)

	heightPlateaus := make([]*models.GeoJSONFeature, 0)
	heightPlateaus = append(heightPlateaus, project.Data.HeightPlateaus.Features...)

	for len(buildingLimits) > 0 {
		// pop first element of buildingLimits
		bl := buildingLimits[0]
		buildingLimits = buildingLimits[1:]

		var (
			isMatched bool
		)
		for i, hl := range heightPlateaus {
			b := polygol.Geom{
				bl.Geometry.Coordinates,
			}

			h := polygol.Geom{
				hl.Geometry.Coordinates,
			}
			intersection, _ := polygol.Intersection(b, h)
			differenceB, _ := polygol.Difference(b, h)
			differenceH, _ := polygol.Difference(h, b)

			// if there is intersection cut it and add to building splits
			if len(intersection) > 0 {
				log := logger.Log.WithOptions(zap.Fields(
					zap.Any("buildingLimits", bl.Geometry),
					zap.Any("heightPlateaus", hl.Geometry),
					zap.Any("difference", differenceB),
					zap.Any("intersection", intersection),
					zap.Float64("elevation", hl.Properties["elevation"].(float64)),
				))
				log.Info("matched building limit and heigh plateaus")

				obj := models.GeoJSONFeature{
					Properties: map[string]interface{}{
						"status":    "matched",
						"type":      "BuildingLimits",
						"elevation": hl.Properties["elevation"],
					},
					Geometry: models.GeoJSONGeometry{
						Coordinates: intersection[0],
					},
				}
				buildingSplits = append(buildingSplits, &obj)

				isMatched = true

				// remove heigh plateaus from pool
				heightPlateaus = append(heightPlateaus[:i], heightPlateaus[i+1:]...)

				// if there is non 100% match return heigh plateaus back to pool
				if len(differenceB) > 0 {
					log.Warn("returning back difference to building limits")

					obj := models.GeoJSONFeature{
						Properties: models.GetProperty("BuildingLimits"),
						Geometry: models.GeoJSONGeometry{
							Coordinates: differenceB[0],
						},
					}
					buildingLimits = append(buildingLimits, &obj)
				}

				if len(differenceH) > 0 {
					log.Warn("returning back difference to heigh plateaus")

					obj := models.GeoJSONFeature{
						Properties: map[string]interface{}{
							"elevation": hl.Properties["elevation"],
							"type":      "HeightPlateaus",
						},
						Geometry: models.GeoJSONGeometry{
							Coordinates: differenceH[0],
						},
					}
					heightPlateaus = append(heightPlateaus, &obj)
				}

				break
			}
		}

		// add bl to unmatched
		if !isMatched {
			logger.Log.Info("building limit isn't matched to any height plateaus",
				zap.Any("buildingLimits", bl.Geometry),
				zap.Any("buildingLimits", bl.Geometry),
			)

			bl.Properties["status"] = "unmatched"
			bl.Properties["type"] = "BuildingLimits"
			buildingSplits = append(buildingSplits, bl)
		}
	}

	for _, hl := range heightPlateaus {
		logger.Log.Info("heigh plateaus is unmatched",
			zap.Any("heightPlateaus", hl.Geometry),
		)

		hl.Properties["status"] = "unmatched"
		hl.Properties["type"] = "HeightPlateaus"
		buildingSplits = append(buildingSplits, hl)
	}

	return
}
