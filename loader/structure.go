package loader

import (
	"fmt"
	"sort"
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
}

func mergeInformation(d, s Information) Information {
	r := Information{}
	r.Firstname = mergeString(d.Firstname, s.Firstname)
	r.Lastname = mergeString(d.Lastname, s.Lastname)
	r.Phone = mergeString(d.Phone, s.Phone)
	r.Email = mergeString(d.Email, s.Email)
	r.Twitter = mergeString(d.Twitter, s.Twitter)
	r.Website = mergeString(d.Website, s.Website)
	r.Title = mergeString(d.Title, s.Title)
	return r
}

func mergeString(s1, s2 string) string {
	if s1 == "" {
		fmt.Println(s1, s2)
		return s2
	}
	return s1
}

// Resume represents a resume structure
type Resume struct {
	Inherit     string
	Information Information
	Experiences map[string]Experience
	Educations  map[string]Education
	Projects    map[string]Project
}

//SortedExperiences sorts and returns the experiences list
func (r *Resume) SortedExperiences() []Experience {
	s := make(Experiences, len(r.Experiences))
	idx := 0
	for _, e := range r.Experiences {
		s[idx] = e
		idx++
	}
	sort.Sort(s)
	return s
}

//SortedEducations sorts and returns the educations list
func (r *Resume) SortedEducations() []Education {
	s := make(Educations, len(r.Educations))
	idx := 0
	for _, e := range r.Educations {
		s[idx] = e
		idx++
	}
	sort.Sort(s)
	return s
}

//SortedProjects sorts and returns the projects list
func (r *Resume) SortedProjects() []Project {
	s := make(Projects, len(r.Projects))
	idx := 0
	for _, e := range r.Projects {
		s[idx] = e
		idx++
	}
	sort.Sort(s)
	return s
}

// Merge merges two resume
func (r *Resume) Merge(other *Resume) error {
	var err error
	r.Information = mergeInformation(r.Information, other.Information)
	r.Experiences, err = MergeExperiences(r.Experiences, other.Experiences)
	if err != nil {
		return err
	}
	r.Educations, err = MergeEducations(r.Educations, other.Educations)
	if err != nil {
		return err
	}
	r.Projects, err = MergeProjects(r.Projects, other.Projects)
	if err != nil {
		return err
	}
	return nil
}
