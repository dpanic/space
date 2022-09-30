package logic

import (
	"errors"
	"space/backend/models"
	"time"
)

func Update(currentProject, newProject *models.Project) (updatedProject *models.Project, err error) {

	// prevent concurrency issues
	if currentProject.Revision != newProject.Revision-1 {
		err := errors.New("project revision should be incremented by 1")
		return nil, err
	}

	updatedProject = currentProject
	updatedProject.Data = newProject.Data
	updatedProject.Revision++
	updatedProject.UpdatedAt = time.Now()

	updatedProject = currentProject

	return
}
