package loader

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Loader load toml file and merge them
type Loader struct {
	loadedFiles map[string]*Resume
}

// NewLoader creates a new loader with an initialized map
func NewLoader() *Loader {
	return &Loader{
		loadedFiles: make(map[string]*Resume),
	}
}

// Load loads a toml data file and merege inherited data
func (g *Loader) Load(dir string, name string) (*Resume, error) {
	return g.load(dir, name, make([]string, 10))
}

// Load loads toml data file and merge into inherited data. It tracks loaded data
// to detect circular inheritance
func (g *Loader) load(dir string, name string, seen []string) (*Resume, error) {
	if contains(name, seen) {
		return nil, fmt.Errorf("Circular inheritance detected with file %s", name)
	}
	var resume *Resume
	resume = g.loadedFiles[name]
	if resume != nil {
		return resume, nil
	}
	//load the toml file
	filepath := filepath.Join(dir, name)
	tomlData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	// out the file content
	if _, err := toml.Decode(string(tomlData), &resume); err != nil {
		return nil, err
	}
	if resume.Inherit != "" {
		newSeen := append(seen, name)
		parent, err := g.load(dir, resume.Inherit, newSeen)
		if err != nil {
			return nil, err
		}
		resume.Merge(parent)
	}
	g.loadedFiles[name] = resume
	return resume, nil
}

func contains(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
