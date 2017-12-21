# go-strings
String utilities for go

#### String Builder

The string builder allows to easily build strings using a byte.buffer underneath.

Supports
- **Append**:      Appends the given text to the builder
- **AppendLine**:  Appends the given text to the builder and adds a line break at the end.
- **AppendInt**:   Appends an int to the builder converting it to a string.
- **Br**:          Appends a line break to the builder.
- **IsEmpty**:     Indicates whether the builder is empty.

Usage:
```Go
builder := stringx.Builder()

builder.Append("some text ")
builder.AppendLine("something else")
builder.AppendInt(1234)
builder.Br()
builder.Append("End")

result := builder.Build()
```

The result will be equal to

```
some text something else
1234
End
```

#### Conversion methods

- **CamelToSnake**: Converts CamelCase to snake_case
- **SnakeToCamel**: Converts snake_case to CamelCase
- **CamelToDash**:  Converts CamelCase to dash-separated-string
- **DashToCamel**:  Converts a dash-separated-string to CamelCase
