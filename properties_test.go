package xray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-xray"
)

type sampleStruct struct {
	Name  string `form:"name" json:"name"`
	Title string `form:"title" json:"title"`
}

const sampleName = "Pieter"
const sampleTitle = "CTO"

func TestPropertiesValid(t *testing.T) {

	actual, err := xray.Properties(sampleStruct{
		Name:  sampleName,
		Title: sampleTitle,
	})

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, []string([]string{"Name", "Title"}), actual)

}

func TestPropertiesInvalidType(t *testing.T) {

	actual, err := xray.Properties(1)

	assert.Error(t, err, "error")
	assert.Nil(t, actual)

}

func TestPropertiesAsMapValid(t *testing.T) {

	actual, err := xray.PropertiesAsMap(sampleStruct{
		Name:  sampleName,
		Title: sampleTitle,
	})

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"Name": "Pieter", "Title": "CTO"}), actual)

}

func TestPropertiesAsMapInvalidType(t *testing.T) {

	actual, err := xray.PropertiesAsMap(1)

	assert.Error(t, err, "error")
	assert.Equal(t, map[string]interface{}{}, actual)

}

func TestPropertyValid(t *testing.T) {

	actual, err := xray.Property(sampleStruct{
		Name:  sampleName,
		Title: sampleTitle,
	}, "Name")

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, "Pieter", actual)

}

func TestPropertyNonExistingPropery(t *testing.T) {

	actual, err := xray.Property(sampleStruct{
		Name:  sampleName,
		Title: sampleTitle,
	}, "UnknownProperty")

	assert.Error(t, err, "error")
	assert.Nil(t, actual)

}

func TestPropertyInvalidType(t *testing.T) {

	actual, err := xray.Property(1, "a")

	assert.Error(t, err, "error")
	assert.Equal(t, nil, actual)

}
