package auth

type Auth struct {
	BasicAuth *BasicAuth
	Token     *Token
}

type BasicAuth struct {
	Username string
	Password string
}

type Token struct {
	Type  string
	Value string
}
