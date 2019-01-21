package flog_test

import (
	. "go.coder.com/flog"
	"testing"
)

func TestLogger(t *testing.T) {
	Log(INFO, "hello %.3f", 1/3.0)
	Log(SUCCESS, "finished that")
	Log(ERROR, "oops")

	log := New().WithPrefix("user %v: ", 500)

	log.Log(ERROR, "didn't work out")

	// Short-hand

	Info("something happened")
	Error("something bad happened")
	Success("something good happened")
}
