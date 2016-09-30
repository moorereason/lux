package lux

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	NoColor = false

	tests := []struct {
		text   string
		fn     luxFn
		styles string
	}{
		{"bold", Bold, "1"},
		{"faint", Faint, "2"},
		{"italic", Italic, "3"},
		{"underline", Underline, "4"},
		{"blinkSlow", BlinkSlow, "5"},
		{"blinkRapid", BlinkRapid, "6"},
		{"reverseVideo", ReverseVideo, "7"},
		{"concealed", Concealed, "8"},
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

		fmt.Printf("%s\n", s) // visuals in verbose or error mode
		if q != wantq {
			t.Errorf("Given: %q\nwant: %q\ngot: %q", test.text, wantq, q)
		}
	}
}

func TestMulti(t *testing.T) {
	NoColor = false

	tests := []struct {
		text   string
		fns    []luxFn
		styles string
	}{
		{"BlackOnBgGreen", []luxFn{BgGreen, Black}, "42;30"},
		{"BgGreenUnderBlack", []luxFn{Black, BgGreen}, "30;42"},
		{"RedGreen", []luxFn{Green, Red}, "32;31"},

		{"Red", []luxFn{Red}, "31"},
		{"RedBold", []luxFn{Red, Bold}, "31;1"},
		{"HiRedBold", []luxFn{HiRed, Bold}, "91;1"},

		{"GreenFaint", []luxFn{Green, Faint}, "32;2"},
		{"YellowItalic", []luxFn{Yellow, Italic}, "33;3"},
		{"BlueUnderline", []luxFn{Blue, Underline}, "34;4"},
		{"MagentaBlinkSlow", []luxFn{Magenta, BlinkSlow}, "35;5"},
		{"CyanBlinkRapid", []luxFn{Cyan, BlinkRapid}, "36;6"},
		{"WhiteReverseVideo", []luxFn{White, ReverseVideo}, "37;7"},
		{"WhiteConcealed", []luxFn{White, Concealed}, "37;8"},
		{"WhiteCrossedOut", []luxFn{White, CrossedOut}, "37;9"},
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

		fmt.Printf("%s\n", v) // visuals in verbose or error mode
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
