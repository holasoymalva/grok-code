package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

type model struct {
	messages []string
	input    string
	quitting bool
}

func InitialModel() tea.Model {
	return model{
		messages: []string{"Grok Code: Hello! How can I help you build today?"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.input != "" {
				m.messages = append(m.messages, "You: "+m.input)
				m.messages = append(m.messages, "Grok Code: Thinking... (not implemented)")
				m.input = ""
			}
			return m, nil
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	s := strings.Join(m.messages, "\n\n")
	s += "\n\n"
	s += "Input: " + m.input + "_"
	s += "\n\n(Press q to quit)"

	return baseStyle.Render(s)
}

func RunTUI() error {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return err
	}
	return nil
}
