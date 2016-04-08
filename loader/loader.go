package loader

import (
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

// Load load toml data file and merge into inherited data
func (g *Loader) Load(dir string, name string) (*Resume, error) {
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
		parent, err := g.Load(dir, resume.Inherit)
		if err != nil {
			return nil, err
		}
		resume.Merge(parent)
	}
	g.loadedFiles[name] = resume
	return resume, nil
}
