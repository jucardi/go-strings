package stringx

import (
	"bytes"
	"strconv"
)

const LINE_BREAK = "\n"

type StringBuilder struct {
	buffer bytes.Buffer
}

// Creates a new StringBuilder
func Builder() *StringBuilder {
	return &StringBuilder{}
}

// Appends the text
func (s *StringBuilder) Append(str string) *StringBuilder {
	s.buffer.WriteString(str)
	return s
}

// Appends the number
func (s *StringBuilder) AppendInt(i int) *StringBuilder {
	return s.Append(strconv.Itoa(i))
}

// Appends the text and breaks to the next line
func (s *StringBuilder) AppendLine(str string) *StringBuilder {
	return s.Append(str).Br()
}

// Appends a single character to the builder
func (s *StringBuilder) AppendRune(char rune) *StringBuilder {
	s.buffer.WriteRune(char)
	return s
}

// Breaks to the next line
func (s *StringBuilder) Br() *StringBuilder {
	return s.Append(LINE_BREAK)
}

// Indicates whether the builder is empty
func (s *StringBuilder) IsEmpty() bool {
	return s.buffer.Len() == 0
}

// Builds the string
func (s *StringBuilder) Build() string {
	return s.buffer.String()
}
