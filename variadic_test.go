package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	defaultKey string
	innerMap   map[string]string
}

func (t *testStruct) Get(name ...string) string {
	return t.innerMap[GetOrDefault(t.defaultKey, name...)]
}

func TestGetOrDefault(t *testing.T) {
	v := &testStruct{
		defaultKey: "default",
		innerMap: map[string]string{
			"default": "a",
			"key1":    "b",
			"key2":    "c",
		},
	}

	assert.Equal(t, "a", v.Get("default"))
	assert.Equal(t, "b", v.Get("key1"))
	assert.Equal(t, "a", v.Get())
}
