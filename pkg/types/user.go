package types

import (
	"regexp"
	"strings"
)

type User struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
	Role     UserRole `json:"role"`
}

type Username string

// Sanitizes username. Removes leading and trailing whitespace and
// puts characters to lower case.
func (username *Username) Sanitize() {
	*username = Username(strings.ToLower(string(*username)))
	*username = Username(strings.TrimSpace(string(*username)))
}

// Validates username. Contstraints: 4-30 chars, chars: {a-z, A-Z, 0-9, `@`, `.`, `_`, `-`}
func (username Username) Valid() bool {
	validUser := regexp.MustCompile(`^[a-zA-Z0-9@\._-]{4,30}$`)
	return validUser.MatchString(string(username))
}

type Password string

// Validates Password. Constraints: 10-32 Characters except whitespace
func (pass Password) Valid() bool {
	validPassword := regexp.MustCompile(`^\S{10,32}$`)
	return validPassword.MatchString(string(pass))
}

type UserRole string

const (
	UserRole_ADMIN UserRole = "admin"
	UserRole_VIEW  UserRole = "view"
)

func (role UserRole) Valid() bool {
	return role == UserRole_ADMIN || role == UserRole_VIEW
}
