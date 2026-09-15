package program

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/urlinput"
)

type Program struct {
	currentFocus int
	urlInput     *urlinput.UrlInput
	method       *method.Method
}

func NewProgram() Program {
	method := method.NewMethod()
	urlInput := urlinput.NewUrlInput(method)
	program := Program{
		currentFocus: 0,
		method:       method,
		urlInput:     urlInput,
	}
	program.updateFocus()

	return program
}

func (m Program) Init() tea.Cmd {
	return textinput.Blink
}
