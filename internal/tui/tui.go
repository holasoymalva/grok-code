package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/holasoymalva/grok-code/internal/agent"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

type model struct {
	agent    *agent.Agent
	messages []string
	input    string
	quitting bool
	loading  bool
	spin     spinner.Model
}

func InitialModel(ag *agent.Agent) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		agent:    ag,
		messages: []string{"Grok Code: Hello! How can I help you build today?"},
		spin:     s,
	}
}

type responseMsg string
type errMsg error

func (m model) Init() tea.Cmd {
	return m.spin.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.input == "/exit" || m.input == "/quit" {
				m.quitting = true
				return m, tea.Quit
			}
			if m.input != "" && !m.loading {
				userText := m.input
				m.messages = append(m.messages, "You: "+userText)
				m.input = ""
				m.loading = true

				return m, tea.Batch(
					m.spin.Tick,
					func() tea.Msg {
						reply, err := m.agent.RunLoop(context.Background(), userText)
						if err != nil {
							return errMsg(err)
						}
						return responseMsg(reply)
					},
				)
			}
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(msg.String()) == 1 && !m.loading {
				m.input += msg.String()
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		if m.loading {
			m.spin, cmd = m.spin.Update(msg)
			return m, cmd
		}
		return m, nil

	case responseMsg:
		m.loading = false
		m.messages = append(m.messages, "Grok Code: "+string(msg))
		return m, nil

	case errMsg:
		m.loading = false
		m.messages = append(m.messages, "Grok Code: Error API: "+msg.Error())
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	s := strings.Join(m.messages, "\n\n")
	s += "\n\n"

	if m.loading {
		s += m.spin.View() + " Grok Code is thinking...\n\n"
	}

	s += "Input: " + m.input + "_"
	s += "\n\n(Type /exit or press ctrl+c to quit)"

	return baseStyle.Render(s)
}

func RunTUI(ag *agent.Agent) error {
	p := tea.NewProgram(InitialModel(ag))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return err
	}
	return nil
}
