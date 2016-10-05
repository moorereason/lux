package lux_test

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/moorereason/lux"
)

func Example() {
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
}

func ExamplePrint() {
	lux.Print(lux.Blue("Skies"))
}

func ExamplePrintln() {
	lux.Println(lux.Blue("Skies"))
}

func ExamplePrintf() {
	lux.Printf("%d\n", lux.HiCyan(20))
}

func ExampleLux_Print() {
	c := lux.New(lux.Red)
	c.Print("ERROR")
}

func ExampleLux_Println() {
	c := lux.New(lux.Yellow)
	c.Println("WARNING")
}

var err error

func ExampleLux_Printf() {
	c := lux.New(lux.BgRed)
	c.Printf("CRITICAL: %s\n", err)
}
