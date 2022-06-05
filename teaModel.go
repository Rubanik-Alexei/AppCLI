package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	taskName        string
	taskDescription string
	event           string
}

var _ tea.Model = (*model)(nil)

func NewModel() (*model, error) {
	return &model{}, nil
}

func (m model) View() string {
	if m.event != "" {
		return fmt.Sprintf("You've selected: %s", m.event)
	}
	return "TODO"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}
