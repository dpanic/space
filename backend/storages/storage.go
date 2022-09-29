package storages

import "space/backend/models"

type Storage interface {
	Create(name string) (id string, err error)
	Delete(id string) (err error)
	Read(id string) (object *models.Data, err error)
	Update(id string, data *models.Data) (err error)
	List() (objects []*models.Data, err error)
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
