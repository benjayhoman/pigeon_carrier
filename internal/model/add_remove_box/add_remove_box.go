package addremovebox

import tea "charm.land/bubbletea/v2"

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
