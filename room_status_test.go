package bitwise_test

import (
	"testing"

	"github.com/kikihakiem/bitwise"
)

func TestRoomStatus(t *testing.T) {
	tests := []struct {
		name           string
		initialStatus  bitwise.RoomStatus
		setStatus      []bitwise.RoomStatus
		clearStatus    []bitwise.RoomStatus
		toggleStatus   []bitwise.RoomStatus
		checkStatus    []bitwise.RoomStatus
		expectedResult []bool
	}{
		{
			name:           "Set and check single status",
			initialStatus:  bitwise.RoomStatusVacant,
			setStatus:      []bitwise.RoomStatus{bitwise.RoomStatusClean},
			checkStatus:    []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean, bitwise.RoomStatusInspected},
			expectedResult: []bool{true, true, false},
		},
		{
			name:           "Set and clear multiple statuses",
			initialStatus:  bitwise.RoomStatusVacant | bitwise.RoomStatusClean,
			setStatus:      []bitwise.RoomStatus{bitwise.RoomStatusInspected},
			clearStatus:    []bitwise.RoomStatus{bitwise.RoomStatusClean},
			checkStatus:    []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean, bitwise.RoomStatusInspected},
			expectedResult: []bool{true, false, true},
		},
		{
			name:           "Toggle statuses",
			initialStatus:  bitwise.RoomStatusVacant,
			toggleStatus:   []bitwise.RoomStatus{bitwise.RoomStatusClean, bitwise.RoomStatusVacant},
			checkStatus:    []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean},
			expectedResult: []bool{false, true},
		},
		{
			name:           "Check combined status",
			initialStatus:  bitwise.RoomStatusVacantCleanInspected,
			checkStatus:    []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean, bitwise.RoomStatusInspected, bitwise.RoomStatusMaintenance},
			expectedResult: []bool{true, true, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			room := bitwise.Room{Name: "TestRoom", Status: tt.initialStatus}

			for _, status := range tt.setStatus {
				room.SetStatus(status)
			}

			for _, status := range tt.clearStatus {
				room.ClearStatus(status)
			}

			for _, status := range tt.toggleStatus {
				room.ToggleStatus(status)
			}

			for i, status := range tt.checkStatus {
				if got := room.HasStatus(status); got != tt.expectedResult[i] {
					t.Errorf("Room.HasStatus(%v) = %v, want %v", status, got, tt.expectedResult[i])
				}
			}
		})
	}
}

func TestRoomStatusString(t *testing.T) {
	tests := []struct {
		status bitwise.RoomStatus
		want   string
	}{
		{bitwise.RoomStatusVacant, "Vacant"},
		{bitwise.RoomStatusClean, "Clean"},
		{bitwise.RoomStatusInspected, "Inspected"},
		{bitwise.RoomStatusMaintenance, "Maintenance"},
		{bitwise.RoomStatusOutOfOrder, "Out of Order"},
		{bitwise.RoomStatusReserved, "Reserved"},
		{bitwise.RoomStatus(255), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("RoomStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinedRoomStatuses(t *testing.T) {
	tests := []struct {
		name           string
		status         bitwise.RoomStatus
		checkStatuses  []bitwise.RoomStatus
		expectedResult []bool
	}{
		{
			name:           "Vacant, Clean, and Inspected",
			status:         bitwise.RoomStatusVacantCleanInspected,
			checkStatuses:  []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean, bitwise.RoomStatusInspected},
			expectedResult: []bool{true, true, true},
		},
		{
			name:           "Vacant and Dirty",
			status:         bitwise.RoomStatusVacantDirty,
			checkStatuses:  []bitwise.RoomStatus{bitwise.RoomStatusVacant, bitwise.RoomStatusClean},
			expectedResult: []bool{true, false},
		},
		{
			name:           "Reversed VCI",
			status:         bitwise.RoomStatusInspected | bitwise.RoomStatusClean | bitwise.RoomStatusVacant | bitwise.RoomStatusReserved,
			checkStatuses:  []bitwise.RoomStatus{bitwise.RoomStatusVacantCleanInspected},
			expectedResult: []bool{true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			room := bitwise.Room{Name: "TestRoom", Status: tt.status}

			for i, status := range tt.checkStatuses {
				if got := room.HasStatus(status); got != tt.expectedResult[i] {
					t.Errorf("Room.HasStatus(%v) = %v, want %v", status, got, tt.expectedResult[i])
				}
			}
		})
	}
}
