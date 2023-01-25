package strcase

// ToFlat converts a string to flatcase
// See https://www.deanpugh.com/a-brief-list-of-programming-naming-conventions/ Flat Case
func ToFlat(s string) string {
	return ToScreamingDelimited(s, 0, "", false)
}
