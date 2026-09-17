package common

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/atotto/clipboard"
)

func HandleCopyAndPaste(msg tea.Msg, textInput *textinput.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			_ = clipboard.WriteAll(textInput.Value()) // hack to get around bubble tea not copying correctly
			return tea.SetClipboard(textInput.Value())
		}

	case tea.PasteMsg:
		textInput.SetValue(textInput.Value() + msg.Content)
		return nil

	case tea.PasteEndMsg:
		textInput.CursorEnd()
		return nil
	}

	return nil
}
