package main

import (
	"fmt"
	"os"

	"github.com/pigeon_carrier/internal/model/program"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(program.NewProgram())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
