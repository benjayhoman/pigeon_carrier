package results

import (
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/pigeon_carrier/internal/model/scrollabletext"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type Results struct {
	log     zerolog.Logger
	Status  int
	Headers map[string]string

	hideHeaders         bool
	toggleHeadersButton *addremovebox.EnterBox

	body *scrollabletext.ScrollableText
}

func NewResults(logger zerolog.Logger) *Results {
	resultsLogger := logger.With().Str("module", "results").Logger()

	toggleHeaderButton := addremovebox.NewEnterBox(resultsLogger, "Show Headers", ToggleResultHeadersOnEnter{})
	body := scrollabletext.NewScrollableText(resultsLogger)

	return &Results{
		log:                 resultsLogger,
		Status:              0,
		Headers:             nil,
		hideHeaders:         true,
		toggleHeadersButton: &toggleHeaderButton,
		body:                &body,
	}
}

func (r Results) Init() tea.Cmd {
	return nil
}

func (r *Results) GetFocusables() []app.Focusable {
	return []app.Focusable{r.toggleHeadersButton, r.body}
}
