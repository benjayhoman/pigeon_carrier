package option

import tea "charm.land/bubbletea/v2"

func (o *Option) Update(msg tea.Msg) tea.Cmd {
	if !o.focused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			o.selected = (o.selected - 1 + len(o.options)) % len(o.options)
		case "down":
			o.selected = (o.selected + 1) % len(o.options)
		}
	}
	return nil
}

func (o *Option) Focus() tea.Cmd {
	o.focused = true
	return nil
}

func (o *Option) Blur() tea.Cmd {
	o.focused = false
	return nil
}
