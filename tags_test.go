package xray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-xray"
)

func Test_Tags_Valid(t *testing.T) {

	type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

	actual, err := xray.Tags(sampleStruct{}, "form")

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, map[string]string{"Name": "name", "Title": "title"}, actual)

}

func Test_Tags_UnknownTag(t *testing.T) {

	type sampleStruct struct {
		Name string `form:"name" json:"name"`
	}

	actual, err := xray.Tags(sampleStruct{}, "db")

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, map[string]string{"Name": ""}, actual)

}

func Test_Tags_InvalidTag(t *testing.T) {

	type sampleStruct struct {
		Name string `form:"name`
	}

	actual, err := xray.Tags(sampleStruct{}, "form")

	assert.NoError(t, err, "error")
	assert.NotNil(t, actual)
	assert.Equal(t, map[string]string{"Name": ""}, actual)

}

func Test_Tags_InvalidType(t *testing.T) {

	actual, err := xray.Tags(1, "form")

	assert.Error(t, err, "error")
	assert.Nil(t, actual)

}
