package styles

import (
	"charm.land/lipgloss/v2"
)

var (
	colorWhite = lipgloss.Color("#FFFFFF")
	colorBlue  = lipgloss.Color("#0C3DB9")
	colorGreen = lipgloss.Color("#2AE173")
	colorError = lipgloss.Color("#ed1010")
)

var ChecksStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorWhite))

var ChecksTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorGreen)).
	Align(lipgloss.Left)

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorWhite)).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground((colorBlue))

var HeaderTreeRootStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorWhite))

var HeaderTreeChildStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorBlue))

var InfoStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground(colorBlue)

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorError))
