package tui

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type MenuModel struct {
	help help.Model
	keys menuKeyMap
	msg  string
}

type menuKeyMap struct {
	Enter    key.Binding
	Navigate key.Binding
	Chat     key.Binding
	Quit     key.Binding
}

var menuKeys = menuKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "to select"),
	),
	Navigate: key.NewBinding(
		key.WithKeys("up", "down", "k", "j"),
		key.WithHelp("↑/k ↓/j", "to navigate"),
	),
	Chat: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "to start chat"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "to quit"),
	),
}

func (k menuKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		menuKeys.Navigate, menuKeys.Enter, menuKeys.Chat, menuKeys.Quit,
	}
}

func (k menuKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{menuKeys.Navigate, menuKeys.Enter},
		{menuKeys.Chat, menuKeys.Quit},
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, menuKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, menuKeys.Chat):
			return m, func() tea.Msg { return SwitchView{} }
		}
	}
	return m, nil
}

func (m MenuModel) View() tea.View {
	return tea.NewView("menu")
}
