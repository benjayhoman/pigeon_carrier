package method

import tea "github.com/charmbracelet/bubbletea"

type MethodType string

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	DELETE MethodType = "DELETE"
	PATCH  MethodType = "PATCH"
)

type Method struct {
	selected int
	focused  bool
}

func (m Method) Init() tea.Cmd {
	return nil
}

func NewMethod() *Method {
	return &Method{
		selected: 0,
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
