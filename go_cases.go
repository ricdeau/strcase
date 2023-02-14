package strcase

// GoExported as go exported structs/interfaces/fields e.g.
// Example: type SomeStruct int
func GoExported(s string) string {
	return ToCamel(s)
}

// GoUnexported as go unexported structs/interfaces/fields e.g.
// Example: type myInt int
func GoUnexported(s string) string {
	return ToLowerCamel(s)
}

// GoFileName as go file name.
// Example: go_file_name
func GoFileName(s string) string {
	return ToSnake(s)
}
