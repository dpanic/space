package logic

import "space/backend/models"

func List(results []*models.Project) {
	for _, r := range results {
		r.Data = nil
	}
}
