package flog_test

import (
	"testing"

	. "github.com/coder/flog"
)

func TestLogger(t *testing.T) {
	// Short-hand

	Info("something happened")
	Error("something bad happened")
	Success("something good happened")
}
