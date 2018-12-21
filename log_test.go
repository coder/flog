package flog_test

import (
	"go.coder.com/flog"
	"testing"
)

func TestLogger(t *testing.T) {
	flog.Log(flog.INFO, "hello %.3f", 1/3.0)
	flog.Log(flog.SUCCESS, "finished that")
	flog.Log(flog.ERROR, "oops")

	log := flog.New().WithPrefix("user %v: ", 500)

	log.Log(flog.ERROR, "didn't work out")
}
