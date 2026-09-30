package config

import (
	"github.com/pigeon_carrier/internal/model/common"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type textInput struct {
	textinput.Model
}

func newTextInput(placeholder string) textInput {
	ti := textInput{textinput.New()}
	ti.Placeholder = placeholder
	ti.CharLimit = MaxSize
	ti.SetWidth(MinWidth)
	ti.SetVirtualCursor(true)
	ti.Prompt = ""
	return ti
}

func (t *textInput) Focus() tea.Cmd {
	return t.Model.Focus()
}

func (t *textInput) Blur() tea.Cmd {
	t.Model.Blur()
	return nil
}

func (t *textInput) Update(msg tea.Msg) tea.Cmd {
	if !t.Model.Focused() {
		return nil
	}

	textLength := len(t.Model.Value())

	if textLength < MinWidth {
		t.Model.SetWidth(MinWidth)

	} else if textLength > MaxWidth {
		t.Model.SetWidth(MaxWidth)

	} else {
		t.Model.SetWidth(textLength + 1)
	}

	copyPasteCmd := common.HandleCopyAndPaste(msg, &t.Model)

	m, cmd := t.Model.Update(msg)
	t.Model = m
	return tea.Batch(copyPasteCmd, cmd)
}
