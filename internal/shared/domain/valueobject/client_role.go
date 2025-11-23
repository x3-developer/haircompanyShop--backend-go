package valueobject

import (
	"errors"
	"strings"
)

type ClientRole string

const (
	RoleUser ClientRole = "user"
)

var validClientRoles = map[ClientRole]struct{}{
	RoleUser: {},
}

func NewClientRole(value string) (ClientRole, error) {
	r := ClientRole(strings.ToLower(value))
	if _, ok := validClientRoles[r]; !ok {
		return "", errors.New("invalid role: " + value)
	}
	return r, nil
}

func (r ClientRole) String() string {
	return string(r)
}

func (r ClientRole) IsUser() bool {
	return r == RoleUser
}
