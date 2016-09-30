package lux

// Base ints
const (
	reset int = iota
	bold
	faint
	italic
	underline
	blinkSlow
	blinkRapid
	reverseVideo
	concealed
	crossedOut
)

// Foreground text colors
const (
	fgBlack int = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

// Foreground Hi-Intensity text colors
const (
	fgHiBlack int = iota + 90
	fgHiRed
	fgHiGreen
	fgHiYellow
	fgHiBlue
	fgHiMagenta
	fgHiCyan
	fgHiWhite
)

// Background text colors
const (
	bgBlack int = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
)

// Background Hi-Intensity text colors
const (
	bgHiBlack int = iota + 100
	bgHiRed
	bgHiGreen
	bgHiYellow
	bgHiBlue
	bgHiMagenta
	bgHiCyan
	bgHiWhite
)

var attrs = [...][]byte{
	reset:        []byte{0x30},
	bold:         []byte{0x31},
	faint:        []byte{0x32},
	italic:       []byte{0x33},
	underline:    []byte{0x34},
	blinkSlow:    []byte{0x35},
	blinkRapid:   []byte{0x36},
	reverseVideo: []byte{0x37},
	concealed:    []byte{0x38},
	crossedOut:   []byte{0x39},

	fgBlack:   []byte{0x33, 0x30},
	fgRed:     []byte{0x33, 0x31},
	fgGreen:   []byte{0x33, 0x32},
	fgYellow:  []byte{0x33, 0x33},
	fgBlue:    []byte{0x33, 0x34},
	fgMagenta: []byte{0x33, 0x35},
	fgCyan:    []byte{0x33, 0x36},
	fgWhite:   []byte{0x33, 0x37},

	bgBlack:   []byte{0x34, 0x30},
	bgRed:     []byte{0x34, 0x31},
	bgGreen:   []byte{0x34, 0x32},
	bgYellow:  []byte{0x34, 0x33},
	bgBlue:    []byte{0x34, 0x34},
	bgMagenta: []byte{0x34, 0x35},
	bgCyan:    []byte{0x34, 0x36},
	bgWhite:   []byte{0x34, 0x37},

	fgHiBlack:   []byte{0x39, 0x30},
	fgHiRed:     []byte{0x39, 0x31},
	fgHiGreen:   []byte{0x39, 0x32},
	fgHiYellow:  []byte{0x39, 0x33},
	fgHiBlue:    []byte{0x39, 0x34},
	fgHiMagenta: []byte{0x39, 0x35},
	fgHiCyan:    []byte{0x39, 0x36},
	fgHiWhite:   []byte{0x39, 0x37},

	bgHiBlack:   []byte{0x31, 0x30, 0x30},
	bgHiRed:     []byte{0x31, 0x30, 0x31},
	bgHiGreen:   []byte{0x31, 0x30, 0x32},
	bgHiYellow:  []byte{0x31, 0x30, 0x33},
	bgHiBlue:    []byte{0x31, 0x30, 0x34},
	bgHiMagenta: []byte{0x31, 0x30, 0x35},
	bgHiCyan:    []byte{0x31, 0x30, 0x36},
	bgHiWhite:   []byte{0x31, 0x30, 0x37},
}
