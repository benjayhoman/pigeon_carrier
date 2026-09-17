package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/pigeon_carrier/internal/action"
)

func CallHttp(url string, headers map[string]string) action.UpdateResults {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return action.UpdateResults{
			Body:       fmt.Sprintf("Error Building Request: %v\n", err),
			StatusCode: 0,
			Headers:    nil,
		}
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return action.UpdateResults{
				Body:       "timeout",
				StatusCode: 0,
				Headers:    nil,
			}

		} else {
			return action.UpdateResults{
				Body:       fmt.Sprintf("Network Error: %v\n", err),
				StatusCode: 0,
				Headers:    nil,
			}
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return action.UpdateResults{
			Body:       fmt.Sprintf("Error reading body: %v\n", err),
			StatusCode: 0,
			Headers:    nil,
		}
	}

	responseHeaders := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			responseHeaders[key] = values[0]
		}
	}

	return action.UpdateResults{
		Body:       string(body),
		StatusCode: resp.StatusCode,
		Headers:    responseHeaders,
	}
}
