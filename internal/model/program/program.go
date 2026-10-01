package program

import (
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app/model"
	"github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/pigeon_carrier/internal/model/request"
	"github.com/pigeon_carrier/internal/model/results"
	"github.com/pigeon_carrier/internal/model/sending"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type ResultState string

const (
	ResultStateNone       ResultState = ""
	ResultStateHasResults ResultState = "hasResults"
	ResultStateIsSending  ResultState = "isSending"
)

type HttpClient interface {
	CallHttp(logger zerolog.Logger, httpRequest model.HttpRequest) action.UpdateResults
}

type Program struct {
	log          zerolog.Logger
	currentFocus int

	sendTab *enterbox.EnterBox
	loadTab *enterbox.EnterBox
	envTab  *enterbox.EnterBox

	request *request.Request

	sending     *sending.Sending
	results     *results.Results
	resultState ResultState

	httpClient HttpClient
}

func NewProgram(logger zerolog.Logger, httpClient HttpClient) Program {
	programLogger := logger.With().Str("module", "program").Logger()

	sendTab := enterbox.NewEnterBox(programLogger, "Send", nil)
	loadTab := enterbox.NewEnterBox(programLogger, "Load", nil)
	envTab := enterbox.NewEnterBox(programLogger, "Env", nil)

	request := request.NewRequest(programLogger)

	sending := sending.NewSending(programLogger)
	results := results.NewResults(programLogger)

	program := Program{
		log:          programLogger,
		currentFocus: 0,
		sendTab:      &sendTab,
		loadTab:      &loadTab,
		envTab:       &envTab,
		request:      request,
		sending:      sending,
		results:      results,
		resultState:  ResultStateNone,
		httpClient:   httpClient,
	}
	program.setFocus() // set initial focus to the first focusable element

	program.log.Debug().Msg("Program initialized")
	return program
}

func (m Program) Init() tea.Cmd {
	return textinput.Blink
}
