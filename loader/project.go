package loader

import (
	"fmt"
	"github.com/imdario/mergo"
	"time"
)

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

// MergeProjects merge two projects map
func MergeProjects(dst, src map[string]Project) (map[string]Project, error) {
	if src == nil {
		return nil, fmt.Errorf("Invalid project merge")
	}
	dest := dst
	if dst == nil {
		dest = make(map[string]Project)
	}

	for k, v := range src {
		if dv, ok := dest[k]; !ok {
			dest[k] = v
		} else {
			if err := mergo.Merge(&dv, v); err != nil {
				return nil, err
			}
			dest[k] = dv
		}
	}
	return dest, nil
}

// Projects is a list of projects
type Projects []Project

func (p Projects) Len() int {
	return len(p)
}

func (p Projects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Projects) Less(i, j int) bool {
	return p[i].Date().After(p[j].Date())
}
