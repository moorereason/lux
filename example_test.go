package lux_test

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/moorereason/lux"
)

func ExampleLux() {

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
