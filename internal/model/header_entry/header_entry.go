package headerentry

import (
	"github.com/pigeon_carrier/internal/app"
	enterbox "github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/rs/zerolog"

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
	log          zerolog.Logger
	Id           uuid.UUID
	removeHeader *enterbox.EnterBox
	keyInput     *textInput
	valueInput   *textInput
}

func NewEntry(logger zerolog.Logger) Entry {
	entryId := uuid.New()
	entryLogger := logger.With().Str("module", "header_entry").Logger()

	removeHeaderBox := enterbox.NewEnterBox(entryLogger, "-", HeaderEntryRemoveOnEnter{id: entryId})
	keyInput := newTextInput(keyPlaceholder)
	valueInput := newTextInput(valuePlaceholder)

	return Entry{
		log:          entryLogger,
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
