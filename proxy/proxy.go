package proxy

import "github.com/MXLange/royalfetch/v2/auth"

type Proxy struct {
	// Protocol string //TODO
	Host string
	Port int
	Auth *auth.BasicAuth // Seting this field to nil will disable the proxy authentication
}
