package xray

import (
	"fmt"
	"strings"
)

// KeyValue abstracts key with a value
type KeyValue struct {
	Key   string
	Value interface{}
}

// String returns the key value pair in the form <key>=<value>
func (kv *KeyValue) String() string {
	key := strings.ToLower(kv.Key)
	value := kv.Value
	if key == "" && value == "" {
		return ""
	}
	return fmt.Sprintf("%s=%s", key, value)
}
