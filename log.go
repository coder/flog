// Package flog provides human-optimized, formatted logging.
package flog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// level wraps the time string with a level, by default
// a color.
type level func(msg string) string

func makeLevel(f func(s string, a ...interface{}) string) level {
	return func(msg string) string {
		return f("[%v]", msg)
	}
}

var (
	INFO    = makeLevel(color.HiBlueString)
	SUCCESS = makeLevel(color.HiGreenString)
	ERROR   = makeLevel(color.RedString)
	FATAL   = makeLevel(color.RedString)
)

// Time formats
const (
	ClockFormat     = "15:04:05"
	DateClockFormat = "2006-01-02" + ClockFormat
)

type Logger struct {
	W          io.Writer
	TimeFormat string
	Prefix     string
}

func (l *Logger) WithPrefix(p string, args ...interface{}) *Logger {
	ll := *l
	ll.Prefix += fmt.Sprintf(p, args...)
	return &ll
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log(INFO, msg, args...)
}

func (l *Logger) Success(msg string, args ...interface{}) {
	l.Log(SUCCESS, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log(ERROR, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Log(FATAL, msg, args...)
	os.Exit(1)
}

func (l *Logger) Log(lvl level, msg string, args ...interface{}) {
	fmt.Fprintf(
		l.W,
		fmt.Sprintf("%v ", lvl(time.Now().Format(l.TimeFormat)))+l.Prefix+msg+"\n",
		args...,
	)
}

// New returns a new logger.
func New() *Logger {
	return &Logger{
		W:          os.Stderr,
		TimeFormat: ClockFormat,
	}
}

func Info(msg string, args ...interface{}) {
	New().Log(INFO, msg, args...)
}

func Success(msg string, args ...interface{}) {
	New().Log(SUCCESS, msg, args...)
}

func Error(msg string, args ...interface{}) {
	New().Log(ERROR, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	New().Log(FATAL, msg, args...)
}

// Log logs a message with the default logger.
func Log(l level, m string, args ...interface{}) {
	New().Log(l, m, args...)
}
