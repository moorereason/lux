# lux

Package lux provides the ability to color and enhance terminal output with ANSI escape codes.
The package is implemented as a `fmt.Formatter` interface.

## Install

```text
go get github.com/moorereason/lux
```

## Usage

```go
fmt.Printf("%s\n", lux.Red("RED!"))
fmt.Printf("%s\n", lux.BgWhite(lux.Red("RED!")))
```
