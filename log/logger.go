package log

// Context holds the structured data for a log as a key-value map.
type Context map[string]interface{}

// StructuredLogger allows to attach context to a logger.
type StructuredLogger interface {
	WithField(key string, value interface{}) StructuredLogger
	WithContext(ctx Context) StructuredLogger
}

// LevelLogger attaches a level to the log.
type LevelLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}

// StdLogger declares an interface for the Logger in the standard library.
type StdLogger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// StructuredLevelLogger combines structured and level-based logging.
type StructuredLevelLogger interface {
	StructuredLogger
	LevelLogger
}

// Logger combines all goods of logging.
type Logger interface {
	StructuredLevelLogger
	StdLogger
}
