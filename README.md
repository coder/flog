# flog

`flog` is a minimal, formatted, pretty logging package for Go.

It's optimized for human readability.

[Uber's Zap](https://github.com/uber-go/zap) is recommended for more robust logging.

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/go.coder.com/flog)

## Install

`go get go.coder.com/flog`

## Usage

```go
// Use a dot import for optimal readability.
Log(INFO, "hello %.3f", 1/3.0)
Log(SUCCESS, "finished that")
Log(ERROR, "oops")

log := NewLogger().WithPrefix("user %v: ", 500)

log.Log(ERROR, "didn't work out")
```

produces

![example](docs/usage.png)


