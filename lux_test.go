// +build !windows

package lux

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	NoColor = false

	tests := []struct {
		text   string
		fn     CodeFn
		styles string
	}{
		{"bold", Bold, "1"},
		{"faint", Faint, "2"},
		{"italic", Italic, "3"},
		{"underline", Underline, "4"},
		{"blinkSlow", BlinkSlow, "5"},
		{"blinkRapid", BlinkRapid, "6"},
		{"reverseVideo", ReverseVideo, "7"},
		{"conceal", Conceal, "8"},
		{"crossedOut", CrossedOut, "9"},

		{"black", Black, "30"},
		{"red", Red, "31"},
		{"green", Green, "32"},
		{"yellow", Yellow, "33"},
		{"blue", Blue, "34"},
		{"magenta", Magenta, "35"},
		{"cyan", Cyan, "36"},
		{"white", White, "37"},

		{"bgBlack", BgBlack, "40"},
		{"bgRed", BgRed, "41"},
		{"bgGreen", BgGreen, "42"},
		{"bgYellow", BgYellow, "43"},
		{"bgBlue", BgBlue, "44"},
		{"bgMagenta", BgMagenta, "45"},
		{"bgCyan", BgCyan, "46"},
		{"bgWhite", BgWhite, "47"},

		{"hiBlack", HiBlack, "90"},
		{"hiRed", HiRed, "91"},
		{"hiGreen", HiGreen, "92"},
		{"hiYellow", HiYellow, "93"},
		{"hiBlue", HiBlue, "94"},
		{"hiMagenta", HiMagenta, "95"},
		{"hiCyan", HiCyan, "96"},
		{"hiWhite", HiWhite, "97"},

		{"bgHiBlack", BgHiBlack, "100"},
		{"bgHiRed", BgHiRed, "101"},
		{"bgHiGreen", BgHiGreen, "102"},
		{"bgHiYellow", BgHiYellow, "103"},
		{"bgHiBlue", BgHiBlue, "104"},
		{"bgHiMagenta", BgHiMagenta, "105"},
		{"bgHiCyan", BgHiCyan, "106"},
		{"bgHiWhite", BgHiWhite, "107"},
	}

	for _, test := range tests {
		s := fmt.Sprintf("%s", test.fn(test.text))
		q := fmt.Sprintf("%q", s)
		want := fmt.Sprintf("%s[%sm%s%s[0m", escape, test.styles, test.text, escape)
		wantq := fmt.Sprintf("%q", want)

		// fmt.Printf("%s\n", s) // visuals in verbose or error mode
		if q != wantq {
			t.Errorf("Given: %q\nwant: %q\ngot: %q", test.text, wantq, q)
		}
	}
}

func TestMulti(t *testing.T) {
	NoColor = false

	tests := []struct {
		text   string
		fns    []CodeFn
		styles string
	}{
		{"BlackOnBgGreen", []CodeFn{BgGreen, Black}, "42;30"},
		{"BgGreenUnderBlack", []CodeFn{Black, BgGreen}, "30;42"},
		{"RedGreen", []CodeFn{Green, Red}, "32;31"},

		{"Red", []CodeFn{Red}, "31"},
		{"RedBold", []CodeFn{Red, Bold}, "31;1"},
		{"HiRedBold", []CodeFn{HiRed, Bold}, "91;1"},

		{"GreenFaint", []CodeFn{Green, Faint}, "32;2"},
		{"YellowItalic", []CodeFn{Yellow, Italic}, "33;3"},
		{"BlueUnderline", []CodeFn{Blue, Underline}, "34;4"},
		{"MagentaBlinkSlow", []CodeFn{Magenta, BlinkSlow}, "35;5"},
		{"CyanBlinkRapid", []CodeFn{Cyan, BlinkRapid}, "36;6"},
		{"WhiteReverseVideo", []CodeFn{White, ReverseVideo}, "37;7"},
		{"WhiteConceal", []CodeFn{White, Conceal}, "37;8"},
		{"WhiteCrossedOut", []CodeFn{White, CrossedOut}, "37;9"},
	}

	for _, test := range tests {
		var v *Lux
		for i, f := range test.fns {
			if i == 0 {
				v = f(test.text)
			} else {
				v = f(v)
			}
		}

		s := fmt.Sprintf("%s", v)
		q := fmt.Sprintf("%q", s)
		want := fmt.Sprintf("%s[%sm%s%s[0m", escape, test.styles, test.text, escape)
		wantq := fmt.Sprintf("%q", want)

		// fmt.Printf("%s\n", v) // visuals in verbose or error mode
		if q != wantq {
			t.Errorf("Given: %q\nwant: %q\ngot: %q", test.text, wantq, q)
		}
	}
}

func BenchmarkColor(b *testing.B) {
	NoColor = false
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", Blue("test"))
	}
}

func BenchmarkNoColor(b *testing.B) {
	NoColor = true
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", Blue("test"))
	}
}
