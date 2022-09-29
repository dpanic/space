package storages

import (
	"encoding/json"
	"errors"
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

// Create create project on disk based on it's name
func (d *Disk) Create(name string) (id string, err error) {
	// uuid, err := crypto.UUID()
	id = crypto.SHA256(name)[0:8]

	projectPath := filepath.Join(d.RootDir, id)
	stat, _ := os.Stat(projectPath)
	if stat != nil {
		err = errors.New("project already exist")
		return
	}

	err = os.MkdirAll(projectPath, 0755)
	if err != nil {
		return
	}

	fileLoc := filepath.Join(projectPath, "data")

	project := models.Project{
		Id:       id,
		Name:     name,
		Revision: 1,
	}
	raw, _ := json.MarshalIndent(project, "", "\t")
	os.WriteFile(fileLoc, raw, 0755)

	return
}

// Read project from disk by it's ID
func (d *Disk) Read(id string) (project *models.Project, err error) {
	fileLoc := filepath.Join(d.RootDir, id, "data")
	raw, err := os.ReadFile(fileLoc)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &project)
	return
}

// List list all projects
func (d *Disk) List() (objects []*models.Project, err error) {
	objects = make([]*models.Project, 0)

	res, err := os.ReadDir(d.RootDir)
	if err != nil {
		return
	}

	for _, r := range res {
		id := r.Name()
		project, err := d.Read(id)
		if err != nil {
			project.Error = err
		}

		objects = append(objects, project)
	}

	return
}

// Update project on disk based on ID
func (d *Disk) Update(id string, data *models.Data) (err error) {
	return
}

// Delete project on disk by ID
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
