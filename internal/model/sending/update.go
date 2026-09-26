package sending

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

var (
	sendAnimation = []string{
		"====>",
		" ====>",
		"  ====>",
		"   ====>",
		"    ====>",
		"     ====>",
		"      ====>",
		"       ====>",
		"        ====>",
		"         ====>",
		"          ====>",
		"           ====>",
		"           <====",
		"          <====",
		"         <====",
		"        <====",
		"       <====",
		"      <====",
		"     <====",
		"    <====",
		"   <====",
		"  <====",
		" <====",
		"<====",
	}
)

type tickMsg time.Time

func (s *Sending) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case action.UpdateSending:
		// start animation
		s.frame = 0
		s.isRunning = true
		return tick()

	case action.UpdateResults:
		// stop animation
		s.isRunning = false
		return nil

	case tickMsg:
		s.log.Debug().Int("sendAnimationFrame", s.frame).Msg("Updating send animation frame")
		if !s.isRunning {
			s.frame = 0
			return nil
		}
		s.frame++
		if s.frame >= len(sendAnimation) {
			s.frame = 0
		}
		return tick()
	}

	return nil
}

func tick() tea.Cmd {
	return tea.Tick(time.Second/24, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
