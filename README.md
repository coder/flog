# flog

`flog` is a minimal, formatted, pretty logging package for Go.

It's optimized for human readability and CLIs.
* Log levels color-code the timestamp
* Log level doesn't change the width of the prefix

[slog](https://github.com/cdr/slog) is recommended for robust logging.

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/coder/flog)

## Install

`go get github.com/coder/flog`

## Usage

```go
flog.Infof("hello %.3f", 1/3.0)
flog.Successf("finished that")
flog.Errorf("oops")

log := flog.New(os.Stderr).WithPrefix("user %v: ", 500)

log.Errorf("didn't work out")
```

## Output

![example](docs/usage.png)
