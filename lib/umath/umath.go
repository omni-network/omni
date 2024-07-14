// Package umath provides some useful unsigned math functions to prevent underflows.
package umath

// Subtract returns a - b and true if a >= b, otherwise 0 and false.
func Subtract(a, b uint64) (uint64, bool) {
	if a < b {
		return 0, false
	}

	return a - b, true
}

// SubtractOrZero returns a - b if a >= b, otherwise 0.
// This is a convenience function for inline usage.
func SubtractOrZero(a, b uint64) uint64 {
	resp, _ := Subtract(a, b)
	return resp
}
