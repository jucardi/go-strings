# go-strings
String utilities for go

#### Getting started

To keep up to date with the most recent version:

```bash
go get github.com/jucardi/go-strings
```

To get a specific version:

```bash
go get gopkg.in/jucardi/go-strings.v1
```

### String Builder

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

### String struct wrapper

What inspired me to create this wrapper was the fact that strings do not have functions built in like in other languages, so doing functional operations over the same string is
not possible with the default `string` object.

Have you ever ran into a situation where you needed to do multiple replaces to the same string? and maybe plus a `ToLower` or `ToUpper` to the result?

**The Go way**
```go
    // Given a string value `str`
    result := strings.ToLower(strings.Replace(strings.Replace(strings.Replace(str, "aaaa", "ZX", -1), "012", "www", -1), "bcd", "ppp", -1))
```

**The String way in `stringx`**
```go
    // Given a string value `str`

    result := stringx.New(str).	Replace("aaaa", "ZX", -1).Replace("012", "www", -1).Replace("bcd", "ppp", -1).ToLower().S()

```

**If you like splitting functional syntax in multiple lines**
```go
    result := stringx.New(str).
        Replace("aaaa", "ZX", -1).
        Replace("012", "www", -1).
        Replace("bcd", "ppp", -1).
        ToLower().
        S()
```

### Conversion methods

- **CamelToSnake**: Converts CamelCase to snake_case
- **SnakeToCamel**: Converts snake_case to CamelCase
- **CamelToDash**:  Converts CamelCase to dash-separated-string
- **DashToCamel**:  Converts a dash-separated-string to CamelCase
