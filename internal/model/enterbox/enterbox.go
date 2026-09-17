package enterbox

import tea "charm.land/bubbletea/v2"

type OnEnter interface {
	Do() tea.Cmd
}

type EnterBox struct {
	symbol  string
	focused bool
	action  OnEnter
}

func (b EnterBox) Init() tea.Cmd {
	return nil
}

func NewEnterBox(symbol string, action OnEnter) EnterBox {
	return EnterBox{
		symbol: symbol,
		action: action,
	}
}
