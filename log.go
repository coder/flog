// Package flog provides human-optimized, formatted logging.
package flog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// Level enum type, representing the severity level of a log entry.
type Level string

// Level enum values.
var (
	InfoLevel    = Level(color.HiBlueString("INFO"))
	SuccessLevel = Level(color.HiGreenString("SUCCESS"))
	ErrorLevel   = Level(color.RedString("ERROR"))
	FatalLevel   = Level(color.RedString("FATAL"))
)

// Deprecated: Upper case values exists for historical compatibility
// and should not be used. Prefer the linted, capitalized xxxLevel version.
var (
	INFO    = InfoLevel
	SUCCESS = SuccessLevel
	ERROR   = ErrorLevel
	FATAL   = FatalLevel
)

// Time format used as prefix in all log messages.
const timePrefixFormat = `2006-01-02 15:04:05`

// Logger is the main logger struct, wrapping a writer and settings.
type Logger struct {
	W      io.Writer
	Prefix string
}

// WithPrefixf makes creates a copy of the logger and sets a appends the given prefix to it.
func (l *Logger) WithPrefixf(p string, args ...interface{}) *Logger {
	ll := *l
	ll.Prefix += fmt.Sprintf(p, args...)

	return &ll
}

// Infof logs an entry with the Infof level.
func (l *Logger) Infof(msg string, args ...interface{}) { l.Logf(InfoLevel, msg, args...) }

// Successf logs an entry with the Successf level.
func (l *Logger) Successf(msg string, args ...interface{}) { l.Logf(SuccessLevel, msg, args...) }

// Errorf logs an entry with the Errorf level.
func (l *Logger) Errorf(msg string, args ...interface{}) { l.Logf(ErrorLevel, msg, args...) }

// Fatalf logs an entry with the Fatalf level. Exits the process.
func (l *Logger) Fatalf(msg string, args ...interface{}) { l.Logf(FatalLevel, msg, args...) }

// Logf the given message with a severity level. If Fatal, exits the process.
func (l *Logger) Logf(lvl Level, msg string, args ...interface{}) {
	fmt.Fprintf(l.W, //  Log to the underlying writer.
		"%s %s\t%s%s\n", // [current time] [severity level]\t[preset perfix][formatted msg].
		time.Now().Format(timePrefixFormat),
		lvl,
		l.Prefix,
		fmt.Sprintf(msg, args...),
	)

	if lvl == FatalLevel {
		os.Exit(1)
	}
}

// WithPrefix maps to WithPrefixf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) WithPrefix(p string, args ...interface{}) *Logger { return l.WithPrefixf(p, args...) }

// Info maps to Infof.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) Info(msg string, args ...interface{}) { l.Infof(msg, args...) }

// Success maps to Successf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) Success(msg string, args ...interface{}) { l.Successf(msg, args...) }

// Error maps to Errorf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) Error(msg string, args ...interface{}) { l.Errorf(msg, args...) }

// Fatal maps to Fatalf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) Fatal(msg string, args ...interface{}) { l.Fatalf(msg, args...) }

// Log maps to Logf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func (l *Logger) Log(lvl Level, msg string, args ...interface{}) { l.Logf(lvl, msg, args...) }

// New returns a new logger.
func New() *Logger { return &Logger{W: os.Stderr} }

// Infof creates a new default logger and logs an entry with the Infof level.
func Infof(msg string, args ...interface{}) { New().Logf(InfoLevel, msg, args...) }

// Successf creates a new default logger and logs an entry with the Successf level.
func Successf(msg string, args ...interface{}) { New().Logf(SuccessLevel, msg, args...) }

// Errorf creates a new default logger and logs an entry with the Errorf level.
func Errorf(msg string, args ...interface{}) { New().Logf(ErrorLevel, msg, args...) }

// Fatalf creates a new default logger and logs an entry with the Fatalf level. Exits the process.
func Fatalf(msg string, args ...interface{}) { New().Logf(FatalLevel, msg, args...) }

// Logf creates a new default logger and logs an entry with the given severity level.
func Logf(l Level, m string, args ...interface{}) { New().Logf(l, m, args...) }

// Info maps to Infof.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func Info(msg string, args ...interface{}) { Infof(msg, args...) }

// Success maps to Successf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func Success(msg string, args ...interface{}) { Successf(msg, args...) }

// Error maps to Errorf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func Error(msg string, args ...interface{}) { Errorf(msg, args...) }

// Fatal maps to Fatalf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func Fatal(msg string, args ...interface{}) { Fatalf(msg, args...) }

// Log maps to Logf.
// Deprecated: Non-standard and prevents tools like go vet to check for errors.
// The xxxf version should be used.
func Log(l Level, m string, args ...interface{}) { Logf(l, m, args...) }
