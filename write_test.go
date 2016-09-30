// +build !windows

package lux

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrints(t *testing.T) {
	tests := []struct {
		text   string
		styles []LuxFn
		want   string
	}{
		{"AA", []LuxFn{Red, BgGreen}, "\x1b[31;42mAA\x1b[0m"},
		{"BB", []LuxFn{Bold}, "\x1b[1mBB\x1b[0m"},
		{"BB", []LuxFn{Bold, Red, BgGreen}, "\x1b[1;31;42mBB\x1b[0m"},
		{"BB", []LuxFn{BgGreen, Red, Bold}, "\x1b[42;31;1mBB\x1b[0m"},
	}

	for i, test := range tests {
		c := New(test.styles...)

		// Print
		out := captureOutput(func() { c.Print(test.text) })

		if out != test.want {
			t.Errorf("[%d] Print failed, want:\n\t%q\ngot:\n\t%q", i, test.want, out)
		}

		// Println
		out = captureOutput(func() { c.Println(test.text) })
		want := test.want + "\n"
		if out != want {
			t.Errorf("[%d] Println failed, want:\n\t%q\ngot:\n\t%q", i, want, out)
		}

		// Printf
		out = captureOutput(func() { c.Printf("A%sZ", test.text) })
		want = "A" + test.want + "Z"
		if out != want {
			t.Errorf("[%d] Printf failed, want:\n\t%q\ngot:\n\t%q", i, want, out)
		}
	}

}

/*
	blue := Blue
	s := blue("test")
	fmt.Printf("%s\n", s)
	fmt.Println(s)

	test := Blue("Test")
	test.Add(BgWhite)
	fmt.Printf("%s\n", test)

	c := New(Red, BgGreen)
	c.Print("Print: Red on Green")
	println()
	c.Println("Println: Red on Green")
	c.Printf("%s", "Printf: Red on Green\n")

	Printf("%s\n", Red(BgGreen("Red on Green")))
}
*/

func captureOutput(f func()) string {
	old := Output
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	Output = w

	f()

	w.Close()
	Output = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
