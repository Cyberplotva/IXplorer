package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dirPath    string
	dirEntries table.Model
	isQuitting bool
}

func newModel() *model {
	startDirAbsPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Error getting absolute path of start directory: %v", err)
	}

	columns := []table.Column{
		{Title: "Type", Width: 4},
		{Title: "Name", Width: 30},
		{Title: "Date modified", Width: 20},
		{Title: "Size", Width: 16},
	}

	startDirEntries := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(16),
		table.WithStyles(lipglossStyleTable),
	)

	return &model{
		dirPath:    startDirAbsPath,
		dirEntries: startDirEntries,
		isQuitting: false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Sequence(
		tea.ClearScreen,
		getNewRowsForDirEntries(m.dirPath),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.isQuitting = true
			return m, tea.Sequence(tea.ClearScreen, tea.Quit)
		case "right":
			if len(m.dirEntries.Rows()) == 0 {
				return m, nil
			}
			entryName := m.dirEntries.SelectedRow()[1]
			entryPath := filepath.Join(m.dirPath, entryName)
			fi, err := os.Stat(entryPath)
			if err != nil {
				log.Fatalf("Error reading file info: %v", err)
			}
			
			if fi.IsDir() {
				if _, err := os.ReadDir(entryPath); err == nil {
					storage.cursorPosition[m.dirPath] = m.dirEntries.Cursor()
					m.dirPath = filepath.Join(m.dirPath, entryName)
					return m, getNewRowsForDirEntries(m.dirPath)
				}
			}
			return m, nil
		case "left":
			storage.cursorPosition[m.dirPath] = m.dirEntries.Cursor()
			m.dirPath = filepath.Dir(m.dirPath)
			return m, getNewRowsForDirEntries(m.dirPath)
		}
	case newDirEntriesMsg:
		m.dirEntries.SetRows(msg.rows)
		m.dirEntries.Focus()
		m.dirEntries.SetCursor(
			storage.cursorPosition[m.dirPath])
	}
	var cmd tea.Cmd
	m.dirEntries, cmd = m.dirEntries.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.isQuitting {
		return ""
	}

	sb := new(strings.Builder)

	sb.WriteString(
		lipglossStyleBorderedTop.Render(m.dirPath))

	sb.WriteString("\n")

	sb.WriteString(
		lipglossStyleBorderedBlock.Render(
			m.dirEntries.View()))

	sb.WriteString(
		lipglossStyleHelperText.Render("\nPress 'q' to quit"))

	return sb.String()
}
