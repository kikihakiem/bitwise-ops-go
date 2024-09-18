package bitwise_test

import (
	"testing"

	"github.com/kikihakiem/bitwise"
)

func TestUserRoles(t *testing.T) {
	tests := []struct {
		name           string
		initialRoles   bitwise.Role
		assignRoles    []bitwise.Role
		checkRoles     []bitwise.Role
		expectedResult []bool
	}{
		{
			name:           "Assign and check single role",
			initialRoles:   bitwise.RoleNone,
			assignRoles:    []bitwise.Role{bitwise.RoleRead},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite},
			expectedResult: []bool{true, false},
		},
		{
			name:           "Assign and check multiple roles",
			initialRoles:   bitwise.RoleRead,
			assignRoles:    []bitwise.Role{bitwise.RoleWrite, bitwise.RoleTriage},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite, bitwise.RoleTriage, bitwise.RoleAdmin},
			expectedResult: []bool{true, true, true, false},
		},
		{
			name:           "Check roles without assignment",
			initialRoles:   bitwise.RoleMaintain,
			assignRoles:    []bitwise.Role{},
			checkRoles:     []bitwise.Role{bitwise.RoleMaintain, bitwise.RoleAdmin},
			expectedResult: []bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := bitwise.User{Name: "TestUser", Roles: tt.initialRoles}

			for _, role := range tt.assignRoles {
				user.AssignRole(role)
			}

			for i, role := range tt.checkRoles {
				if got := user.HasRole(role); got != tt.expectedResult[i] {
					t.Errorf("User.HasRole(%v) = %v, want %v", role, got, tt.expectedResult[i])
				}
			}
		})
	}
}

func TestUserRoleRevocation(t *testing.T) {
	tests := []struct {
		name           string
		initialRoles   bitwise.Role
		revokeRoles    []bitwise.Role
		checkRoles     []bitwise.Role
		expectedResult []bool
	}{
		{
			name:           "Revoke single role",
			initialRoles:   bitwise.RoleRead | bitwise.RoleWrite,
			revokeRoles:    []bitwise.Role{bitwise.RoleRead},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite},
			expectedResult: []bool{false, true},
		},
		{
			name:           "Revoke multiple roles",
			initialRoles:   bitwise.RoleRead | bitwise.RoleWrite | bitwise.RoleTriage | bitwise.RoleAdmin,
			revokeRoles:    []bitwise.Role{bitwise.RoleWrite, bitwise.RoleAdmin},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite, bitwise.RoleTriage, bitwise.RoleAdmin},
			expectedResult: []bool{true, false, true, false},
		},
		{
			name:           "Revoke non-existent role",
			initialRoles:   bitwise.RoleRead | bitwise.RoleTriage,
			revokeRoles:    []bitwise.Role{bitwise.RoleWrite},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite, bitwise.RoleTriage},
			expectedResult: []bool{true, false, true},
		},
		{
			name:           "Revoke all roles",
			initialRoles:   bitwise.RoleRead | bitwise.RoleWrite | bitwise.RoleTriage | bitwise.RoleMaintain | bitwise.RoleAdmin,
			revokeRoles:    []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite, bitwise.RoleTriage, bitwise.RoleMaintain, bitwise.RoleAdmin},
			checkRoles:     []bitwise.Role{bitwise.RoleRead, bitwise.RoleWrite, bitwise.RoleTriage, bitwise.RoleMaintain, bitwise.RoleAdmin},
			expectedResult: []bool{false, false, false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := bitwise.User{Name: "TestUser", Roles: tt.initialRoles}

			for _, role := range tt.revokeRoles {
				user.RevokeRole(role)
			}

			for i, role := range tt.checkRoles {
				if got := user.HasRole(role); got != tt.expectedResult[i] {
					t.Errorf("User.HasRole(%v) = %v, want %v", role, got, tt.expectedResult[i])
				}
			}
		})
	}
}

func TestRoleString(t *testing.T) {
	tests := []struct {
		role     bitwise.Role
		expected string
	}{
		{bitwise.RoleNone, "None"},
		{bitwise.RoleRead, "Read"},
		{bitwise.RoleTriage, "Triage"},
		{bitwise.RoleWrite, "Write"},
		{bitwise.RoleMaintain, "Maintain"},
		{bitwise.RoleAdmin, "Admin"},
		{bitwise.Role(64), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.role.String(); got != tt.expected {
				t.Errorf("Role.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
