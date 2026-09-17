package enterbox

import (
	tea "charm.land/bubbletea/v2"
)

func (b *EnterBox) Update(msg tea.Msg) tea.Cmd {
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

func (b *EnterBox) Focus() tea.Cmd {
	b.focused = true
	return nil
}

func (b *EnterBox) Blur() tea.Cmd {
	b.focused = false
	return nil
}
