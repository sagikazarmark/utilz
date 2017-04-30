// Package utilz contains all kinds of utilities and extensions of the standard library
package utilz

// Must checks if a parameter is an error and panics if so.
// Useful when you want to force a call to suceed.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
