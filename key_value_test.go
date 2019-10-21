package xray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-xray"
)

func TestKeyValueString(t *testing.T) {

	type test struct {
		name     string
		key      string
		value    string
		expected string
	}

	var tests = []test{
		{"key+value", "key", "value", "key=value"},
		{"key-only", "key", "", "key="},
		{"empty", "", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			obj := &xray.KeyValue{
				Key:   tc.key,
				Value: tc.value,
			}

			actual := obj.String()

			assert.Equal(t, tc.expected, actual, tc.name)

		})
	}

}
