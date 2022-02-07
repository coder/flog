package flog_test

import (
	"testing"

	. "github.com/coder/flog"
)

func TestLogger(t *testing.T) {
	// Short-hand

	Infof("something happened")
	Errorf("something bad happened")
	Successf("something good happened")
}
