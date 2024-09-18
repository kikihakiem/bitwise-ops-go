package bitwise

type Role uint8

// Role constants using bitmasks
const (
	RoleNone     Role = 0
	RoleRead     Role = 1 << iota // 2
	RoleTriage                    // 4
	RoleWrite                     // 8
	RoleMaintain                  // 16
	RoleAdmin                     // 32
)

// String returns the string representation of the Role
func (r Role) String() string {
	switch r {
	case RoleNone:
		return "None"
	case RoleRead:
		return "Read"
	case RoleTriage:
		return "Triage"
	case RoleWrite:
		return "Write"
	case RoleMaintain:
		return "Maintain"
	case RoleAdmin:
		return "Admin"
	default:
		return "Unknown"
	}
}

type User struct {
	Name  string
	Roles Role
}

func (u User) HasRole(r Role) bool {
	return u.Roles&r != 0
}

// AssignRole assigns the given role to the user
func (u *User) AssignRole(r Role) {
	u.Roles |= r
}

// RevokeRole revokes the given role from the user
func (u *User) RevokeRole(r Role) {
	u.Roles &^= r
}
