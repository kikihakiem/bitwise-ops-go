package bitwise

type (
	Role       int64
	Permission int64
)

// Role constants using bitmasks
const (
	RoleNone Role = 0
	RoleRead Role = 1 << iota // 2
	_
	RoleTriage // 8
	RoleWrite  // 16
	_
	RoleMaintain // 64
	_
	_
	RoleAdmin // 512

	PermissionOrgRead     Permission = 1 << iota // 2
	PermissionOrgWrite                           // 4
	PermissionRepoRead                           // 8
	PermissionRepoWrite                          // 16
	PermissionCI                                 // 32
	PermissionSecretRead                         // 64
	PermissionSecretWrite                        // 128
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
