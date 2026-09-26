package sending

import (
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type Sending struct {
	log       zerolog.Logger
	frame     int
	isRunning bool
}

func NewSending(logger zerolog.Logger) *Sending {
	return &Sending{
		log:       logger,
		frame:     0,
		isRunning: false,
	}
}

func (s Sending) Init() tea.Cmd {
	return nil
}
