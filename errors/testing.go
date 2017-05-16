package errors

// TestHandler is a test implementation of errors.Handler
type TestHandler struct {
	errors []error
}

// NewTestHandler returns a new LogHandler.
func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Handle takes care of an error by logging it.
func (h *TestHandler) Handle(err error) {
	h.errors = append(h.errors, err)
}

// Errors returns the list of handled errors.
func (h *TestHandler) Errors() []error {
	return h.errors
}

// Last returns the last error (if any).
func (h *TestHandler) Last() error {
	count := len(h.errors)

	// Return the last error (if any)
	if count > 0 {
		return h.errors[count-1]
	}

	return nil
}
