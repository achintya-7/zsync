package ui

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

type Model struct {
	Table         table.Model
	Search        textinput.Model
	Done          bool
	OriginalRows  []table.Row
	FocusedInput  string // can be "table" or "search"
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Done = true
			return m, tea.Quit
		case "enter":
			if m.FocusedInput == "table" {
				m.Done = true
				cmdParts := strings.Fields(m.Table.SelectedRow()[1])
				if len(cmdParts) == 0 {
					return m, tea.Quit
				}
				
				selectedCmd := exec.Command(cmdParts[0], cmdParts[1:]...)
				return m, tea.Sequence(
					tea.ExecProcess(selectedCmd, nil),
					tea.Quit,
				)
			}
		case "/":
			if m.FocusedInput == "table" {
				m.FocusedInput = "search"
				m.Search.Focus()
				return m, textinput.Blink
			}
		case "esc":
			if m.FocusedInput == "search" {
				m.FocusedInput = "table"
				m.Search.Blur()
				return m, nil
			}
		}

		if m.FocusedInput == "search" {
			var searchCmd tea.Cmd
			m.Search, searchCmd = m.Search.Update(msg)
			
			// Filter table based on search input
			filteredRows := []table.Row{}
			searchText := strings.ToLower(m.Search.Value())
			for _, row := range m.OriginalRows {
				if strings.Contains(strings.ToLower(row[1]), searchText) {
					filteredRows = append(filteredRows, row)
				}
			}
			m.Table.SetRows(filteredRows)
			
			return m, searchCmd
		}
	}

	if m.FocusedInput == "table" {
		m.Table, cmd = m.Table.Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString(baseStyle.Render(m.Table.View()))
	s.WriteString("\n  ")
	s.WriteString(m.Table.HelpView())
	s.WriteString("\n  ")
	
	// Add search prompt
	if m.FocusedInput == "search" {
		s.WriteString(m.Search.View())
	} else {
		s.WriteString("Press / to search")
	}
	s.WriteString("\n")

	return s.String()
}

func (m Model) ClearView() string {
	return lipgloss.NewStyle().Render(m.Table.View())
}

// NewModel returns a new Model
func NewModel(table table.Model) Model {
	ti := textinput.New()
	ti.Placeholder = "Search commands..."
	ti.CharLimit = 100
	ti.Width = 30

	return Model{
		Table:        table,
		Search:       ti,
		OriginalRows: table.Rows(),
		FocusedInput: "table",
	}
}