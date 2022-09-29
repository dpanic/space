package storages

import (
	"os"
	"space/backend/models"
	"sync"
)

type Disk struct {
	RootDir string
	Mutex   *sync.Mutex
}

func NewDisk(rootDir string) *Disk {
	return &Disk{
		RootDir: rootDir,
		Mutex:   &sync.Mutex{},
	}
}

func (d *Disk) Create(name string) (id string, err error) {

	return
}

func (d *Disk) Read(id string) (data *models.Data, err error) {
	return
}

func (d *Disk) List() (objects []*models.Data, err error) {
	objects = make([]*models.Data, 0)

	return
}

func (d *Disk) Update(id string, data *models.Data) (err error) {
	return
}

func (d *Disk) Delete(id string) (err error) {
	return
}

func init() {
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		rootDir = "/tmp"
	}

	instances["disk"] = NewDisk(rootDir)

	os.MkdirAll(rootDir, 0755)
}
