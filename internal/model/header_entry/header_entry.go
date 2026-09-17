package headerentry

import (
	"github.com/pigeon_carrier/internal/app"
	enterbox "github.com/pigeon_carrier/internal/model/enterbox"

	"charm.land/bubbles/v2/textinput"
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
	removeHeader *enterbox.EnterBox
	keyInput     *textInput
	valueInput   *textInput
}

func NewEntry() Entry {
	entryId := uuid.New()

	removeHeaderBox := enterbox.NewEnterBox("-", HeaderEntryRemoveOnEnter{id: entryId})
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
	ti.SetWidth(MinWidth)
	ti.SetVirtualCursor(true)
	ti.Prompt = ""
	return ti
}

func (e Entry) GetFocusables() []app.Focusable {
	return []app.Focusable{e.removeHeader, e.keyInput, e.valueInput}
}

func (e Entry) GetKeyValue() (string, string) {
	return e.keyInput.Value(), e.valueInput.Value()
}
