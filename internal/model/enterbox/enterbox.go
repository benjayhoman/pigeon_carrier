package enterbox

import (
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type OnEnter interface {
	Do() tea.Cmd
}

type EnterBox struct {
	log     zerolog.Logger
	symbol  string
	focused bool
	action  OnEnter
}

func (b EnterBox) Init() tea.Cmd {
	return nil
}

func NewEnterBox(logger zerolog.Logger, symbol string, action OnEnter) EnterBox {
	return EnterBox{
		log:    logger.With().Str("module", "enterbox").Logger(),
		symbol: symbol,
		action: action,
	}
}

func (b *EnterBox) SetSymbol(symbol string) {
	b.symbol = symbol
}
