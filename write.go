package lux

import (
	"fmt"

	"github.com/mattn/go-colorable"
)

// Output is the io.Writer used by Print, Printf, and Println.
var Output = colorable.NewColorableStdout()

// Print is a colorable frontend to fmt.Print.
func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(Output, a...)
}

// Printf is a colorable frontend to fmt.Printf.
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Output, format, a...)
}

// Println is a colorable frontend to fmt.Println.
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Output, a...)
}
