package storages

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"space/backend/models"
)

type Disk struct {
	RootDir string
}

func NewDisk(rootDir string) *Disk {
	return &Disk{
		RootDir: rootDir,
	}
}

// Create create project on disk based on it's name
func (d *Disk) Create(project *models.Project) (err error) {
	projectPath := d.getProjectPath(project.Id)

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

func (d *Disk) getProjectPath(id string) (projectPath string) {
	projectPath = filepath.Join(d.RootDir, id)
	return
}

// Update project on disk based on ID
func (d *Disk) Update(id string, updatedProject *models.Project) (*models.Project, error) {
	fileLoc := filepath.Join(d.getProjectPath(id), "data")

	raw, _ := json.MarshalIndent(updatedProject, "", "\t")
	err := os.WriteFile(fileLoc, raw, 0755)

	return updatedProject, err
}

// Delete project on disk by ID
func (d *Disk) Delete(id string) (err error) {
	projectPath := filepath.Join(d.RootDir, id)
	err = os.RemoveAll(projectPath)

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
