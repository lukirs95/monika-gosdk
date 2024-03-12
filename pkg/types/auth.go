package types

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Expires  uint64 `json:"expires"`
}

func NewAuth(username string, password string, expires uint64) *Auth {
	return &Auth{
		Username: username,
		Password: password,
		Expires:  expires,
	}
}
