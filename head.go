package royalfetch

import (
	"net/http"
	"strings"
	"time"
)

func (r RoyalFetch) Head(url string) (*http.Response, error) {

	if !strings.Contains(url, r.BaseURL) {
		url = r.BaseURL + url
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	if r.Auth != nil {
		err := SetHttpAuth(req, r.Auth)
		if err != nil {
			return nil, err
		}
	}

	if r.Proxy != nil {
		err := SetHttpProxy(&r.httpClient, r.Proxy)
		if err != nil {
			return nil, err
		}
	}

	response, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if r.Retry > 0 && r.ContainsCode(response.StatusCode) {
		r.Retry--
		response.Body.Close()
		if r.WaitingTime > 0 {
			time.Sleep(time.Duration(r.WaitingTime) * time.Millisecond)
			if r.WaitTimeIncreaseRate > 1.0 {
				r.WaitingTime *= r.WaitTimeIncreaseRate
			}
		}
		return r.Head(url)
	}

	return response, nil
}
