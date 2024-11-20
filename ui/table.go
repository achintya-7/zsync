package ui

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

type Model struct {
	Table table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, runCommand(m.Table.SelectedRow()[1])
		case "tab":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func runCommand(cmd string) tea.Cmd {
	execCommand := exec.Command(cmd)

	return tea.Batch(
		tea.ExecProcess(execCommand, nil),
		tea.Quit,
	)
}

func (m Model) View() string {
	return baseStyle.Render(m.Table.View()) + "\n  " + m.Table.HelpView() + "\n"
}

// NewModel returns a new Model
func NewModel(table table.Model) Model {
	return Model{Table: table}
}
