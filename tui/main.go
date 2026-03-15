package tui

import (
	tea "charm.land/bubbletea/v2"
)

type focusedModel int

const (
	menu focusedModel = iota
	chat
)

type MainModel struct {
	State focusedModel
	Menu  tea.Model
	Chat  tea.Model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch m.State {
	case menu:
		menuModel, menuCmd := m.Menu.Update(msg)
		m.Menu = menuModel
		cmd = menuCmd
	case chat:
		chatModel, chatCmd := m.Chat.Update(msg)
		m.Chat = chatModel
		cmd = chatCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() tea.View {
	v := tea.NewView("loading...")
	switch m.State {
	case menu:
		v = m.Menu.View()
	case chat:
		v = m.Chat.View()
	}
	v.AltScreen = true
	return v
}
