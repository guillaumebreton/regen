package loader

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Loader struct {
	loadedFiles map[string]*Resume
}

func NewLoader() *Loader {
	return &Loader{
		loadedFiles: make(map[string]*Resume),
	}
}

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
