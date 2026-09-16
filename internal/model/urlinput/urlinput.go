package urlinput

import (
	"github.com/pigeon_carrier/internal/model/method"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	urlPlaceholder = "https://example.com/path?key=value"
)

var (
	stylePrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	styleCursor = lipgloss.NewStyle().Background(lipgloss.Color("15")).Foreground(lipgloss.Color("0")).Bold(true)
)

type UrlInput struct {
	method    *method.Method
	textInput textinput.Model
	focused   bool
}

func (m UrlInput) Init() tea.Cmd {
	return nil
}

func NewUrlInput(method *method.Method) *UrlInput {
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
		method:    method,
		textInput: ti,
	}
}
