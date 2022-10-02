package logic

import (
	"errors"
	"space/backend/models"
	"space/lib/crypto"
	"time"
)

func Create(name string) (project *models.Project, err error) {
	if name == "" {
		err = errors.New("project name can't be empty")
		return
	}

	id := crypto.SHA256(name)[0:8]

	project = &models.Project{
		Id:        id,
		Name:      name,
		Revision:  1,
		CreatedAt: time.Now(),
		Data: &models.Data{
			BuildingLimits: &models.GeoJSONFeatureCollection{
				Features: make([]*models.GeoJSONFeature, 0),
			},
			HeighPlateaus: &models.GeoJSONFeatureCollection{
				Features: make([]*models.GeoJSONFeature, 0),
			},
			BuildingSplits: &models.GeoJSONFeatureCollection{
				Features: make([]*models.GeoJSONFeature, 0),
			},
		},
	}

	return
}
