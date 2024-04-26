package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin/go-humanize"
)

type newRowsForDirEntriesMsg []table.Row

func getNewRowsForDirEntries(dirPath string) tea.Cmd {
	return func() tea.Msg {
		dirEntries, err := os.ReadDir(dirPath)
		if err != nil {
			panic(fmt.Sprintf("Error reading directory: %v", err))
		}

		rows := make([]table.Row, len(dirEntries))
		for i, startDirEntry := range dirEntries {
			fi, err := startDirEntry.Info()
			if err != nil {
				panic(fmt.Sprintf("Error getting file info: %v", err))
			}

			rows[i] = make(table.Row, 4)
			rows[i][0] = "f"
			if fi.IsDir() {
				rows[i][0] = "d"
			} else {
				rows[i][3] = humanize.Bytes(uint64(fi.Size()))
			}
			rows[i][1] = fi.Name()
			rows[i][2] = fi.ModTime().Format("2006.02.01 15:04")
		}

		// Sort by date, type and then name
		sort.Slice(rows, func(i, j int) bool {
			if rows[i][2] != rows[j][2] {
				return rows[i][2] > rows[j][2]
			}
			if rows[i][0] != rows[j][0] {
				return rows[i][0] < rows[j][0]
			}
			return rows[i][1] < rows[j][1]
		})

		return newRowsForDirEntriesMsg(rows)
	}
}