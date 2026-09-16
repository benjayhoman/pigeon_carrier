package results

import tea "github.com/charmbracelet/bubbletea"

type Results struct {
	hasResults     bool
	responseBody   string
	responseStatus int
}

func NewResults() *Results {
	return &Results{
		hasResults:     false,
		responseBody:   "",
		responseStatus: 0,
	}
}

func (r Results) Init() tea.Cmd {
	return nil
}
