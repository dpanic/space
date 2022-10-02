package logic

import (
	"errors"
	"fmt"
	"space/backend/models"
	"time"
)

func Update(currentProject, newProject *models.Project) (updatedProject *models.Project, sErrors []error) {
	// prevent concurrency issues
	if currentProject.Revision != newProject.Revision-1 {
		err := errors.New("project revision should be incremented by 1")
		sErrors = append(sErrors, err)
		return
	}

	// check if any feature is defined in building limits
	if len(newProject.Data.BuildingLimits.Features) == 0 {
		err := errors.New("buildingLimits aren't defined")
		sErrors = append(sErrors, err)
	}

	// check if any feature is defined in building limits
	if len(newProject.Data.BuildingLimits.Features) != 1 {
		err := errors.New("buildingLimits are different than 1, should be defined all in 1 feature")
		sErrors = append(sErrors, err)
	}

	// check if any feature is defined in heigh plateaus
	if len(newProject.Data.HeighPlateaus.Features) == 0 {
		err := errors.New("heighPlateaus aren't defined")
		sErrors = append(sErrors, err)
	}

	// check if elevations are defined in each heigh plateaus
	for i, feature := range newProject.Data.HeighPlateaus.Features {
		if _, ok := feature.Properties["elevation"]; !ok {
			err := fmt.Errorf("elevation not defined on heigh plateaus with index %d", i)
			sErrors = append(sErrors, err)
		}
	}

	if len(sErrors) > 0 {
		return
	}

	// perform building splits
	newProject.Data.BuildingSplits.Features = Splits(newProject)

	// add colors and names to geojson
	newProject.Data.Populate()

	// update data
	updatedProject = currentProject
	updatedProject.Data = newProject.Data
	updatedProject.Revision++
	updatedProject.UpdatedAt = time.Now()

	return
}
