package xray

import (
	"reflect"
	"strings"
)

// Name returns the name of an object
func Name(obj interface{}) string {
	if obj == nil {
		return "<nil>"
	}
	name := "<unnamed>"
	if t := reflect.TypeOf(obj); t != nil {
		if t.Kind() == reflect.Ptr {
			name = t.Elem().Name()
		} else {
			name = t.Name()
		}
	}
	return strings.ToLower(name)
}
