package addremovebox

import (
	tea "charm.land/bubbletea/v2"
)

func (b *AddRemoveBox) Update(msg tea.Msg) tea.Cmd {
	if !b.focused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
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
