package domain

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleMaster Role = "manager"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleMaster:
		return true
	}
	return false
}
