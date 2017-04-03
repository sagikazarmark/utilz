package strings

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// ToSnake converts a string (camel or spinal) to snake case.
func ToSnake(s string) string {
	// Skip processing for an empty string
	if len(s) == 0 {
		return ""
	}

	// Build the results in this buffer
	buf := new(bytes.Buffer)

	// Trick: if the first rune is uppercase, do not prepend an underscore
	prev := '_'

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		switch {
		case unicode.IsUpper(r):
			var uc bool

			if len(s) > 0 {
				next, _ := utf8.DecodeRuneInString(s)
				uc = prev != '_' && (!unicode.IsUpper(prev) || (unicode.IsUpper(prev) && !unicode.IsUpper(next)))
			} else {
				uc = prev != '_' && !unicode.IsUpper(prev)
			}

			// Prepend an underscore if the previous char is not an underscore
			// and the current char is not part of an abbreviation
			if uc {
				buf.WriteRune('_')
			}

			buf.WriteRune(unicode.ToLower(r))

		default:
			if r == '-' || r == ' ' {
				r = '_'
			}

			buf.WriteRune(r)
		}

		prev = r
	}

	return buf.String()
}
