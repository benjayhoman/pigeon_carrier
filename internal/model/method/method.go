package method

import (
	"github.com/pigeon_carrier/internal/model/option"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type MethodType string

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	DELETE MethodType = "DELETE"
	PATCH  MethodType = "PATCH"
)

type Method struct {
	log    zerolog.Logger
	option *option.Option
}

func (m Method) Init() tea.Cmd {
	return nil
}

func NewMethod(logger zerolog.Logger) *Method {
	return &Method{
		log:    logger.With().Str("module", "method").Logger(),
		option: option.NewOption(listTypes()),
	}
}

func listTypes() []string {
	return []string{
		string(GET),
		string(POST),
		string(PUT),
		string(DELETE),
		string(PATCH),
	}
}
