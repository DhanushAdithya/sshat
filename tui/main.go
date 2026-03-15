package tui

import (
	tea "charm.land/bubbletea/v2"
)

type focusedModel int

const (
	menu focusedModel = iota
	chat
)

type SwitchView struct{}

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

	switch msg.(type) {
	case SwitchView:
		switch m.State {
		case menu:
			m.State = chat
			return m, m.Chat.Init()
		case chat:
			m.State = menu
			return m, m.Menu.Init()
		}
	}

	switch m.State {
	case menu:
		var updatedMenu tea.Model
		updatedMenu, cmd = m.Menu.Update(msg)
		m.Menu = updatedMenu
	case chat:
		var updatedChat tea.Model
		updatedChat, cmd = m.Chat.Update(msg)
		m.Chat = updatedChat
	}
	return m, cmd
}

func (m MainModel) View() tea.View {
	var v tea.View
	switch m.State {
	case menu:
		v = m.Menu.View()
	case chat:
		v = m.Chat.View()
	}
	v.AltScreen = true
	return v
}
