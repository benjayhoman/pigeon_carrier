package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app/model"

	"github.com/rs/zerolog"
)

type HttpClient struct{}

func NewHttpClient() HttpClient {
	return HttpClient{}
}

func (c HttpClient) CallHttp(logger zerolog.Logger, httpRequest model.HttpRequest) action.UpdateResults {
	transport, err := createTransport(logger, httpRequest.ClientCertificatePath, httpRequest.ClientKeyPath, httpRequest.CaCertificatePath)
	if err != nil {
		return action.UpdateResults{
			Body:       fmt.Sprintf("Error creating transport: %v\n", err),
			StatusCode: 0,
			Headers:    nil,
		}
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { // do not follow redirects
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(httpRequest.Method, httpRequest.Url, strings.NewReader(httpRequest.Body))
	if err != nil {
		return action.UpdateResults{
			Body:       fmt.Sprintf("Error Building Request: %v\n", err),
			StatusCode: 0,
			Headers:    nil,
		}
	}

	req.Header.Set("User-Agent", "PigeonCarrier/1.0")
	for key, value := range httpRequest.Headers {
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

func createTransport(logger zerolog.Logger, certPath, keyPath, caCertPath string) (*http.Transport, error) {
	certificate := make([]tls.Certificate, 0)

	// 1. Load the client's certificate and private key
	if strings.TrimSpace(certPath) != "" && strings.TrimSpace(keyPath) != "" {
		clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to load client certificate/key")
			return nil, err

		} else {
			certificate = append(certificate, clientCert)
		}
	}

	// Create a CA certificate pool
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		// Fallback to empty pool if system certs can't be loaded
		caCertPool = x509.NewCertPool()
	}

	// Load the CA certificate if a path is provided
	if strings.TrimSpace(caCertPath) != "" {
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to load CA certificate")
			return nil, err

		} else {
			caCertPool.AppendCertsFromPEM(caCert)
		}
	}

	tlsConfig := &tls.Config{
		Certificates: certificate,
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12, // Best practice: enforce secure TLS versions
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	return transport, nil
}
