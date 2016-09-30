package lux

import (
	"bytes"
	"fmt"
	"os"

	isatty "github.com/mattn/go-isatty"
)

var (
	// NoColor defines whether to use escape sequences or not.  By default,
	// escape sequences are enabled when running on a TTY.
	NoColor bool = !isatty.IsTerminal(os.Stdout.Fd())

	escape = []byte{0x1b} // "\x1b"
)

// Lux represents the state of a particular stylized string.
type Lux struct {
	val    string
	styles []int
}

// New creates a new Lux instances with a series of LuxFn styles.
func New(fns ...LuxFn) *Lux {
	x := &Lux{}
	for _, fn := range fns {
		x = fn(x)
	}
	return x
}

// Add appends a style to the list of styles in a Lux.
func (x *Lux) Add(f LuxFn) *Lux {
	return f(x)
}

// Print is a colorable frontend to fmt.Print.
func (x *Lux) Print(a ...interface{}) (n int, err error) {
	mod := make([]interface{}, len(a))
	for i, v := range a {
		mod[i] = &Lux{
			val:    fmt.Sprintf("%v", v),
			styles: x.styles,
		}
	}

	return fmt.Fprintf(Output, "%s", mod...)
}

// Printf is a colorable frontend to fmt.Printf.
func (x *Lux) Printf(format string, a ...interface{}) (n int, err error) {
	mod := make([]interface{}, len(a))
	for i, v := range a {
		mod[i] = &Lux{
			val:    fmt.Sprintf("%v", v),
			styles: x.styles,
		}
	}

	return fmt.Fprintf(Output, format, mod...)
}

// Println is a colorable frontend to fmt.Println.
func (x *Lux) Println(a ...interface{}) (n int, err error) {
	mod := make([]interface{}, len(a))
	for i, v := range a {
		mod[i] = &Lux{
			val:    fmt.Sprintf("%v", v),
			styles: x.styles,
		}
	}

	return fmt.Fprintf(Output, "%s\n", mod...)
}

// Format implements the fmt.Formatter interface.  Not to be used directly.
// Instead, pass the Lux as a parameter to fmt format functions.
func (x *Lux) Format(f fmt.State, format rune) {
	if NoColor {
		f.Write([]byte(x.val))
		return
	}
	if format == 'v' {
		if f.Flag('+') {
			f.Write([]byte(fmt.Sprintf("{val: %v, styles: %v}", x.val, x.styles)))

		} else {
			f.Write([]byte(fmt.Sprintf("{%v %v}", x.val, x.styles)))
		}
		return
	}

	x.write(f)
}

// make sure we implement fmt.Formatter interface
var _ fmt.Formatter = (*Lux)(nil)

func (x *Lux) write(f fmt.State) {
	f.Write(x.format())
	f.Write([]byte(x.val))
	f.Write(x.unformat())
}

// format writes the beginning format codes.
func (x *Lux) format() []byte {
	var f bytes.Buffer

	f.Write(escape)
	f.Write([]byte{0x5b}) // "["

	for i := 0; i < len(x.styles); i++ {
		f.Write(attrs[x.styles[i]])
		if i < len(x.styles)-1 {
			f.Write([]byte{0x3b}) // ";"
		}
	}

	f.Write([]byte{0x6d}) // "m"

	return f.Bytes()
}

// unformat writes the format reset codes.
func (x *Lux) unformat() []byte {
	var f bytes.Buffer
	f.Write([]byte{0x1b, 0x5b, 0x30, 0x6d}) // "\033[0m"
	return f.Bytes()
}

// add is a helper function that creates a new Lux or extends an existing Lux.
func add(style int, v interface{}) (res *Lux) {
	switch t := v.(type) {
	case *Lux:
		res = t
		res.styles = append(res.styles, style)
	case string:
		res = &Lux{val: t, styles: []int{style}}
	default:
		if str, ok := t.(fmt.Stringer); ok {
			res = &Lux{val: str.String(), styles: []int{style}}
		} else {
			res = &Lux{val: fmt.Sprintf("%v", t), styles: []int{style}}
		}
	}
	return res
}
