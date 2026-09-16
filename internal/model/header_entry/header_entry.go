package headerentry

import (
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/add_remove_box"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/google/uuid"
)

const (
	keyPlaceholder   = "Key"
	valuePlaceholder = "Value"

	MinWidth = 5
	MaxWidth = 30
	MaxSize  = 300
)

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
	ti.CharLimit = MaxSize
	ti.Width = MinWidth
	ti.Prompt = ""
	return ti
}

func (e Entry) GetFocusables() []app.Focusable {
	return []app.Focusable{e.removeHeader, e.keyInput, e.valueInput}
}
