package loader

import (
	"fmt"
	"github.com/imdario/mergo"
	"time"
)

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

// MergeEducations merge two educations map
func MergeEducations(dst, src map[string]Education) (map[string]Education, error) {
	if src == nil {
		return nil, fmt.Errorf("Invalid education merge")
	}
	dest := dst
	if dst == nil {
		dest = make(map[string]Education)
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

// Educations is a list of educations
type Educations []Education

func (e Educations) Len() int {
	return len(e)
}

func (e Educations) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Educations) Less(i, j int) bool {
	return e[i].Date().After(e[j].Date())
}
