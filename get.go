package royalfetch

import (
	"net/http"
	"strings"
	"time"
)

func (r *RoyalFetch) Get(url string, optional ...RoyalFetch) (*http.Response, error) {

	if len(optional) > 0 {
		return optional[0].Get(url)
	}

	if !strings.Contains(url, r.BaseURL) {
		url = r.BaseURL + url
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(r.Headers) > 0 {
		for key, value := range r.Headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	if r.Timeout > 0 {
		client.Timeout = time.Duration(r.Timeout) * time.Millisecond
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if r.Retry > 0 && r.ContainsCode(response.StatusCode) {
		r.Retry--
		if r.WaitingTime > 0 {
			time.Sleep(time.Duration(r.WaitingTime) * time.Millisecond)
			if r.WaitTimeIncreaseRate > 1.0 {
				r.WaitingTime *= r.WaitTimeIncreaseRate
			}
		}
		return r.Get(url)
	}

	return response, nil
}