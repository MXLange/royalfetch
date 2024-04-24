package auth

type BasicAuth struct {
	Username string
	Password string
}

type Token struct {
	Type  string
	Value string
}

type Proxy struct {
	Host     string
	Port     int
	Username string
	Password string
}
