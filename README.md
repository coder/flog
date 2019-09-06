# flog

`flog` is a minimal, formatted, pretty logging package for Go.

It's optimized for human readability.

[slog](https://github.com/cdr/slog) is recommended for more robust logging.

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/go.coder.com/flog)

## Install

`go get go.coder.com/flog`

## Usage

```go
flog.Info("hello %.3f", 1/3.0)
flog.Success("finished that")
flog.Error("oops")

log := flog.NewLogger().WithPrefix("user %v: ", 500)

log.Error("didn't work out")
```

produces

![example](docs/usage.png)


