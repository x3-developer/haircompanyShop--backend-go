package valueobject

import (
	"errors"
	"strings"
)

type DashboardRole string

const (
	RoleAdmin    DashboardRole = "admin"
	RoleManager  DashboardRole = "manager"
	RoleOperator DashboardRole = "operator"
)

var validDashboardRoles = map[DashboardRole]struct{}{
	RoleAdmin:    {},
	RoleManager:  {},
	RoleOperator: {},
}

func NewDashboardRole(value string) (DashboardRole, error) {
	r := DashboardRole(strings.ToLower(value))
	if _, ok := validDashboardRoles[r]; !ok {
		return "", errors.New("invalid role: " + value)
	}
	return r, nil
}

func (r DashboardRole) String() string {
	return string(r)
}

func (r DashboardRole) IsAdmin() bool {
	return r == RoleAdmin
}

func (r DashboardRole) IsManager() bool {
	return r == RoleManager
}

func (r DashboardRole) IsOperator() bool {
	return r == RoleOperator
}
