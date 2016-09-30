package lux

import (
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

// make sure we implement fmt.Formatter interface
var _ fmt.Formatter = (*Lux)(nil)

// Format implements the fmt.Formatter interface.
func (s *Lux) Format(state fmt.State, format rune) {
	if NoColor {
		state.Write([]byte(s.val))
		return
	}
	if format == 'v' {
		if state.Flag('+') {
			state.Write([]byte(fmt.Sprintf("{val: %v, styles: %v}", s.val, s.styles)))

		} else {
			state.Write([]byte(fmt.Sprintf("{%v %v}", s.val, s.styles)))
		}
		return
	}

	state.Write(escape)
	state.Write([]byte{0x5b}) // "["

	for i := 0; i < len(s.styles); i++ {
		state.Write(attrs[s.styles[i]])
		if i < len(s.styles)-1 {
			state.Write([]byte{0x3b}) // ";"
		}
	}

	state.Write([]byte{0x6d}) // "m"
	state.Write([]byte(s.val))
	state.Write([]byte{0x1b, 0x5b, 0x30, 0x6d}) // "\033[0m"
}

// add creates a new Lux.
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

type luxFn func(interface{}) *Lux

func Reset(v interface{}) *Lux        { return add(0, v) }
func Bold(v interface{}) *Lux         { return add(1, v) }
func Faint(v interface{}) *Lux        { return add(2, v) }
func Italic(v interface{}) *Lux       { return add(3, v) }
func Underline(v interface{}) *Lux    { return add(4, v) }
func BlinkSlow(v interface{}) *Lux    { return add(5, v) }
func BlinkRapid(v interface{}) *Lux   { return add(6, v) }
func ReverseVideo(v interface{}) *Lux { return add(7, v) }
func Concealed(v interface{}) *Lux    { return add(8, v) }
func CrossedOut(v interface{}) *Lux   { return add(9, v) }

func Black(v interface{}) *Lux   { return add(30, v) }
func Red(v interface{}) *Lux     { return add(31, v) }
func Green(v interface{}) *Lux   { return add(32, v) }
func Yellow(v interface{}) *Lux  { return add(33, v) }
func Blue(v interface{}) *Lux    { return add(34, v) }
func Magenta(v interface{}) *Lux { return add(35, v) }
func Cyan(v interface{}) *Lux    { return add(36, v) }
func White(v interface{}) *Lux   { return add(37, v) }

func BgBlack(v interface{}) *Lux   { return add(40, v) }
func BgRed(v interface{}) *Lux     { return add(41, v) }
func BgGreen(v interface{}) *Lux   { return add(42, v) }
func BgYellow(v interface{}) *Lux  { return add(43, v) }
func BgBlue(v interface{}) *Lux    { return add(44, v) }
func BgMagenta(v interface{}) *Lux { return add(45, v) }
func BgCyan(v interface{}) *Lux    { return add(46, v) }
func BgWhite(v interface{}) *Lux   { return add(47, v) }

func HiBlack(v interface{}) *Lux   { return add(90, v) }
func HiRed(v interface{}) *Lux     { return add(91, v) }
func HiGreen(v interface{}) *Lux   { return add(92, v) }
func HiYellow(v interface{}) *Lux  { return add(93, v) }
func HiBlue(v interface{}) *Lux    { return add(94, v) }
func HiMagenta(v interface{}) *Lux { return add(95, v) }
func HiCyan(v interface{}) *Lux    { return add(96, v) }
func HiWhite(v interface{}) *Lux   { return add(97, v) }

func BgHiBlack(v interface{}) *Lux   { return add(100, v) }
func BgHiRed(v interface{}) *Lux     { return add(101, v) }
func BgHiGreen(v interface{}) *Lux   { return add(102, v) }
func BgHiYellow(v interface{}) *Lux  { return add(103, v) }
func BgHiBlue(v interface{}) *Lux    { return add(104, v) }
func BgHiMagenta(v interface{}) *Lux { return add(105, v) }
func BgHiCyan(v interface{}) *Lux    { return add(106, v) }
func BgHiWhite(v interface{}) *Lux   { return add(107, v) }
