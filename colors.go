package markdown

import (
	"github.com/gookit/color"
)

var (
	// we need a bunch of escape code for manual formatting
	boldOn = "\x1b[1m"
	// boldOff       = "\x1b[21m" --> use resetAll + snapshot with bold off instead
	italicOn     = "\x1b[3m"
	crossedOutOn = "\x1b[9m"
	greenOn      = "\x1b[32m"
	resetAll     = "\x1b[0m"

	Green        = color.New(color.FgGreen).Sprint
	HiGreen      = color.New(color.FgLightGreen).Sprint
	GreenBold    = color.New(color.FgGreen, color.Bold).Sprint
	Blue         = color.New(color.FgBlue).Sprint
	BlueBgItalic = color.New(color.BgBlue, color.OpItalic).Sprint
	Red          = color.New(color.FgRed).Sprint
)
