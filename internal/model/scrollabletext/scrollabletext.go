package scrollabletext

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type ScrollableText struct {
	log      zerolog.Logger
	viewport viewport.Model
	focused  bool
}

func NewScrollableText(log zerolog.Logger) ScrollableText {
	return ScrollableText{
		log:      log,
		viewport: NewTextArea(),
	}
}

func NewTextArea() viewport.Model {
	ta := viewport.New(
		viewport.WithWidth(20),
		viewport.WithHeight(20),
	)
	return ta
}

func (r ScrollableText) Init() tea.Cmd {
	return nil
}

func (r *ScrollableText) SetContent(content string) {
	r.viewport.SetContent(content)
}

func (r *ScrollableText) SetWidth(width int) {
	r.viewport.SetWidth(width)
}

func (r *ScrollableText) SetHeight(height int) {
	r.viewport.SetHeight(height)
}
