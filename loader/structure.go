package loader

import (
	"github.com/imdario/mergo"
	"sort"
	"time"
)

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

// Experience define a resume experience part
type Experience struct {
	Start       string
	End         string
	Company     string
	Title       string
	Description string
}

//Date defines a sorting date
func (e Experience) Date() time.Time {
	t, err := time.Parse("2006-01", e.Start)
	if err != nil {
		return time.Now() // hiding the error ..
	}
	return t
}

// Education defines a resume education part
type Education struct {
	Year        int
	School      string
	Title       string
	Description string
}

// Date defines a sorting date
func (e Education) Date() time.Time {
	return time.Date(e.Year, 1, 1, 0, 0, 0, 0, time.UTC)
}

// Project defines a personal project
type Project struct {
	Year        int
	Name        string
	URL         string
	Description string
}

// Date defines the sorting date
func (p Project) Date() time.Time {
	return time.Date(p.Year, 1, 1, 0, 0, 0, 0, time.UTC)
}

// Resume represents a resume structure
type Resume struct {
	Inherit     string
	Information Information
	Experiences []Experience
	Educations  []Education
	Projects    []Project
}

//SortedExperiences sorts and returns the experiences list
func (r *Resume) SortedExperiences() []Experience {
	s := make(Sortables, len(r.Experiences))
	for k, v := range r.Experiences {
		s[k] = Sortable(v)
	}
	sort.Sort(s)
	return r.Experiences
}

//SortedEducations sorts and returns the educations list
func (r *Resume) SortedEducations() []Education {
	s := make(Sortables, len(r.Educations))
	for k, v := range r.Educations {
		s[k] = Sortable(v)
	}
	sort.Sort(s)
	return r.Educations
}

//SortedProjects sorts and returns the projects list
func (r *Resume) SortedProjects() []Project {
	s := make(Sortables, len(r.Projects))
	for k, v := range r.Projects {
		s[k] = Sortable(v)
	}
	sort.Sort(s)
	return r.Projects
}

// Merge merges two resume
func (r *Resume) Merge(other *Resume) error {
	mergo.Merge(r.Information, other.Information)
	return mergo.Merge(r, other)
}
