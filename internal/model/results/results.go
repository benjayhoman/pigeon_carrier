package results

import (
	"charm.land/bubbles/v2/viewport"
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type ResultState string

const (
	ResultStateNone       ResultState = ""
	ResultStateHasResults ResultState = "hasResults"
	ResultStateIsSending  ResultState = "isSending"
)

type Results struct {
	log     zerolog.Logger
	Status  int
	Headers map[string]string

	state              ResultState
	sendAnimationFrame int

	hideHeaders         bool
	toggleHeadersButton *addremovebox.EnterBox

	body viewport.Model
}

func NewTextArea() viewport.Model {
	ta := viewport.New(
		viewport.WithWidth(80),
		viewport.WithHeight(8),
	)
	return ta
}

func NewResults(logger zerolog.Logger) *Results {
	resultsLogger := logger.With().Str("module", "results").Logger()

	toggleHeaderButton := addremovebox.NewEnterBox(resultsLogger, "Show Headers", ToggleResultHeadersOnEnter{})
	return &Results{
		log:                 resultsLogger,
		state:               ResultStateNone,
		Status:              0,
		Headers:             nil,
		hideHeaders:         true,
		toggleHeadersButton: &toggleHeaderButton,
		body:                NewTextArea(),
	}
}

func (r Results) Init() tea.Cmd {
	return nil
}

func (r *Results) GetFocusables() []app.Focusable {
	return []app.Focusable{r.toggleHeadersButton}
}
