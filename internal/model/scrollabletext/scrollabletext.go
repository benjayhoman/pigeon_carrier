package scrollabletext

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

const (
	maxHeight = 20
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
		viewport.WithHeight(0),
	)
	return ta
}

func (r ScrollableText) Init() tea.Cmd {
	return nil
}

func (r *ScrollableText) SetContent(content string) {
	contentWithLineNumbers, numberOfLines := prependLineNumbers(content)
	r.viewport.SetContent(contentWithLineNumbers)
	r.setHeight(numberOfLines)
}

func (r *ScrollableText) SetWidth(width int) {
	r.viewport.SetWidth(width)
}

func (r *ScrollableText) setHeight(height int) {
	if height > maxHeight {
		height = maxHeight
	}
	r.viewport.SetHeight(height)
}

func prependLineNumbers(content string) (string, int) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%3d %s", i+1, line)
	}
	return strings.Join(lines, "\n"), len(lines)
}
