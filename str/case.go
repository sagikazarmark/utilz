package str

import (
	"bytes"
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
	buf := new(bytes.Buffer)

	// Trick: if the first rune is uppercase, do not prepend an underscore
	prev := '_'

	// Used to count the number of uppercase runes
	// If more than two is found, we suppress writting until the next rune
	// where we decide if an underscore needs to be prepended or not
	upperCount := 0

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		switch {
		case unicode.IsUpper(r):
			upperCount++

			// Switch to post writting mode
			if upperCount == 2 {
				break
			}

			// If we are in post writting mode, write the previous rune and store the current
			// Otherwise continue as usual
			if upperCount > 2 {
				buf.WriteRune(unicode.ToLower(prev))
			} else {
				// Prepend an underscore if the previous char is not an underscore
				// and the current char is not part of an abbreviation
				if prev != '_' && !unicode.IsUpper(prev) {
					buf.WriteRune('_')
				}

				buf.WriteRune(unicode.ToLower(r))
			}

		default:
			if upperCount > 2 {
				buf.WriteRune('_')
				buf.WriteRune(unicode.ToLower(prev))
			}

			upperCount = 0

			if r == '-' || r == ' ' {
				r = '_'
			}

			buf.WriteRune(r)
		}

		prev = r
	}

	return buf.String()
}
