package loader

import "github.com/imdario/mergo"


// Information represents general information
type Information struct {
	Firstname string
	Lastname  string
	Phone     string
  Email     string
  Twitter   string
  Website   string
  Title     string
  Skills    []string
}

type Experience struct {
  Start       string
  End         string
  Company     string
  Title       string
  Description string
}

type Education struct {
  Year        string
  School      string
  Title       string
  Description string
}

type Project struct {
  Name string
  Url string
  Description string
}

// Resume represents a resume structure
type Resume struct {
	Inherit     string
	Information Information
  Experiences []Experience
  Educations  []Education
  Projects    []Project
}

// Merge merges two resume
func (r *Resume) Merge(other *Resume) error {
	return mergo.Merge(r, other)
}
