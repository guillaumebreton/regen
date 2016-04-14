package loader

import (
	"fmt"
	"github.com/imdario/mergo"
	"time"
)

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

// MergeExperiences merge two experiences map
func MergeExperiences(dst, src map[string]Experience) (map[string]Experience, error) {
	if src == nil {
		return nil, fmt.Errorf("Invalid experience merge")
	}
	dest := dst
	if dst == nil {
		dest = make(map[string]Experience)
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

// Experiences is a list of experiences
type Experiences []Experience

func (e Experiences) Len() int {
	return len(e)
}
func (e Experiences) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Experiences) Less(i, j int) bool {
	return e[i].Date().After(e[j].Date())
}
