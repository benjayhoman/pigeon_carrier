package urlinput

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/rs/zerolog"
)

const (
	urlPlaceholder = "https://example.com/path?key=value"
)

var (
	stylePrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	styleCursor = lipgloss.NewStyle().Background(lipgloss.Color("15")).Foreground(lipgloss.Color("0")).Bold(true)
)

type UrlInput struct {
	log       zerolog.Logger
	method    *method.Method
	textInput textinput.Model
	focused   bool
}

func (m UrlInput) Init() tea.Cmd {
	return nil
}

func NewUrlInput(logger zerolog.Logger, method *method.Method) *UrlInput {
	ti := textinput.New()
	ti.Placeholder = urlPlaceholder
	ti.SetVirtualCursor(true)
	ti.Focus()
	ti.CharLimit = 512
	ti.SetWidth(60)
	styles := ti.Styles()
	styles.Focused.Prompt = stylePrompt
	styles.Blurred.Prompt = stylePrompt
	styles.Cursor.Color = lipgloss.Color("15")
	ti.SetStyles(styles)

	return &UrlInput{
		log:       logger.With().Str("module", "urlinput").Logger(),
		method:    method,
		textInput: ti,
		focused:   false,
	}
}

func (m UrlInput) Value() string {
	return m.textInput.Value()
}
