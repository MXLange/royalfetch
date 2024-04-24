package royalfetch

type RoyalFetch struct {
	BaseURL              string
	Retry                int
	CodesToRetry         []int
	WaitingTime          float32
	WaitTimeIncreaseRate float32
	Timeout              int
	Headers              map[string]string
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
	}
}

func (o *RoyalFetch) ContainsCode(code int) bool {
	for _, c := range o.CodesToRetry {
		if c == code {
			return true
		}
	}
	return false
}
