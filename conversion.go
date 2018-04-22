package stringx

import (
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// Converts CamelCase to snake_case.
func CamelToSnake(s string) string {
	return camelToSymbolSeparated(s, "_")
}

// Converts snake_case to CamelCase
func SnakeToCamel(s string) string {
	return symbolSeparatedToCamel(s, "_")
}

// Converts CamelCase to dash-separated-string
func CamelToDash(s string) string {
	return camelToSymbolSeparated(s, "-")
}

// Converts a dash-separated-string to CamelCase
func DashToCamel(s string) string {
	return symbolSeparatedToCamel(s, "-")
}

func symbolSeparatedToCamel(s string, separator string) string {
	var ret []string

	for _, v := range strings.Split(s, separator) {
		ret = append(ret, strings.Title(v))
	}

	return strings.Join(ret, "")
}

func camelToSymbolSeparated(s string, separator string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, separator))
}
