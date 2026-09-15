package urlinput

import (
	"github.com/pigeon_carrier/internal/model/method"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	ti.Focus()
	ti.CharLimit = 512
	ti.Width = 60
	ti.PromptStyle = stylePrompt
	ti.Cursor.Style = styleCursor

	return &UrlInput{
		method:    method,
		textInput: ti,
	}
}
