package types

import (
	"regexp"
	"strings"
)

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type Username string

func (username *Username) Sanitize() {
	*username = Username(strings.ToLower(string(*username)))
	*username = Username(strings.TrimSpace(string(*username)))
}

func (username Username) Validate() bool {
	validUser := regexp.MustCompile(`^[a-zA-Z0-9@\._-]{4,30}$`)
	return validUser.MatchString(string(username))
}

type UserRole string

const (
	UserRole_ADMIN UserRole = "admin"
	UserRole_VIEW  UserRole = "view"
)

func (role UserRole) Valid() bool {
	return role == UserRole_ADMIN || role == UserRole_VIEW
}
