package scrollabletext

import (
	tea "charm.land/bubbletea/v2"
)

func (r *ScrollableText) Update(msg tea.Msg) tea.Cmd {
	if !r.focused {
		return nil
	}

	var cmd tea.Cmd
	r.viewport, cmd = r.viewport.Update(msg)
	return cmd
}

func (m *ScrollableText) Focus() tea.Cmd {
	m.focused = true
	return nil
}

func (m *ScrollableText) Blur() tea.Cmd {
	m.focused = false
	return nil
}
