package action

import tea "charm.land/bubbletea/v2"

type AddHeader struct{}

func NewAddHeader() tea.Msg {
	return AddHeader{}
}
