package program

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeon_carrier/internal/app"
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/urlinput"
)

type Program struct {
	currentFocus int
	urlInput     *urlinput.UrlInput
	method       *method.Method
	headers      *headers.Headers
}

func NewProgram() Program {
	method := method.NewMethod()
	urlInput := urlinput.NewUrlInput(method)
	headers := headers.NewHeaders()
	program := Program{
		currentFocus: 0,
		method:       method,
		urlInput:     urlInput,
		headers:      headers,
	}
	program.setFocus()

	return program
}

func (m Program) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Program) focusables() []app.Focusable {
	focusables := make([]app.Focusable, 0)
	focusables = append(focusables, m.method, m.urlInput)

	for _, focusable := range m.headers.GetFocusables() {
		focusables = append(focusables, focusable)
	}
	return focusables
}

func (m *Program) setFocus() (tea.Model, tea.Cmd) {
	return m.incrementAndSetFocus(0)
}

func (m *Program) incrementAndSetFocus(increment int) (tea.Model, tea.Cmd) {
	focusables := m.focusables()
	length := len(focusables)

	if increment >= 0 {
		m.currentFocus = (m.currentFocus + increment) % length

	} else {
		m.currentFocus = (m.currentFocus + increment + length) % length
	}

	var cmd []tea.Cmd

	for i, f := range focusables {
		if i == m.currentFocus {
			cmd = append(cmd, f.Focus())

		} else {
			cmd = append(cmd, f.Blur())
		}
	}

	return m, tea.Batch(cmd...)
}
