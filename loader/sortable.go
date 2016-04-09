package loader

import "time"

// Sortable defines a data sortable interface
type Sortable interface {
	Date() time.Time
}

//Sortables defines an array of sortable
type Sortables []Sortable

func (s Sortables) Len() int {
	return len(s)
}
func (s Sortables) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sortables) Less(i, j int) bool {
	return s[i].Date().After(s[j].Date())
}
