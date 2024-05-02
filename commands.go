package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin/go-humanize"
)


func getRowsFromDir(dirPath string) []table.Row {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("Error reading directory: ", err)
	}

	rows := make([]table.Row, len(dirEntries))
	for i, startDirEntry := range dirEntries {
		fi, err := startDirEntry.Info()
		if err != nil {
			log.Fatal("Error getting file info: ", err)
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

	return rows
}

func updateCursorPositionsForParentDirs(currDirPath string) {
	storage.cursorPosition.Store(currDirPath, 0)
	
	for {
		parentDirPath := filepath.Dir(currDirPath)
		if parentDirPath == currDirPath {
			break
		}
		currDirBase := filepath.Base(currDirPath)

		rows := getRowsFromDir(parentDirPath)
		for i, row := range rows {
			if row[1] == currDirBase {
				storage.cursorPosition.Store(parentDirPath, i)
				break
			}
		}

		currDirPath = parentDirPath
	}
}

type newDirEntriesMsg struct {
	rows []table.Row
}

func getNewRowsForDirEntries(dirPath string) tea.Cmd {
	return func() tea.Msg {
		return newDirEntriesMsg{
			rows: getRowsFromDir(dirPath),
		}
	}
}
