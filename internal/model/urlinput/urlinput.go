package urlinput

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/pigeon_carrier/internal/app"
	"github.com/pigeon_carrier/internal/model/method"
	modeltextinput "github.com/pigeon_carrier/internal/model/textinput"
	"github.com/rs/zerolog"
)

const (
	urlPlaceholder = "https://example.com/path?key=value"
)

var (
	styleCursor = lipgloss.NewStyle().Background(lipgloss.Color("15")).Foreground(lipgloss.Color("0")).Bold(true)
)

type UrlInput struct {
	log       zerolog.Logger
	method    *method.Method
	textInput *modeltextinput.TextInput
}

func (m UrlInput) Init() tea.Cmd {
	return nil
}

func NewUrlInput(logger zerolog.Logger) *UrlInput {
	method := method.NewMethod(logger)

	return &UrlInput{
		log:       logger.With().Str("module", "urlinput").Logger(),
		method:    method,
		textInput: modeltextinput.NewTextInput(urlPlaceholder, 512, 60),
	}
}

func (m UrlInput) Value() string {
	return m.textInput.Value()
}

func (m UrlInput) GetFocusables() []app.Focusable {
	return []app.Focusable{m.method, m.textInput}
}
