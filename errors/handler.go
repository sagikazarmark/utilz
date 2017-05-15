package errors

// Handler is responsible for handling an error.
//
// This interface allows libraries to decouple from logging solutions.
// In most cases the implementation will provide some log functionalities though.
type Handler interface {
	Handle(err error)
}

// LogHandler accepts an errorLogger instance and logs an error.
//
// Compatible with most level-based loggers.
type LogHandler struct {
	logger errorLogger
}

// errorLogger covers most of the level-based logging solutions.
type errorLogger interface {
	Error(args ...interface{})
}

// NewLogHandler returns a new LogHandler.
func NewLogHandler(logger errorLogger) Handler {
	return &LogHandler{logger}
}

// Handle takes care of an error by logging it.
func (h *LogHandler) Handle(err error) {
	h.logger.Error(err)
}
