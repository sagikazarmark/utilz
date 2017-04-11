package strings

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// ToSnake converts a string (camel or spinal) to snake case.
func ToSnake(s string) string {
	return toSeparated(s, '_')
}

// ToSpinal converts a string (camel or snake) to spinal case.
func ToSpinal(s string) string {
	return toSeparated(s, '-')
}

// toSeparated converts a string (camel, snake or spinal) to a lower-cased separated one (essentially snake or spinal based on the separator).
func toSeparated(s string, sep rune) string {
	// Skip processing for an empty string
	if len(s) == 0 {
		return ""
	}

	// Build the results in this buffer
	buf := new(bytes.Buffer)

	// Trick: if the first rune is uppercase, do not prepend a separator
	prev := sep

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		switch {
		case unicode.IsUpper(r):
			var writeSep bool

			if len(s) > 0 {
				next, _ := utf8.DecodeRuneInString(s)
				writeSep = prev != sep && (!unicode.IsUpper(prev) || (unicode.IsUpper(prev) && !unicode.IsUpper(next)))
			} else {
				writeSep = prev != sep && !unicode.IsUpper(prev)
			}

			// Prepend a separator if the previous char is not a separator
			// and the current char is not part of an abbreviation
			if writeSep {
				buf.WriteRune(sep)
			}

			buf.WriteRune(unicode.ToLower(r))

		default:
			if r == '_' || r == '-' || r == ' ' {
				r = sep
			}

			buf.WriteRune(r)
		}

		prev = r
	}

	return buf.String()
}
