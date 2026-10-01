package config

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	styleDim = lipgloss.NewStyle().Faint(true)
)

func (c Config) View() string {
	result := fmt.Sprintf("%s %s", "Config:", c.toggleButton.View())
	if !c.showFields {
		return result
	}
	result += fmt.Sprintf("\n    %s %s\n", styleDim.Render("script:"), c.scriptInput.View())
	result += fmt.Sprintf("    %s %s\n", styleDim.Render("client certificate:"), c.clientCertInput.View())
	result += fmt.Sprintf("    %s %s", styleDim.Render("client key:"), c.clientKeyInput.View())
	return result
}
