package tui

import (
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type ChatModel struct {
	lastEscAt time.Time
	chat      []string
	input     textinput.Model
}

func (m ChatModel) Init() tea.Cmd {
	return nil
}

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ChatModel) View() tea.View {
	return tea.NewView("chat")
}
