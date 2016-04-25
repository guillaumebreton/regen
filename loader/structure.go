package loader

import (
  "strconv"
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

// Experience define a resume experience part
type Experience struct {
	Start       string
	End         string
	Company     string
	Title       string
	Description string
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

func (r *Resume) SortedExperiences() []Experience {
  ranks := make([]int, len(r.Experiences))
  i := 0
  for k, _ := range r.Experiences {
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
  for i,v := range ranks {
    sv := strconv.Itoa(v)
    result[i] = r.Experiences[sv]
  }
  return result
}
func (r *Resume) SortedEducations() []Education {
  return nil
}
func (r *Resume) SortedProjets() []Project {
  return nil
}

