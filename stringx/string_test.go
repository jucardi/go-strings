package stringx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Replace(t *testing.T) {
	str := "aaaabcde01234"
	strx := New("aaaabcde01234")

	assert.Equal(t, strings.Replace(str, "aaaa", "ZX", -1), strx.Replace("aaaa", "ZX", -1).S())

	multiReplace1 := strings.Replace(strings.Replace(strings.Replace(str, "aaaa", "ZX", -1), "012", "www", -1), "bcd", "ppp", -1)
	multiReplace2 := strx.
		Replace("aaaa", "ZX", -1).
		Replace("012", "www", -1).
		Replace("bcd", "ppp", -1).
		S()

	assert.Equal(t, multiReplace1, multiReplace2)
}

func TestString_StringTransforms(t *testing.T) {
	str := "aaaabcde01234"
	strx := New("aaaabcde01234")

	multiReplace1 := strings.Replace(strings.Replace(strings.Replace(str, "aaaa", "ZX", -1), "012", "www", -1), "bcd", "ppp", -1)
	multiReplace1 = strings.ToTitle(strings.ToLower(strings.ToUpper(multiReplace1)))

	multiReplace2 := strx.
		Replace("aaaa", "ZX", -1).
		Replace("012", "www", -1).
		Replace("bcd", "ppp", -1).
		ToUpper().
		ToLower().
		ToTitle().
		S()

	assert.Equal(t, multiReplace1, multiReplace2)
}
