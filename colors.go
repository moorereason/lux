package lux

// Basic styles
const (
	reset int = iota
	bold
	faint
	italic
	underline
	blinkSlow
	blinkRapid
	reverseVideo
	conceal
	crossedOut
)

// Foreground colors
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

// Background Colors
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

// High-Intensity, Foreground Colors
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

// High-Intensity, Background Colors
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
	reset:        {0x30},
	bold:         {0x31},
	faint:        {0x32},
	italic:       {0x33},
	underline:    {0x34},
	blinkSlow:    {0x35},
	blinkRapid:   {0x36},
	reverseVideo: {0x37},
	conceal:      {0x38},
	crossedOut:   {0x39},

	fgBlack:   {0x33, 0x30},
	fgRed:     {0x33, 0x31},
	fgGreen:   {0x33, 0x32},
	fgYellow:  {0x33, 0x33},
	fgBlue:    {0x33, 0x34},
	fgMagenta: {0x33, 0x35},
	fgCyan:    {0x33, 0x36},
	fgWhite:   {0x33, 0x37},

	bgBlack:   {0x34, 0x30},
	bgRed:     {0x34, 0x31},
	bgGreen:   {0x34, 0x32},
	bgYellow:  {0x34, 0x33},
	bgBlue:    {0x34, 0x34},
	bgMagenta: {0x34, 0x35},
	bgCyan:    {0x34, 0x36},
	bgWhite:   {0x34, 0x37},

	fgHiBlack:   {0x39, 0x30},
	fgHiRed:     {0x39, 0x31},
	fgHiGreen:   {0x39, 0x32},
	fgHiYellow:  {0x39, 0x33},
	fgHiBlue:    {0x39, 0x34},
	fgHiMagenta: {0x39, 0x35},
	fgHiCyan:    {0x39, 0x36},
	fgHiWhite:   {0x39, 0x37},

	bgHiBlack:   {0x31, 0x30, 0x30},
	bgHiRed:     {0x31, 0x30, 0x31},
	bgHiGreen:   {0x31, 0x30, 0x32},
	bgHiYellow:  {0x31, 0x30, 0x33},
	bgHiBlue:    {0x31, 0x30, 0x34},
	bgHiMagenta: {0x31, 0x30, 0x35},
	bgHiCyan:    {0x31, 0x30, 0x36},
	bgHiWhite:   {0x31, 0x30, 0x37},
}
