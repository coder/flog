# flog

[![GitHub Release](https://img.shields.io/github/v/release/cdr/flog?color=6b9ded&sort=semver)](https://github.com/cdr/flog/releases)
[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/go.coder.com/flog?tab=doc)
[![license](https://img.shields.io/github/license/cdr/flog)](https://raw.githubusercontent.com/cdr/flog/master/LICENSE)

`flog` is a minimal, formatted, pretty logging package for Go.

It's optimized for human readability.

[slog](https://github.com/cdr/slog) is recommended for more robust logging.

## Install

`go get go.coder.com/flog`

## Usage

```go
flog.Infof("hello %.3f", 1/3.0)
flog.Successf("finished that")
flog.Errorf("oops")

log := flog.NewLogger().WithPrefixf("user %d: ", 500)

log.Errorf("didn't work out")
```

produces

![example](docs/usage.png)


## Linter settings

By default, the linters will not understand that `flog` is a logging library expected formatting. To fix this, we need to update the settings.

### golangci-lint

When using [golangci-lint](https://github.com/golangci/golangci-lint), the `govet` settings can be set in the `.golangci.yml` file:

```yaml
linters-settings:
  govet:
    settings:
      printf: # Analyzer name.
        funcs: # Run `go tool vet help printf` to see available settings for `printf` analyzer.
          - (go.coder.com/flog).Errorf
          - (go.coder.com/flog).Fatalf
          - (go.coder.com/flog).Infof
          - (go.coder.com/flog).Logf
          - (go.coder.com/flog).Successf
          - (*go.coder.com/flog.Logger).Errorf
          - (*go.coder.com/flog.Logger).Fatalf
          - (*go.coder.com/flog.Logger).Infof
          - (*go.coder.com/flog.Logger).Logf
          - (*go.coder.com/flog.Logger).Successf
          - (*go.coder.com/flog.Logger).WithPrefixf
```

### go vet

When using `go vet` directly, `-printf.funcs` can be used.

```bash
go vet -printf.funcs Errorf,Fatalf,Infof,Logf,Successf,WithPrefixf ./...
```
