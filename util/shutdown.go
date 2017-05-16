package util

import "github.com/sagikazarmark/utilz/errors"

// ShutdownManager manages an application shutdown by calling the registered handlers.
type ShutdownManager struct {
	handlers     []ShutdownHandler
	errorHandler errors.Handler
}

// ShutdownHandler is any function that has no parameters and can return an error.
//
// Shutdown handlers are the last resort for the application when there is no more
// flow control in the user's hand.
//
// Returned errors are logged.
type ShutdownHandler func() error

// ShutdownFunc wraps a function withot an error return type.
//
// To make sure there are no silenced errors, panics are also recovered.
func ShutdownFunc(fn func()) ShutdownHandler {
	return func() (err error) {
		defer func() {
			err = errors.Recover(recover())
		}()

		fn()

		return err
	}
}

// NewShutdownManager creates a new Shutdown manager.
func NewShutdownManager(errorHandler errors.Handler) *ShutdownManager {
	if errorHandler == nil {
		errorHandler = errors.NewNullHandler()
	}

	return &ShutdownManager{
		errorHandler: errorHandler,
	}
}

// Register appends new shutdown handlers to the list of existing ones.
func (sm *ShutdownManager) Register(handlers ...ShutdownHandler) {
	sm.handlers = append(sm.handlers, handlers...)
}

// RegisterAsFirst prepends new shutdown handlers to the list of existing ones.
func (sm *ShutdownManager) RegisterAsFirst(handlers ...ShutdownHandler) {
	sm.handlers = append(handlers, sm.handlers...)
}

// Shutdown is the panic recovery and shutdown handler.
//
// It should be called as the last method in `main` (eg. using defer).
func (sm *ShutdownManager) Shutdown() {
	// Try recovering from panic first
	err := errors.Recover(recover())
	if err != nil {
		sm.errorHandler.Handle(err)
	}

	// Loop through all the handlers and call them
	// Log any errors that may occur
	for _, handler := range sm.handlers {
		err := handler()
		if err != nil {
			sm.errorHandler.Handle(err)
		}
	}
}
