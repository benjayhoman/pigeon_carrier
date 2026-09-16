package addremovebox

import tea "github.com/charmbracelet/bubbletea"

type OnEnter interface {
	Do() tea.Cmd
}

type AddRemoveBox struct {
	symbol  string
	focused bool
	action  OnEnter
}

func (b AddRemoveBox) Init() tea.Cmd {
	return nil
}

func NewAddRemoveBox(symbol string, action OnEnter) AddRemoveBox {
	return AddRemoveBox{
		symbol: symbol,
		action: action,
	}
}
