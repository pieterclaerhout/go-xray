package xray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-xray"
)

func TestName(t *testing.T) {

	type test struct {
		name     string
		obj      interface{}
		expected string
	}

	var tests = []test{
		{"nil", nil, "<nil>"},
		{"value", test{}, "test"},
		{"pointer", &test{}, "test"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := xray.Name(tc.obj)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}

}
