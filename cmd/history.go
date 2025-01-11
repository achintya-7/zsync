/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/achintya-7/zsync/db"
	"github.com/achintya-7/zsync/ui"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get the history of the commands you have run",
	Long:  `Get the history of the commands you have run. This command will show you the history of the commands you have run in the past.`,
	Run: func(cmd *cobra.Command, args []string) {
		history()
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}

func history() {
	// get all the commands from SQL

	// initialize the table and run the terminal table
	model := ui.NewModel(initTable())
	p := tea.NewProgram(model)
	_, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
	}
}

func getTerminalWidth() int {
	var ws struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	)
	if err != 0 {
		return 80 // default width
	}
	return int(ws.Col)
}

func initTable() table.Model {
	// Get terminal width
	width := getTerminalWidth()

	// Calculate column widths
	mainColumnWidth := int(float64(width) * 0.8)

	columns := []table.Column{
		{Title: "Rank", Width: 5},
		{Title: "Command", Width: mainColumnWidth},
		{Title: "Frequency", Width: 10},
	}

	rows := generateRowsFromSql()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return t
}

func generateRowsFromSql() []table.Row {
	rows := []table.Row{}

	commands, err := db.GetAllCommands()
	if err != nil {
		fmt.Println("Error getting all the commands")
		return rows
	}

	for i, command := range commands {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", i+1),
			command.Command,
			fmt.Sprintf("%d", command.Frequency),
		})
	}

	return rows
}
