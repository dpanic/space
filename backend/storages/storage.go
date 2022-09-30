package storages

import "space/backend/models"

type Storage interface {
	Create(project *models.Project) error
	Delete(id string) error
	Read(id string) (*models.Project, error)
	Update(id string, project *models.Project) (*models.Project, error)
	List() ([]*models.Project, error)
}

var (
	instances = make(map[string]Storage, 0)
)

// Get gets desired storage by name
func Get(name string) (obj *Storage, found bool) {
	if val, ok := instances[name]; ok {
		obj = &val
		found = true
	}

	return
}
