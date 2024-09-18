package bitwise

// RoomStatus represents the status of a hotel room
type RoomStatus uint8

// Room status constants using bitmasks
const (
	RoomStatusVacant      RoomStatus = 1 << iota // 1
	RoomStatusClean                              // 2
	RoomStatusInspected                          // 4
	RoomStatusMaintenance                        // 8
	RoomStatusOutOfOrder                         // 16
	RoomStatusReserved                           // 32

	// Combined status
	RoomStatusVacantCleanInspected RoomStatus = RoomStatusVacant | RoomStatusClean | RoomStatusInspected
	RoomStatusVacantDirty          RoomStatus = RoomStatusVacant | (^RoomStatusClean)
)

// String returns the string representation of the RoomStatus
func (rs RoomStatus) String() string {
	switch rs {
	case RoomStatusVacant:
		return "Vacant"
	case RoomStatusClean:
		return "Clean"
	case RoomStatusInspected:
		return "Inspected"
	case RoomStatusMaintenance:
		return "Maintenance"
	case RoomStatusOutOfOrder:
		return "Out of Order"
	case RoomStatusReserved:
		return "Reserved"
	default:
		return "Unknown"
	}
}

type Room struct {
	Name   string
	Status RoomStatus
}

// HasStatus checks if the room has the given status
func (r Room) HasStatus(rs RoomStatus) bool {
	return r.Status&rs != 0
}

// SetStatus sets the given status for the room
func (r *Room) SetStatus(rs RoomStatus) {
	r.Status |= rs
}

// ClearStatus clears the given status for the room
func (r *Room) ClearStatus(rs RoomStatus) {
	r.Status &^= rs
}

// ToggleStatus toggles the given status for the room
func (r *Room) ToggleStatus(rs RoomStatus) {
	r.Status ^= rs
}
