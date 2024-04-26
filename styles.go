package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// ANSI color codes
const (
	ANSIColorLightGrey   = "241"
	ANSIColorLightYellow = "229"
	ANSIColorBlue        = "57"
)

// Lipgloss styles
var borderMiddle = lipgloss.Border{
	Top:          "─",
	Bottom:       "─",
	Left:         "│",
	Right:        "│",
	TopLeft:      "├", // Changed from default
	TopRight:     "┤", // Changed from default
	BottomLeft:   "├", // Changed from default
	BottomRight:  "┤", // Changed from default
	MiddleLeft:   "├",
	MiddleRight:  "┤",
	Middle:       "┼",
	MiddleTop:    "┬",
	MiddleBottom: "┴",
}

const terminalSize = 78

var (
	// Path
	lipglossStyleUpperBlock = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), true, true, false, true).
				Width(terminalSize).
				Bold(true)

	// Directory entries
	lipglossStyleMiddleBlock = lipgloss.NewStyle().
					BorderStyle(borderMiddle).
					Width(terminalSize)

	// Help
	lipglossStyleBottomBlock = lipgloss.NewStyle().
					Border(lipgloss.NormalBorder(), false, true, true, true).
					Width(terminalSize)

	lipglossStyleTable = table.DefaultStyles()
)

func init() {
	lipglossStyleTable.Header = lipglossStyleTable.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Bold(false)
	lipglossStyleTable.Selected = lipglossStyleTable.Selected.
		Foreground(lipgloss.Color(ANSIColorLightYellow)).
		Background(lipgloss.Color(ANSIColorBlue)).
		Bold(false)
}
