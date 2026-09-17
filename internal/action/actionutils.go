package action

import tea "charm.land/bubbletea/v2"

type MsgFunc = func() tea.Msg

func NewDefaultMsg(action any) MsgFunc {
	return func() tea.Msg {
		return action
	}
}
