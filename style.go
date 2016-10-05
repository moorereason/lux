package lux

// CodeFn represents a style code function.
type CodeFn func(interface{}) *Lux

//
// Basic Styles
//

// Reset writes the reset / normal escape code.
func Reset(v interface{}) *Lux { return add(reset, v) }

// Bold writes the bold code.
func Bold(v interface{}) *Lux { return add(bold, v) }

// Faint writes the faint code.
func Faint(v interface{}) *Lux { return add(faint, v) }

// Italic writes the italic code.
func Italic(v interface{}) *Lux { return add(italic, v) }

// Underline writes the underline code.
func Underline(v interface{}) *Lux { return add(underline, v) }

// BlinkSlow writes the blink code.
func BlinkSlow(v interface{}) *Lux { return add(blinkSlow, v) }

// BlinkRapid writes the rapid blink code.
func BlinkRapid(v interface{}) *Lux { return add(blinkRapid, v) }

// ReverseVideo writes the reverse video code.
func ReverseVideo(v interface{}) *Lux { return add(reverseVideo, v) }

// Conceal writes the conceal code.
func Conceal(v interface{}) *Lux { return add(conceal, v) }

// CrossedOut writes the crossed-out code.
func CrossedOut(v interface{}) *Lux { return add(crossedOut, v) }

//
// Foreground Colors
//

// Black writes the black foreground code.
func Black(v interface{}) *Lux { return add(fgBlack, v) }

// Red writes the red foreground code.
func Red(v interface{}) *Lux { return add(fgRed, v) }

// Green writes the green foreground code.
func Green(v interface{}) *Lux { return add(fgGreen, v) }

// Yellow writes the yellow foreground code.
func Yellow(v interface{}) *Lux { return add(fgYellow, v) }

// Blue writes the blue foreground code.
func Blue(v interface{}) *Lux { return add(fgBlue, v) }

// Magenta writes the magenta foreground code.
func Magenta(v interface{}) *Lux { return add(fgMagenta, v) }

// Cyan writes the cyan foreground code.
func Cyan(v interface{}) *Lux { return add(fgCyan, v) }

// White writes the white foreground code.
func White(v interface{}) *Lux { return add(fgWhite, v) }

//
// Background Colors
//

// BgBlack writes the black background code.
func BgBlack(v interface{}) *Lux { return add(bgBlack, v) }

// BgRed writes the red background code.
func BgRed(v interface{}) *Lux { return add(bgRed, v) }

// BgGreen writes the green background code.
func BgGreen(v interface{}) *Lux { return add(bgGreen, v) }

// BgYellow writes the yellow background code.
func BgYellow(v interface{}) *Lux { return add(bgYellow, v) }

// BgBlue writes the blue background code.
func BgBlue(v interface{}) *Lux { return add(bgBlue, v) }

// BgMagenta writes the magenta background code.
func BgMagenta(v interface{}) *Lux { return add(bgMagenta, v) }

// BgCyan writes the cyan background code.
func BgCyan(v interface{}) *Lux { return add(bgCyan, v) }

// BgWhite writes the white background code.
func BgWhite(v interface{}) *Lux { return add(bgWhite, v) }

//
// High-Intensity Foreground Colors
//

// HiBlack writes the high-intensity, black foreground code.
func HiBlack(v interface{}) *Lux { return add(fgHiBlack, v) }

// HiRed writes the high-intensity, red foreground code.
func HiRed(v interface{}) *Lux { return add(fgHiRed, v) }

// HiGreen writes the high-intensity, green foreground code.
func HiGreen(v interface{}) *Lux { return add(fgHiGreen, v) }

// HiYellow writes the high-intensity, yellow foreground code.
func HiYellow(v interface{}) *Lux { return add(fgHiYellow, v) }

// HiBlue writes the high-intensity, blue foreground code.
func HiBlue(v interface{}) *Lux { return add(fgHiBlue, v) }

// HiMagenta writes the high-intensity, magenta foreground code.
func HiMagenta(v interface{}) *Lux { return add(fgHiMagenta, v) }

// HiCyan writes the high-intensity, cyan foreground code.
func HiCyan(v interface{}) *Lux { return add(fgHiCyan, v) }

// HiWhite writes the high-intensity, white foreground code.
func HiWhite(v interface{}) *Lux { return add(fgHiWhite, v) }

//
// High-Intensity Background Colors
//

// BgHiBlack writes the high-intensity, black background code.
func BgHiBlack(v interface{}) *Lux { return add(bgHiBlack, v) }

// BgHiRed writes the high-intensity, red background code.
func BgHiRed(v interface{}) *Lux { return add(bgHiRed, v) }

// BgHiGreen writes the high-intensity, green background code.
func BgHiGreen(v interface{}) *Lux { return add(bgHiGreen, v) }

// BgHiYellow writes the high-intensity, yellow background code.
func BgHiYellow(v interface{}) *Lux { return add(bgHiYellow, v) }

// BgHiBlue writes the high-intensity, blue background code.
func BgHiBlue(v interface{}) *Lux { return add(bgHiBlue, v) }

// BgHiMagenta writes the high-intensity, magenta background code.
func BgHiMagenta(v interface{}) *Lux { return add(bgHiMagenta, v) }

// BgHiCyan writes the high-intensity, cyan background code.
func BgHiCyan(v interface{}) *Lux { return add(bgHiCyan, v) }

// BgHiWhite writes the high-intensity, white background code.
func BgHiWhite(v interface{}) *Lux { return add(bgHiWhite, v) }
