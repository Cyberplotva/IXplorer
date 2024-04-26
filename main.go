package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := newModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(fmt.Sprintf("Error starting program: %v", err))
	}
}
