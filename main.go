package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/DhanushAdithya/sshat/tui"
)

func main() {
	p := tea.NewProgram(tui.MainModel{
		State: 0,
		Menu:  tui.MenuModel{},
		Chat:  tui.ChatModel{},
	})

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
