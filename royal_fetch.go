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
		httpClient = defaultHttpClient()
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

func defaultHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(5000) * time.Millisecond,
	}
}

func (r RoyalFetch) Clone() *RoyalFetch {

	new := new(RoyalFetch)

	new.BaseURL = r.BaseURL
	new.Retry = r.Retry

	if r.CodesToRetry != nil {
		new.CodesToRetry = make([]int, len(r.CodesToRetry))
		copy(new.CodesToRetry, r.CodesToRetry)
	}

	new.WaitingTime = r.WaitingTime
	new.WaitTimeIncreaseRate = r.WaitTimeIncreaseRate
	if r.Headers != nil {
		new.Headers = make(map[string]string, len(r.Headers))
		for key, value := range r.Headers {
			new.Headers[key] = value
		}
	}
	new.Proxy = r.Proxy
	new.Auth = r.Auth
	new.httpClient = r.httpClient

	return new
}

func (r RoyalFetch) ContainsCode(code int) bool {
	for _, c := range r.CodesToRetry {
		if c == code {
			return true
		}
	}
	return false
}
