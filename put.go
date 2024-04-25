package royalfetch

import (
	"net/http"
	"strings"
	"time"
)

func (r *RoyalFetch) Put(url string, body string, optional ...RoyalFetch) (*http.Response, error) {

	if len(optional) > 0 {
		return optional[0].Put(url, body)
	}

	if !strings.Contains(url, r.BaseURL) {
		url = r.BaseURL + url
	}

	req, err := http.NewRequest("PUT", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	if len(r.Headers) > 0 {
		for key, value := range r.Headers {
			req.Header.Set(key, value)
		}
	}

	if r.Auth != nil {
		err := SetHttpAuth(req, r.Auth)
		if err != nil {
			return nil, err
		}
	}

	client := &http.Client{}
	if r.Timeout > 0 {
		client.Timeout = time.Duration(r.Timeout) * time.Millisecond
	}

	if r.Proxy != nil {
		err := SetHttpProxy(client, r.Proxy)
		if err != nil {
			return nil, err
		}
	}

	response, err := client.Do(req)
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
		return r.Put(url, body)
	}

	return response, nil

}
