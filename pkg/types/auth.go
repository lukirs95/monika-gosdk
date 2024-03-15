package types

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Expires  bool   `json:"expires"`
}

func NewAuth(username string, password string, expires bool) *Auth {
	return &Auth{
		Username: username,
		Password: password,
		Expires:  expires,
	}
}
