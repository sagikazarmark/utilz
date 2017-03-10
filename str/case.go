package str

import (
	"unicode"
	"unicode/utf8"
)

// ToSnake converts a string (camel or spinal) to snake case
func ToSnake(str string) string {
	// Skip processing for an empty string
	if len(str) == 0 {
		return ""
	}

	// Build the results in this buffer
	buf := ""

	// Trick: if the first character is uppercase, do not prepend an underscore
	prev := '_'

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		switch {
		case unicode.IsUpper(r):
			// Prepend an underscore if the previous char is not an underscore
			// and the current char is not part of an abbreviation
			if prev != '_' && !unicode.IsUpper(prev) {
				buf += "_"
			}

			buf += string(unicode.ToLower(r))

		default:
			if r == '-' || r == ' ' {
				r = '_'
			}

			buf += string(r)
		}

		prev = r
	}

	return buf
}