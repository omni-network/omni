package solvernet

var transitions = map[OrderStatus][]OrderStatus{
	StatusPending: {StatusRejected, StatusClosed, StatusFilled},
	StatusFilled:  {StatusClaimed},
}

// ValidTarget returns true if the target status can be reached from the current status.
func (s OrderStatus) ValidTarget(target OrderStatus) bool {
	current := s
	if current == target {
		return true
	}

	for _, next := range transitions[current] {
		if next.ValidTarget(target) {
			return true
		}
	}

	return false
}
