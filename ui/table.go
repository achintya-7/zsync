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
	Done  bool
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Done = true
			return m, tea.Quit
		case "enter":
			m.Done = true
			return m, runCommand(m.Table.SelectedRow()[1])
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func runCommand(cmd string) tea.Cmd {
	execCommand := exec.Command(cmd)

	return tea.Sequence(
		tea.ExecProcess(exec.Command("clear"), nil),
		tea.ExecProcess(execCommand, nil),
		tea.ClearScreen,
		tea.Quit,
	)
}

func (m Model) View() string {
	return baseStyle.Render(m.Table.View()) + "\n  " + m.Table.HelpView() + "\n"
}

func (m Model) ClearView() string {
	return lipgloss.NewStyle().Render(m.Table.View())
}

// NewModel returns a new Model
func NewModel(table table.Model) Model {
	return Model{Table: table}
}
