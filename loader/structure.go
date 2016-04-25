package loader

import (
	"sort"
	"strconv"
	"strings"
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

// Experience define a resume experience part
type Experience struct {
	Start       string
	End         string
	Company     string
	Title       string
	Description string
}

// DescriptionLines returns the description splitted by lines
func (e Experience) DescriptionLines() []string {
	d := strings.Split(strings.Trim(e.Description, "\n "), "\n")
	for k, v := range d {
		d[k] = strings.TrimSpace(v)
	}
	return d
}

// Education defines a resume education part
type Education struct {
	Year        int
	School      string
	Title       string
	Description string
}

// Project defines a personal project
type Project struct {
	Year        int
	Name        string
	URL         string
	Description string
}

// Resume represents a resume structure
type Resume struct {
	Information Information
	Experiences map[string]Experience
	Educations  map[string]Education
	Projects    map[string]Project
}

// SortedExperiences returns the sorted experiences
func (r *Resume) SortedExperiences() []Experience {
	ranks := make([]int, len(r.Experiences))
	i := 0
	for k := range r.Experiences {
		v, err := strconv.Atoi(k)
		if err != nil {
			panic("Key is not an int")
		}
		ranks[i] = v
		i++
	}
	//sort keys
	sort.Ints(ranks)
	result := make([]Experience, len(r.Experiences))
	for i, v := range ranks {
		sv := strconv.Itoa(v)
		result[i] = r.Experiences[sv]
	}
	return result
}

// SortedEducations returns the sorted education list
func (r *Resume) SortedEducations() []Education {
	ranks := make([]int, len(r.Educations))
	i := 0
	for k := range r.Educations {
		v, err := strconv.Atoi(k)
		if err != nil {
			panic("Key is not an int")
		}
		ranks[i] = v
		i++
	}
	//sort keys
	sort.Ints(ranks)
	result := make([]Education, len(r.Educations))
	for i, v := range ranks {
		sv := strconv.Itoa(v)
		result[i] = r.Educations[sv]
	}
	return result
}

// SortedProjects returns the sorted projects list
func (r *Resume) SortedProjects() []Project {
	ranks := make([]int, len(r.Projects))
	i := 0
	for k := range r.Projects {
		v, err := strconv.Atoi(k)
		if err != nil {
			panic("Key is not an int")
		}
		ranks[i] = v
		i++
	}
	//sort keys
	sort.Ints(ranks)
	result := make([]Project, len(r.Projects))
	for i, v := range ranks {
		sv := strconv.Itoa(v)
		result[i] = r.Projects[sv]
	}
	return result
}
