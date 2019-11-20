package xray

import (
	"github.com/oleiade/reflections"
)

// Tags lists the struct tag fields. obj can
// be a structure or pointer to structure.
func Tags(obj interface{}, tagname string) (map[string]string, error) {
	tags, err := reflections.Tags(obj, tagname)
	if err != nil {
		return tags, err
	}
	return tags, nil
}
