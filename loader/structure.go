package loader

import "github.com/imdario/mergo"


// Information represents general information
type Information struct {
	Firstname string
	Lastname  string
	Phone     string
}

// Resume represents a resume structure
type Resume struct {
	Inherit     string
	Information Information
}

// Merge merges two resume
func (r *Resume) Merge(other *Resume) error {
	return mergo.Merge(r, other)
}
