package loader

import "github.com/imdario/mergo"

type Information struct {
	Firstname string
	Lastname  string
	Phone     string
}

type Resume struct {
	Inherit     string
	Information Information
}

func (r *Resume) Merge(other *Resume) error {
	return mergo.Merge(r, other)
}

func (r *Resume) Print() {

}
