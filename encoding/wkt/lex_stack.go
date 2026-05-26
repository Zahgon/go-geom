package wkt

import (
	"github.com/twpayne/go-geom"
)

// layoutStackObj is a stack object used in the layout parsing stack.
type layoutStackObj struct {
	// layout is the currently parsed geometry type.
	layout geom.Layout
	// inBaseTypeCollection is a bool where true means we are at the top-level or in a base type GEOMETRYCOLLECTION.
	inBaseTypeCollection bool
	// nextPointMustBeEmpty is a bool where true means the next scanned point must be EMPTY. It is used to handle
	// the edge case where a base type geometry is allowed in a GEOMETRYCOLLECTIONM but only if it is EMPTY.
	nextPointMustBeEmpty bool
}

// layoutStack is a stack used for parsing the geometry type. An initial frame is pushed for the top level context.
// After that, a frame is pushed for each (nested) geometrycollection is encountered and it is popped when we
// finish scanning that geometrycollection. The initial frame should never be popped off.
type layoutStack struct {
	data []layoutStackObj
}

// makeLayoutStack returns a newly created layoutStack. An initial frame is pushed for the top level context.
func makeLayoutStack() layoutStack { _ = "STUB: not implemented"; return *new(layoutStack) }

// push constructs a layoutStackObj for a layout and pushes it onto the layout stack.
func (s *layoutStack) push(layout geom.Layout) {
	_ = "STUB: not implemented"
	// inBaseTypeCollection inherits from outer context.
	return
}

// This should never happen.

// pop pops a layoutStackObj from the layout stack and returns its layout.
func (s *layoutStack) pop() geom.Layout { _ = "STUB: not implemented"; return *new(geom.Layout) }

// top returns a pointer to the layoutStackObj currently at the top of the stack.
func (s *layoutStack) top() *layoutStackObj { _ = "STUB: not implemented"; return nil }

// topLayout returns the layout field of the topmost layoutStackObj.
func (s *layoutStack) topLayout() geom.Layout {
	_ = "STUB: not implemented"
	return *

	// topLayout returns the inBaseTypeCollection field of the topmost layoutStackObj.
	new(geom.Layout)
}

func (s *layoutStack) topInBaseTypeCollection() bool { _ = "STUB: not implemented"; return false }

// topLayout returns the nextPointMustBeEmpty field of the topmost layoutStackObj.
func (s *layoutStack) topNextPointMustBeEmpty() bool { _ = "STUB: not implemented"; return false }

// setTopLayout sets the layout field of the topmost layoutStackObj.
func (s *layoutStack) setTopLayout(layout geom.Layout) { _ = "STUB: not implemented"; return }

// This should never happen.

// setTopNextPointMustBeEmpty sets the nextPointMustBeEmpty field of the topmost layoutStackObj.
func (s *layoutStack) setTopNextPointMustBeEmpty(nextPointMustBeEmpty bool) {
	_ = "STUB: not implemented"
	return
}

// assertNotEmpty checks that the stack is not empty and panics if it is.
func (s *layoutStack) assertNotEmpty() {
	_ = "STUB: not implemented"
	// Layout stack should never be empty.
	return
}

// assertNoGeometryCollectionFramesLeft checks that no frames corresponding to geometrycollections are left on the stack.
func (s *layoutStack) assertNoGeometryCollectionFramesLeft() {
	_ = "STUB: not implemented"
	// The initial stack frame should be the only one remaining at the end.
	return
}

// atTopLevel returns whether or not the stack has only the first frame which represents that we are currently
// not inside a geometrycollection.
func (s *layoutStack) atTopLevel() bool { _ = "STUB: not implemented"; return false }
