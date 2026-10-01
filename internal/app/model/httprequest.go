package model

type HttpRequest struct {
	Url        string
	Headers    map[string]string
	Body       string
	Method     string
	ScriptPath string

	CaCertificatePath     string
	ClientCertificatePath string
	ClientKeyPath         string
}
