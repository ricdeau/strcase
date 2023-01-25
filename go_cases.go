package strcase

// GoExported as go exported structs/interfaces/fields e.g.
// Example: type SomeStruct int
func GoExported(s string) string {
	return ToCamel(s)
}

// GoUnexported as go unexported structs/interfaces/fields e.g.
// Example: var name = ""
func GoUnexported(s string) string {
	return ToLowerCamel(s)
}
