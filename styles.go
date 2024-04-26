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

var specialBorder = lipgloss.Border{
	Top:          "─",
	Bottom:       "─",
	Left:         "│",
	Right:        "│",
	TopLeft:      "│", // Changed
	TopRight:     "│", // Changed
	BottomLeft:   "└",
	BottomRight:  "┘",
	MiddleLeft:   "├",
	MiddleRight:  "┤",
	Middle:       "┼",
	MiddleTop:    "┬",
	MiddleBottom: "┴",
}

var (
	lipglossStyleHelperText = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ANSIColorLightGrey)).
				Italic(true)
	lipglossStyleBorderedBlock = lipgloss.NewStyle().
					BorderStyle(specialBorder).
					BorderForeground(lipgloss.Color(ANSIColorLightGrey))
	lipglossStyleBorderedTop = lipgloss.NewStyle().
					Border(lipgloss.NormalBorder(), true, true, false, true).
					BorderForeground(lipgloss.Color(ANSIColorLightGrey)).
					// PaddingRight(75).
					Width(78)

	lipglossStyleTable = table.DefaultStyles()
)

func init() {
	lipglossStyleTable.Header = lipglossStyleTable.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(ANSIColorLightGrey)).
		BorderBottom(true)
	lipglossStyleTable.Selected = lipglossStyleTable.Selected.
		Foreground(lipgloss.Color(ANSIColorLightYellow)).
		Background(lipgloss.Color(ANSIColorBlue)).
		Bold(false)
}
