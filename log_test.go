package flog_test

import (
	"testing"

	. "go.coder.com/flog"
)

func TestLogger(_ *testing.T) {
	Logf(InfoLevel, "hello %.3f", 1/3.0)
	Logf(SuccessLevel, "finished that")
	Logf(ErrorLevel, "oops")

	log := New().WithPrefixf("user %d: ", 500)

	log.Logf(ErrorLevel, "didn't work out")

	// Short-hand

	Infof("something happened")
	Errorf("something bad happened")
	Successf("something good happened")
}
