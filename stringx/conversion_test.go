package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelToDash(t *testing.T) {
	converted := CamelToDash("somethingCamelCase")
	assert.Equal(t, "something-camel-case", converted, "Conversion mismatch")
}

func TestCamelToSnake(t *testing.T) {
	converted := CamelToSnake("somethingCamelCase")
	assert.Equal(t, "something_camel_case", converted, "Conversion mismatch")
}

func TestDashToCamel(t *testing.T) {
	converted := DashToCamel("something-dash-separated")
	assert.Equal(t, "somethingDashSeparated", converted, "Conversion mismatch")
}

func TestSnakeToCamel(t *testing.T) {
	converted := SnakeToCamel("SOMETHING_SNAKE_CASE")
	assert.Equal(t, "somethingSnakeCase", converted, "Conversion mismatch")
}

func TestPascalToDash(t *testing.T) {
	converted := CamelToDash("SomethingPascalCase")
	assert.Equal(t, "something-pascal-case", converted, "Conversion mismatch")
}

func TestPascalToSnake(t *testing.T) {
	converted := CamelToSnake("SomethingPascalCase")
	assert.Equal(t, "something_pascal_case", converted, "Conversion mismatch")
}

func TestDashToPascal(t *testing.T) {
	converted := DashToPascal("something-dash-separated")
	assert.Equal(t, "SomethingDashSeparated", converted, "Conversion mismatch")
}

func TestSnakeToPascal(t *testing.T) {
	converted := SnakeToPascal("something_snake_case")
	assert.Equal(t, "SomethingSnakeCase", converted, "Conversion mismatch")
}