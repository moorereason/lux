# lux [![GoDoc](https://godoc.org/github.com/moorereason/lux?status.svg)](http://godoc.org/github.com/moorereason/lux) [![Travis-CI](https://travis-ci.org/moorereason/lux.svg)](https://travis-ci.org/moorereason/lux) [![Report card](https://goreportcard.com/badge/github.com/moorereason/lux)](https://goreportcard.com/report/github.com/moorereason/lux)

Package `lux` provides a `fmt.Formatter` interface for outputting ANSI escape codes on the terminal.

## Install

```text
go get github.com/moorereason/lux
```

## Usage

```go
import (
  "fmt"

  "github.com/moorereason/lux"
)

func main() {
	// No Windows support
	fmt.Printf("%s\n", lux.Red("Red!"))

	// Windows supported - see explanation in README
	lux.Printf("%s\n", lux.Red("Red!"))
}
```

## Snarky, Kitchen Sink Example

```go
package main

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/moorereason/lux"
)

func main() {

Start:
	// Don't care about Windows?
	fmt.Printf("%s\n", lux.Red("Red!"))
	fmt.Printf("%s\n", lux.Red(lux.BgWhite("Red on White")))

	// Oh, you do care about Windows. Good call...
	Stdout := colorable.NewColorableStdout()
	fmt.Fprintf(Stdout, "%s\n", lux.Red("Red!")) // *sigh*

	// Want lux to handle Stdout? Use built-in Print functions (Print, Println, Printf)
	lux.Printf("%s\n", lux.Red("Red!"))

	// Hang on to your styles
	p := lux.New(lux.Red, lux.BgWhite)
	p.Println("Big Red")

	// First-Class Prints!
	bigred := lux.New(lux.Bold, lux.Red).Println
	bigred("Notice!")

	// Make Terminals Great Again
	r := lux.HiRed
	w := lux.HiWhite
	b := lux.HiBlue
	lux.Printf("%s%s%s! %s%s%s!\n", r("U"), w("S"), b("A"), r("U"), w("S"), b("A"))

	if lux.NoColor {
		goto Exit
	} else {
		// Let's do it again without colors.
		// I know. Crazy, right?
		lux.NoColor = true
		goto Start
	}
Exit:
}
```

## Windows Support

TL;DR: If you want cross-platform support, use lux.Print/Println/Printf instead
of the `fmt` package.

Two things make Windows support difficult:

1. Windows doesn't use escape codes but uses console drawing API calls
2. Go's `io.Stdout` is a `*io.File` type, so you can't simply set it to an
   `io.Writer`. See [golang/go#13473][stdout]

Because of these limitations, creating cross-platform console colors requires
writing wrappers around `fmt.Fprintf` to write to an `io.Writer` provided by
the [colorable][colorable] package.

## Acknowledgments

- [@mattn][mattn] for the [colorable][colorable] package.
- [@fatih][fatih] for the [color] package,
  which is a more widely used package and accomplishes many of the same goals as `lux`.
  The main difference is that `lux` is a `fmt.Formatter` implementation.
    - Bonus shoutout to [@fatih][fatih] for the awesome [vim-go][vim-go] package.


[fatih]:  https://github.com/fatih
[color]:  https://github.com/fatih/color
[vim-go]: https://github.com/fatih/vim-go
[mattn]:  https://github.com/mattn
[colorable]: https://github.com/mattn/go-colorable
[stdout]: https://github.com/golang/go/issues/13473
