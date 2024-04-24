package auth

type RoyalFetch struct {
	Retry   int
	Timeout int
	Headers map[string]string
	Proxy   string
	Body    string
}

func NewOptions() *RoyalFetch {
	return &RoyalFetch{
		Retry:   1,
		Timeout: 10000,
		Headers: make(map[string]string),
		Proxy:   "",
		Body:    "",
	}
}
