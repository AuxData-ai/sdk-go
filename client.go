package auxdataaisdkgo

import (
	"os"
	"time"

	"github.com/AuxData-ai/utilities"
)

// NewAuxDataClientDefault creates a new instance of AuxDataClient with default settings.
// It requires an API key as a parameter and uses default values for URL, maximum retries, and timeout.
//
// Parameters:
//   - apiKey: A string representing the API key for authentication.
//
// Returns:
//   - A pointer to an instance of AuxDataClient configured with default settings.
func NewAuxDataClientDefault(apiKey string) *AuxDataClient {
	return NewAuxDataClient(apiKey, DEFAULT_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)
}

// NewAuxDataClient creates a new instance of AuxDataClient with the provided parameters.
// If the apiKey is an empty string, it attempts to retrieve the API key from the environment variable "AUXDTA_API_KEY".
// Parameters:
//   - apiKey: The API key for authentication. If empty, it will be fetched from the environment variable.
//   - url: The base URL for the API.
//   - maxRetries: The maximum number of retry attempts for failed requests.
//   - timeout: The timeout duration for the client.
//
// Returns:
//
//	A pointer to an initialized AuxDataClient.
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

func generateHttpClient(c *AuxDataClient, url string, method utilities.HTTP_METHOD) utilities.SimpleHttpClient {
	var httpClient utilities.SimpleHttpClient
	httpClient.AddBearerAuthentificationToken(c.apiKey)
	httpClient.Method = method
	httpClient.ContentType = "application/json"

	httpClient.Url = c.url + url
	return httpClient
}
