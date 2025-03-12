package auxdataaisdkgo

import (
	"os"
	"time"
)

type AuxDataClient struct {
	apiKey     string
	url        string
	maxRetries int
	timeout    time.Duration
}

func NewAuxDataClientDefault(apiKey string) *AuxDataClient {
	return NewAuxDataClient(apiKey, DEFAULT_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)
}

func NewAuxDataClient(apiKey string, url string, maxRetries int, timeout time.Duration) *AuxDataClient {
	if apiKey == "" {
		apiKey = os.Getenv("AUXDTA_API_KEY")
	}

	return &AuxDataClient{
		apiKey:     apiKey,
		url:        url,
		maxRetries: maxRetries,
		timeout:    timeout,
	}
}
