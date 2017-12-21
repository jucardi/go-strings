package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringBuilder_Append(t *testing.T) {
	result := Builder().Append("some").Append("thing").Build()
	assert.Equal(t, "something", result, "Builder result mismatch")
}

func TestStringBuilder_AppendLine(t *testing.T) {
	result := Builder().AppendLine("some").AppendLine("thing").Build()
	assert.Equal(t, "some\nthing\n", result, "Builder result mismatch")
}

func TestStringBuilder_Br(t *testing.T) {
	result := Builder().Br().Br().Build()
	assert.Equal(t, "\n\n", result, "Builder result mismatch")
}

func TestStringBuilder_AppendInt(t *testing.T) {
	result := Builder().AppendInt(12).AppendInt(34).Build()
	assert.Equal(t, "1234", result, "Builder result mismatch")
}

func TestStringBuilder_IsEmpty(t *testing.T) {
	builder := Builder()
	assert.True(t, builder.IsEmpty())
	builder.AppendLine("something")
	assert.False(t, builder.IsEmpty())
}