package royalfetch

import (
	"github.com/MXLange/royalfetch/auth"
	"github.com/MXLange/royalfetch/proxy"
)

type RoyalFetch struct {
	BaseURL              string
	Retry                int
	CodesToRetry         []int
	WaitingTime          float32
	WaitTimeIncreaseRate float32
	Timeout              int
	Headers              map[string]string
	Proxy                *proxy.Proxy
	Auth                 *auth.Auth
}

func New(options RoyalFetch) *RoyalFetch {
	return &RoyalFetch{
		BaseURL:              options.BaseURL,
		Retry:                options.Retry,
		CodesToRetry:         options.CodesToRetry,
		WaitingTime:          options.WaitingTime,
		WaitTimeIncreaseRate: options.WaitTimeIncreaseRate,
		Timeout:              options.Timeout,
		Headers:              options.Headers,
		Proxy:                options.Proxy,
		Auth:                 options.Auth,
	}
}
