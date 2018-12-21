// Package flog provides human-optimized, formatted logging.
// It's symbols' names may stutter, as it's highly recommended to dot import the package.
package flog

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

type Level string

var (
	INFO    = Level(color.HiBlueString("INFO"))
	SUCCESS = Level(color.HiGreenString("SUCCESS"))
	ERROR   = Level(color.RedString("ERROR"))
	FATAL   = Level(color.RedString("FATAL"))
)

type Logger struct {
	W      io.Writer
	Prefix string
}

func (l *Logger) WithPrefix(p string, args ...interface{}) *Logger {
	ll := *l
	ll.Prefix += fmt.Sprintf(p, args...)
	return &ll
}

func (l *Logger) Log(lvl Level, msg string, args ...interface{}) {
	fmt.Fprintf(
		l.W,
		fmt.Sprintf("%v %v\t", time.Now().Format(`2006-01-02 15:04:05`), lvl)+
			l.Prefix + msg+"\n", args...,
	)
	if lvl == FATAL {
		os.Exit(1)
	}
}

// NewLogger returns a new logger.
func NewLogger() *Logger {
	return &Logger{
		W: os.Stderr,
	}
}

// Log logs a message with the default logger.
func Log(l Level, m string, args ...interface{}) {
	NewLogger().Log(l, m, args...)
}
