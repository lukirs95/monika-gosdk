package types

type Auth struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
	Expires  bool     `json:"expires"`
}

func NewAuth(username string, password string, expires bool) *Auth {
	return &Auth{
		Username: Username(username),
		Password: Password(password),
		Expires:  expires,
	}
}
