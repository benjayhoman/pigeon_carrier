package action

// SendRequest represents an action to send a request to a specific URL.
type UpdateResults struct {
	Body       string
	StatusCode int
	Headers    map[string]string
}

func NewUpdateResults(body string, statusCode int, headers map[string]string) UpdateResults {
	return UpdateResults{
		Body:       body,
		StatusCode: statusCode,
		Headers:    headers,
	}
}
