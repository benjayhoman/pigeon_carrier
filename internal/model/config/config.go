package config

import (
	"github.com/pigeon_carrier/internal/app"
	"github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

const (
	MinWidth = 5
	MaxWidth = 30
	MaxSize  = 300

	showSymbol = "+"
	hideSymbol = "-"
)

type Config struct {
	log          zerolog.Logger
	toggleButton *enterbox.EnterBox
	showFields   bool

	scriptInput     *textInput
	clientCertInput *textInput
	clientKeyInput  *textInput
	caCertInput     *textInput
}

func NewConfig(logger zerolog.Logger) *Config {
	configLogger := logger.With().Str("module", "config").Logger()
	toggleButton := enterbox.NewEnterBox(configLogger, showSymbol, ToggleConfigOnEnter{})
	scriptInput := newTextInput("")
	clientCertInput := newTextInput("")
	clientKeyInput := newTextInput("")
	caCertInput := newTextInput("")

	return &Config{
		log:             configLogger,
		toggleButton:    &toggleButton,
		showFields:      false,
		scriptInput:     &scriptInput,
		clientCertInput: &clientCertInput,
		clientKeyInput:  &clientKeyInput,
		caCertInput:     &caCertInput,
	}
}

func (c Config) Init() tea.Cmd {
	return nil
}

func (c Config) GetScriptPath() string {
	return c.scriptInput.Value()
}

func (c Config) GetClientCertificatePath() string {
	return c.clientCertInput.Value()
}

func (c Config) GetClientKeyPath() string {
	return c.clientKeyInput.Value()
}

func (c Config) GetCaCertificatePath() string {
	return c.caCertInput.Value()
}

func (c Config) GetFocusables() []app.Focusable {
	if !c.showFields {
		return []app.Focusable{c.toggleButton}
	}
	return []app.Focusable{c.toggleButton, c.scriptInput, c.clientCertInput, c.clientKeyInput, c.caCertInput}
}
