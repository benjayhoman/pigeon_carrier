package option

import tea "charm.land/bubbletea/v2"

type Option struct {
	options  []string
	selected int
	focused  bool
}

func NewOption(options []string) *Option {
	return &Option{options: options}
}

func (o Option) Init() tea.Cmd {
	return nil
}

func (o Option) Value() string {
	return o.options[o.selected]
}
