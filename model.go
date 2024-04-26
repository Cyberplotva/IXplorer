package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	dirPath    string
	dirEntries table.Model
	help       help.Model
	quitting   bool
}

func newModel() *model {
	startDirAbsPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Error getting absolute path of start directory: ", err)
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
		help:       help.New(),
		quitting:   false,
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
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Sequence(tea.ClearScreen, tea.Quit)

		case key.Matches(msg, keys.Right):
			if len(m.dirEntries.Rows()) == 0 {
				return m, nil
			}
			entryName := m.dirEntries.SelectedRow()[1]
			entryPath := filepath.Join(m.dirPath, entryName)
			fi, err := os.Stat(entryPath)
			if err != nil {
				log.Fatal("Error reading file info: ", err)
			}

			if fi.IsDir() {
				if _, err := os.ReadDir(entryPath); err == nil {
					storage.cursorPosition[m.dirPath] = m.dirEntries.Cursor()
					m.dirPath = filepath.Join(m.dirPath, entryName)
					return m, getNewRowsForDirEntries(m.dirPath)
				}
			}
			return m, nil

		case key.Matches(msg, keys.Left):
			storage.cursorPosition[m.dirPath] = m.dirEntries.Cursor()
			m.dirPath = filepath.Dir(m.dirPath)
			return m, getNewRowsForDirEntries(m.dirPath)
		
		case key.Matches(msg, keys.Help):
			before := lipgloss.Height(m.help.View(keys))
			m.help.ShowAll = !m.help.ShowAll
			diff := lipgloss.Height(m.help.View(keys)) - before

			if m.dirEntries.Cursor()+diff >= m.dirEntries.Height() {
				m.dirEntries.SetCursor(m.dirEntries.Height() - diff - 1)
			}
			m.dirEntries.SetHeight(m.dirEntries.Height() - diff)
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
	if m.quitting {
		return ""
	}

	sb := new(strings.Builder)

	sb.WriteString(
		lipglossStyleUpperBlock.Render(
			m.dirPath) + "\n")

	sb.WriteString(
		lipglossStyleMiddleBlock.Render(
			m.dirEntries.View()) + "\n")

	sb.WriteString(
		lipglossStyleBottomBlock.Render(
			m.help.View(keys)))

	return sb.String()
}
