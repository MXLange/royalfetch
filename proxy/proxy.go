package proxy

import "github.com/MXLange/royalfetch/auth"

type Proxy struct {
	// Protocol string //TODO
	Host string
	Port int
	Auth *auth.BasicAuth // Seting this field to nil will disable the proxy authentication
}
