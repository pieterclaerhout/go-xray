package xray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-xray"
)

func Test_Properties_Valid(t *testing.T) {

	type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

	actual, err := xray.Properties(sampleStruct{
		Name:  "Pieter",
		Title: "CTO",
	})

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, []string([]string{"Name", "Title"}), actual)

}

func Test_Properties_InvalidType(t *testing.T) {

	actual, err := xray.Properties(1)

	assert.Error(t, err, "error")
	assert.Nil(t, actual)

}

func Test_PropertiesAsMap_Valid(t *testing.T) {

	type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

	actual, err := xray.PropertiesAsMap(sampleStruct{
		Name:  "Pieter",
		Title: "CTO",
	})

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"Name": "Pieter", "Title": "CTO"}), actual)

}

func Test_PropertiesAsMap_InvalidType(t *testing.T) {

	actual, err := xray.PropertiesAsMap(1)

	assert.Error(t, err, "error")
	assert.Equal(t, map[string]interface{}{}, actual)

}

func Test_Property_Valid(t *testing.T) {

	type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

	actual, err := xray.Property(sampleStruct{
		Name:  "Pieter",
		Title: "CTO",
	}, "Name")

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, "Pieter", actual)

}

func Test_Property_NonExistingPropery(t *testing.T) {

	type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

	actual, err := xray.Property(sampleStruct{
		Name:  "Pieter",
		Title: "CTO",
	}, "UnknownProperty")

	assert.Error(t, err, "error")
	assert.Nil(t, actual)

}

func Test_Property_InvalidType(t *testing.T) {

	actual, err := xray.Property(1, "a")

	assert.Error(t, err, "error")
	assert.Equal(t, nil, actual)

}
