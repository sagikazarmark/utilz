package strings

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// ToSnake converts a string (camel or spinal) to snake_case.
func ToSnake(s string) string {
	return toSeparated(s, '_', false)
}

// ToSpinal converts a string (camel or snake) to spinal-case.
//
// See https://en.wikipedia.org/wiki/Letter_case#Special_case_styles
func ToSpinal(s string) string {
	return toSeparated(s, '-', false)
}

// ToTrain converts a string (camel, snake or spinal) to Train-Case.
//
// See https://en.wikipedia.org/wiki/Letter_case#Special_case_styles
func ToTrain(s string) string {
	return toSeparated(s, '-', true)
}

// toSeparated converts a string (camel, snake or spinal) to a lower-cased separated one (essentially snake or spinal based on the separator).
func toSeparated(s string, sep rune, t bool) string {
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

			// If train case is enabled and the previous rune us a separator, make it upper case
			if t && (writeSep || prev == sep) {
				buf.WriteRune(r)
			} else {
				buf.WriteRune(unicode.ToLower(r))
			}

		default:
			if r == '_' || r == '-' || r == ' ' {
				r = sep
			}

			// If train case is enabled and the previous rune us a separator, make it upper case
			if t && prev == sep {
				r = unicode.ToUpper(r)
			}

			buf.WriteRune(r)
		}

		prev = r
	}

	return buf.String()
}
