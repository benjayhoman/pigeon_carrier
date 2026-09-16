package program

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	Focus() tea.Cmd
	Blur() tea.Cmd
}
