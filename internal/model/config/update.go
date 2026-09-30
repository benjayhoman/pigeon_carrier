package config

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

func (c *Config) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case action.ToggleConfig:
		c.showFields = !c.showFields
		if c.showFields {
			c.toggleButton.SetSymbol(hideSymbol)
		} else {
			c.toggleButton.SetSymbol(showSymbol)
		}
	}

	toggleCmd := c.toggleButton.Update(msg)
	scriptCmd := c.scriptInput.Update(msg)
	clientCertCmd := c.clientCertInput.Update(msg)
	clientKeyCmd := c.clientKeyInput.Update(msg)
	return tea.Batch(toggleCmd, scriptCmd, clientCertCmd, clientKeyCmd)
}
