package royalfetch

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/MXLange/royalfetch/v2/auth"
	"github.com/MXLange/royalfetch/v2/proxy"
)

func (o *RoyalFetch) ContainsCode(code int) bool {
	for _, c := range o.CodesToRetry {
		if c == code {
			return true
		}
	}
	return false
}

func SetHttpProxy(httpClient *http.Client, proxy *proxy.Proxy) error {

	if proxy == nil {
		return fmt.Errorf("proxy is nil")
	}

	proxyURL, err := url.Parse(proxy.Host)
	if err != nil {
		return err
	}

	if proxy.Port > 0 {
		proxyURL, err = url.Parse(fmt.Sprintf("%s:%d", proxy.Host, proxy.Port))
		if err != nil {
			return err
		}
	} else if proxy.Port < 0 {
		return fmt.Errorf("invalid proxy port: %d", proxy.Port)
	}

	if proxy.Auth != nil {
		proxyURL.User = url.UserPassword(proxy.Auth.Username, proxy.Auth.Password)
	}

	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return nil
}

func SetHttpAuth(req *http.Request, auth *auth.Auth) error {

	var err error = nil

	if auth == nil || (auth.BasicAuth == nil && auth.Token == nil) {
		return fmt.Errorf("auth is nil")
	}

	if auth.BasicAuth != nil && auth.Token != nil {
		return fmt.Errorf("auth must have only one type of authentication")
	}

	if auth.BasicAuth != nil {
		req.SetBasicAuth(auth.BasicAuth.Username, auth.BasicAuth.Password)
		return err
	}

	if auth.Token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", auth.Token.Type, auth.Token.Value))
		return err
	}

	return err
}
