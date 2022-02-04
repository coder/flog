package flog_test

import (
	"testing"

	. "github.com/coder/flog"
)

func TestLogger(t *testing.T) {
	log(infoLevel, "hello %.3f", 1/3.0)
	log(successLevel, "finished that")
	log(error, "oops")

	log := New().WithPrefix("user %v: ", 500)

	log.log(error, "didn't work out")

	// Short-hand

	Info("something happened")
	Error("something bad happened")
	Success("something good happened")
}
