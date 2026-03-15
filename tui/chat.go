package tui

import (
	"time"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type ChatModel struct {
	lastEscAt time.Time
	chat      []string
	input     textinput.Model
}

type chatKeyMap struct {
	ScrollUp   key.Binding
	ScrollDown key.Binding
	Skip       key.Binding
	Back       key.Binding
	Quit       key.Binding
}

var chatKeys = chatKeyMap{
	ScrollUp: key.NewBinding(
		key.WithKeys("ctrl+up", "ctrl+k"),
		key.WithHelp("ctrl+↑/ctrl+k", "to scroll"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys("ctrl+down", "ctrl+j"),
		key.WithHelp("ctrl+↓/ctrl+j", "to scroll"),
	),
	Skip: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc ×2", "to skip"),
	),
	Back: key.NewBinding(
		key.WithKeys("ctrl+x"),
		key.WithHelp("ctrl+x", "back to menu"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "to quit"),
	),
}

func (k chatKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		chatKeys.ScrollUp, chatKeys.ScrollDown,
		chatKeys.Skip, chatKeys.Back, chatKeys.Quit,
	}
}

func (k chatKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{chatKeys.ScrollUp, chatKeys.ScrollDown},
		{chatKeys.Skip, chatKeys.Back, chatKeys.Quit},
	}
}

func (m ChatModel) Init() tea.Cmd {
	return nil
}

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, chatKeys.Back):
			return m, func() tea.Msg { return SwitchView{} }
		case key.Matches(msg, chatKeys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ChatModel) View() tea.View {
	return tea.NewView("chat")
}
