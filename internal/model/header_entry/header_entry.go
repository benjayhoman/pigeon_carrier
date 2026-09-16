package headerentry

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/add_remove_box"
)

const (
	keyPlaceholder   = "Key"
	valuePlaceholder = "Value"
)

var (
	stylePrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	styleCursor = lipgloss.NewStyle()

	MinWidth = 5
	MaxWidth = 30
)

type textInput struct {
	textinput.Model
}

func (t *textInput) Focus() tea.Cmd {
	return t.Model.Focus()
}

func (t *textInput) Blur() tea.Cmd {
	t.Model.Blur()
	return nil
}

func (t *textInput) Update(msg tea.Msg) tea.Cmd {
	textLength := len(t.Model.Value())

	if textLength < MinWidth {
		t.Model.Width = MinWidth

	} else if textLength > MaxWidth {
		t.Model.Width = MaxWidth

	} else {
		t.Model.Width = textLength + 1
	}

	m, cmd := t.Model.Update(msg)
	t.Model = m
	return cmd
}

type HeaderEntryRemoveOnEnter struct {
	id uuid.UUID
}

func (h HeaderEntryRemoveOnEnter) Do() tea.Cmd {
	removeHeaderAction := action.NewRemoveHeader(h.id)
	return func() tea.Msg { return removeHeaderAction }
}

type Entry struct {
	Id           uuid.UUID
	removeHeader *addremovebox.AddRemoveBox
	keyInput     *textInput
	valueInput   *textInput
}

func NewEntry() Entry {
	entryId := uuid.New()

	removeHeaderBox := addremovebox.NewAddRemoveBox("-", HeaderEntryRemoveOnEnter{id: entryId})
	keyInput := newTextInput(keyPlaceholder)
	valueInput := newTextInput(valuePlaceholder)
	return Entry{
		Id:           entryId,
		removeHeader: &removeHeaderBox,
		keyInput:     &keyInput,
		valueInput:   &valueInput,
	}
}

func newTextInput(placeholder string) textInput {
	ti := textInput{textinput.New()}
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 250
	ti.Width = MinWidth
	ti.PromptStyle = stylePrompt
	ti.Prompt = ""
	ti.Cursor.Style = styleCursor
	return ti
}

func (e Entry) GetFocusables() []app.Focusable {
	return []app.Focusable{e.removeHeader, e.keyInput, e.valueInput}
}
