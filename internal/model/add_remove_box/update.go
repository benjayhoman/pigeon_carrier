package addremovebox

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (b *AddRemoveBox) Update(msg tea.Msg) tea.Cmd {
	if !b.focused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if b.action != nil {
				return b.action.Do()
			}
		}
	}

	return nil
}

func (b *AddRemoveBox) Focus() tea.Cmd {
	b.focused = true
	return nil
}

func (b *AddRemoveBox) Blur() tea.Cmd {
	b.focused = false
	return nil
}
