package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelToDash(t *testing.T) {
	converted := CamelToDash("SomethingCamelCase")
	assert.Equal(t, "something-camel-case", converted, "Conversion mismatch")
}

func TestDashToCamel(t *testing.T) {
	converted := DashToCamel("something-dash-separated")
	assert.Equal(t, "SomethingDashSeparated", converted, "Conversion mismatch")
}

func TestCamelToSnake(t *testing.T) {
	converted := CamelToSnake("SomethingCamelCase")
	assert.Equal(t, "something_camel_case", converted, "Conversion mismatch")
}

func TestSnakeToCamel(t *testing.T) {
	converted := SnakeToCamel("something_snake_case")
	assert.Equal(t, "SomethingSnakeCase", converted, "Conversion mismatch")
}