package royalfetch

import (
	"net/http"
	"time"

	"github.com/MXLange/royalfetch/v2/auth"
	"github.com/MXLange/royalfetch/v2/proxy"
)

// RoyalFetch is the main struct for the package
type RoyalFetch struct {

	// BaseURL is the base URL to use for the requests
	BaseURL string

	// Retry is the number of retries to do if the request fails, based CodesToRetry
	Retry int

	// CodesToRetry is the list of status codes to retry if the request fails
	CodesToRetry []int

	// WaitingTime is the time to wait between retries
	WaitingTime float32

	// WaitTimeIncreaseRate is the rate to increase the waiting time between retries
	WaitTimeIncreaseRate float32

	// Timeout is the timeout for the request
	TimeoutMS int

	// Headers is the list of headers to add to the request
	Headers map[string]string

	// Proxy is the proxy to use for the request
	Proxy *proxy.Proxy

	// Auth is the authentication to use for the request
	Auth *auth.Auth

	// httpClient is the http.Client to use for the requests
	httpClient http.Client
}

// New creates a new RoyalFetch instance
// options: RoyalFetch struct with the configuration
// httpClient: http.Client to use, if nil a default one will be created
func New(options RoyalFetch, httpClient *http.Client) *RoyalFetch {

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Duration(options.TimeoutMS) * time.Millisecond,
		}
	} else {
		if options.TimeoutMS > 0 {
			httpClient.Timeout = time.Duration(options.TimeoutMS) * time.Millisecond
		}
	}

	return &RoyalFetch{
		BaseURL:              options.BaseURL,
		Retry:                options.Retry,
		CodesToRetry:         options.CodesToRetry,
		WaitingTime:          options.WaitingTime,
		WaitTimeIncreaseRate: options.WaitTimeIncreaseRate,
		Headers:              options.Headers,
		Proxy:                options.Proxy,
		Auth:                 options.Auth,
		httpClient:           *httpClient,
	}
}