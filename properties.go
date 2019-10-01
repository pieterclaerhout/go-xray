package xray

import (
	"errors"
	"strings"

	"github.com/oleiade/reflections"
)

// Properties returns the struct fields names list. obj can whether
// be a structure or pointer to structure.
func Properties(obj interface{}) ([]string, error) {
	fields, err := reflections.Fields(obj)
	if err != nil {
		return nil, err
	}
	return fields, err
}

// PropertiesAsMap returns the properties of an item as a map
func PropertiesAsMap(obj interface{}) (map[string]interface{}, error) {

	properties := map[string]interface{}{}

	propertyNames, err := Properties(obj)
	if err != nil {
		return properties, err
	}

	for _, propertyName := range propertyNames {
		propertyValue, _ := Property(obj, propertyName)
		properties[propertyName] = propertyValue
	}

	return properties, nil

}

// Property returns the specific property of obj
func Property(obj interface{}, fieldName string) (interface{}, error) {

	fields, err := reflections.Fields(obj)
	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		if strings.ToLower(field) == strings.ToLower(fieldName) {
			return reflections.GetField(obj, field)
		}
	}

	return nil, errors.New("Object doesn't have a property named: " + fieldName)

}
