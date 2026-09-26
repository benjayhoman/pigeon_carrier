package method

import (
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type MethodType string

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	DELETE MethodType = "DELETE"
	PATCH  MethodType = "PATCH"
)

var (
	listedTypes    = listTypes()      // this should never change. Keep it precomputed
	listedTypesLen = len(listedTypes) // this should never change. Keep it precomputed
)

type Method struct {
	log      zerolog.Logger
	selected int
	focused  bool
}

func (m Method) Init() tea.Cmd {
	return nil
}

func NewMethod(logger zerolog.Logger) *Method {
	return &Method{
		log:      logger.With().Str("module", "method").Logger(),
		selected: 0,
		focused:  false,
	}
}

func listTypes() []MethodType {
	return []MethodType{
		GET,
		POST,
		PUT,
		DELETE,
		PATCH,
	}
}
