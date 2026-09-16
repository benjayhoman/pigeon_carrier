package app

import tea "charm.land/bubbletea/v2"

type Focusable interface {
	Focus() tea.Cmd
	Blur() tea.Cmd
}
