package results

import (
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/enterbox"

	tea "charm.land/bubbletea/v2"
)

type ResultState string

const (
	ResultStateNone       ResultState = ""
	ResultStateHasResults ResultState = "hasResults"
	ResultStateIsSending  ResultState = "isSending"
)

type Results struct {
	Body    string
	Status  int
	Headers map[string]string

	state              ResultState
	sendAnimationFrame int

	hideHeaders         bool
	toggleHeadersButton *addremovebox.EnterBox
}

func NewResults() *Results {
	toggleHeaderButton := addremovebox.NewEnterBox("Toggle Headers", ToggleResultHeadersOnEnter{})
	return &Results{
		state:               ResultStateNone,
		Body:                "",
		Status:              0,
		Headers:             nil,
		toggleHeadersButton: &toggleHeaderButton,
	}
}

func (r Results) Init() tea.Cmd {
	return nil
}

func (r *Results) GetFocusables() []app.Focusable {
	return []app.Focusable{r.toggleHeadersButton}
}
