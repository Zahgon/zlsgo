package zlog

import (
	"os"
	"strings"

	"github.com/sohaha/zlsgo/zutil"
)

// DisableColor DisableColor
var DisableColor = false

// Color Color
type Color int

// Format Format
type Format int

// Op Op
type Op int

const (
	// ColorBlack black
	ColorBlack Color = iota + 30
	// ColorRed gules
	ColorRed
	// ColorGreen green
	ColorGreen
	// ColorYellow yellow
	ColorYellow
	// ColorBlue blue
	ColorBlue
	// ColorMagenta magenta
	ColorMagenta
	// ColorCyan cyan
	ColorCyan
	// ColorWhite white
	ColorWhite
)

const (
	// ColorLightGrey light grey
	ColorLightGrey Color = iota + 90
	// ColorLightRed light red
	ColorLightRed
	// ColorLightGreen light green
	ColorLightGreen
	// ColorLightYellow light yellow
	ColorLightYellow
	// ColorLightBlue light blue
	ColorLightBlue
	// ColorLightMagenta light magenta
	ColorLightMagenta
	// ColorLightCyan lightcyan
	ColorLightCyan
	// ColorLightWhite light white
	ColorLightWhite
	// ColorDefault ColorDefault
	ColorDefault = 49
)
const (
	// OpReset Reset All Settings
	OpReset Op = iota
	// OpBold Bold
	OpBold
	// OpFuzzy Fuzzy (not all terminal emulators support it)
	OpFuzzy
	// OpItalic Italic (not all terminal emulators support it)
	OpItalic
	// OpUnderscore Underline
	OpUnderscore
	// OpBlink Twinkle
	OpBlink
	// OpFastBlink Fast scintillation (not widely supported)
	OpFastBlink
	// OpReverse Reversed Exchange Background and Foreground Colors
	OpReverse
	// OpConcealed Concealed
	OpConcealed
	// OpStrikethrough Deleted lines (not widely supported)
	OpStrikethrough
)

// OpTextWrap OpTextWrap
func OpTextWrap(op Op, text string) string { _ = "STUB: not implemented"; return "" }

// ColorBackgroundWrap ColorBackgroundWrap
func ColorBackgroundWrap(color Color, backgroundColor Color, text string) string {
	_ = "STUB: not implemented"
	return ""
}

// OutAllColor OutAllColor
func OutAllColor() { _ = "STUB: not implemented"; return }

// GetAllColorText GetAllColorText
func GetAllColorText() map[string]Color { _ = "STUB: not implemented"; return nil }

// ColorTextWrap ColorTextWrap
func ColorTextWrap(color Color, text string) string { _ = "STUB: not implemented"; return "" }

var supportColor bool
var isMsystem = os.Getenv("MSYSTEM") != ""

func init() {
	if zutil.IsWin() && isMsystem {
		return
	}
	term := os.Getenv("TERM")
	supportColor = strings.Contains(term, "xterm") || os.Getenv("ConEmuANSI") == "ON" || os.Getenv("ANSICON") != "" || strings.Contains(term, "256color")
}

func isSupportColor() bool { _ = "STUB: not implemented"; return false }

func TrimAnsi(str string) string { _ = "STUB: not implemented"; return "" }
