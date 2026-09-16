package action

import tea "github.com/charmbracelet/bubbletea"

type AddHeader struct{}

func NewAddHeader() tea.Msg {
	return AddHeader{}
}
