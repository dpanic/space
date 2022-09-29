package storages

import (
	"os"
	"path/filepath"
	"space/backend/models"
	"space/lib/crypto"
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
	// uuid, err := crypto.UUID()
	id = crypto.SHA256(name)[0:8]

	projectPath := filepath.Join(d.RootDir, id)
	os.MkdirAll(projectPath, 0755)

	return
}

func (d *Disk) Read(id string) (data *models.Project, err error) {
	return
}

func (d *Disk) List() (objects []*models.Project, err error) {
	objects = make([]*models.Project, 0)

	return
}

func (d *Disk) Update(id string, data *models.Data) (err error) {
	return
}

func (d *Disk) Delete(id string) (err error) {
	projectPath := filepath.Join(d.RootDir, id)
	err = os.Remove(projectPath)

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
