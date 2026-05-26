package internal

import "github.com/twpayne/go-geom"

// CoordStack is a simple stack for pushing coordinates (in []float64 form) onto the stack and getting
// the coordinates back in the normal stack order
// Must be created with the NewCoordStack method
type CoordStack struct {
	// Data is the stack data.  the order is the most recent pushes are at the end of the slice and the oldest
	// are at the start
	Data   []float64
	stride int
}

// NewCoordStack creates a new stack with the stride indicated in the layout
func NewCoordStack(layout geom.Layout) *CoordStack { _ = "STUB: not implemented"; return nil }

// Push puts the coordinate at the location idx onto the stack.
func (stack *CoordStack) Push(data []float64, idx int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// Pop the last pushed coordinate off the stack and return the coordinate
func (stack *CoordStack) Pop() ([]float64, int) { _ = "STUB: not implemented"; return nil, 0 }

// Peek returns the most recently pushed coord without modifying the stack
func (stack *CoordStack) Peek() []float64 { _ = "STUB: not implemented"; return nil }

// Size returns the number of coordinates in the stack
func (stack *CoordStack) Size() int { _ = "STUB: not implemented"; return 0 }
