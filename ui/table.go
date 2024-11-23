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
	Table        table.Model
	Search       textinput.Model
	OriginalRows []table.Row
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var modifiedTable bool

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			m.Search.Reset()

		case "enter":
			command := m.Table.SelectedRow()[1]
			
			cmdParts := strings.Fields(command)
			if len(cmdParts) == 0 {
				return m, tea.Quit
			}

			m.Search.SetValue(command)

			selectedCmd := exec.Command(cmdParts[0], cmdParts[1:]...)
			return m, tea.Sequence(
				tea.ExecProcess(selectedCmd, nil),
				tea.Quit,
			)

		case "up", "down":
			modifiedTable = true
			m.Table, _ = m.Table.Update(msg)

		default:
			m.Search.Focus()
		}
	}

	// Update search input
	var searchCmd tea.Cmd
	m.Search, searchCmd = m.Search.Update(msg)
	cmds = append(cmds, searchCmd)

	// Filter table based on search input
	filteredRows := []table.Row{}

	// If search input is empty, show all rows
	searchText := strings.ToLower(m.Search.Value())
	if searchText == "" {
		m.Table.SetRows(m.OriginalRows)
		return m, tea.Batch(cmds...)
	}

	for _, row := range m.OriginalRows {
		if strings.Contains(strings.ToLower(row[1]), searchText) {
			filteredRows = append(filteredRows, row)
		}
	}
	m.Table.SetRows(filteredRows)

	// Update table if it hasn't been modified
	if !modifiedTable {
		var tableCmd tea.Cmd
		m.Table, tableCmd = m.Table.Update(msg)
		cmds = append(cmds, tableCmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString(baseStyle.Render(m.Table.View()))
	s.WriteString("\n  ")
	s.WriteString(m.Table.HelpView())
	s.WriteString("\n  ")
	s.WriteString(m.Search.View())
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
	}
}
