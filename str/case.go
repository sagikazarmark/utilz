package str

import "unicode"

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

	for _, c := range str {
		switch {
		case unicode.IsUpper(c):
			// Prepend an underscore if the previos char is not an underscore
			// and the current char is not part of an abbreviation
			if prev != '_' && !unicode.IsUpper(prev) {
				buf += "_"
			}

			buf += string(unicode.ToLower(c))

		default:
			if c == '-' || c == ' ' {
				c = '_'
			}

			buf += string(c)
		}

		prev = c
	}

	return buf
}
