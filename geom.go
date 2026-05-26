// Package geom implements efficient geometry types for geospatial
// applications.
package geom

//go:generate bin/goderive . > derived.gen.go
//go:generate bin/gofumpt -w derived.gen.go

import (
	"errors"
)

// A Layout describes the meaning of an N-dimensional coordinate. Layout(N) for
// N > 4 is a valid layout, in which case the first dimensions are interpreted
// to be X, Y, Z, and M and extra dimensions have no special meaning.  M values
// are considered part of a linear referencing system (e.g. classical time or
// distance along a path). 1-dimensional layouts are not supported.
type Layout int

const (
	// NoLayout is an unknown layout.
	NoLayout Layout = iota
	// XY is a 2D layout (X and Y).
	XY
	// XYZ is 3D layout (X, Y, and Z).
	XYZ
	// XYM is a 2D layout with an M value.
	XYM
	// XYZM is a 3D layout with an M value.
	XYZM
)

// An ErrLayoutMismatch is returned when geometries with different layouts
// cannot be combined.
type ErrLayoutMismatch struct {
	Got  Layout
	Want Layout
}

func (e ErrLayoutMismatch) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrStrideMismatch is returned when the stride does not match the expected
// stride.
type ErrStrideMismatch struct {
	Got  int
	Want int
}

func (e ErrStrideMismatch) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrUnsupportedLayout is returned when the requested layout is not
// supported.
type ErrUnsupportedLayout Layout

func (e ErrUnsupportedLayout) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrUnsupportedType is returned when the requested type is not supported.
type ErrUnsupportedType struct {
	Value any
}

func (e ErrUnsupportedType) Error() string { _ = "STUB: not implemented"; return "" }

// A Coord represents an N-dimensional coordinate.
type Coord []float64

// Clone returns a deep copy of c.
func (c Coord) Clone() Coord { _ = "STUB: not implemented"; return *new(Coord) }

// X returns the x coordinate of c. X is assumed to be the first ordinate.
func (c Coord) X() float64 {
	_ = "STUB: not implemented"

	// Y returns the y coordinate of c. Y is assumed to be the second ordinate.
	return 0
}

func (c Coord) Y() float64 {
	_ = "STUB: not implemented"

	// Set copies the ordinate data from the other coord to this coord.
	return 0
}

func (c Coord) Set(other Coord) {
	_ = "STUB: not implemented"

	// Equal compares that all ordinates are the same in this and the other coords.
	// It is assumed that this coord and other coord both have the same (provided)
	// layout.
	return
}

func (c Coord) Equal(layout Layout, other Coord) bool { _ = "STUB: not implemented"; return false }

// T is a generic interface implemented by all geometry types.
type T interface {
	Layout() Layout
	Stride() int
	Bounds() *Bounds
	FlatCoords() []float64
	Ends() []int
	Endss() [][]int
	SRID() int
	Empty() bool
}

// MIndex returns the index of the M dimension, or -1 if the l does not have an
// M dimension.
func (l Layout) MIndex() int { _ = "STUB: not implemented"; return 0 }

// Stride returns l's number of dimensions.
func (l Layout) Stride() int { _ = "STUB: not implemented"; return 0 }

// String returns a human-readable string representing l.
func (l Layout) String() string { _ = "STUB: not implemented"; return "" }

// ZIndex returns the index of l's Z dimension, or -1 if l does not have a Z
// dimension.
func (l Layout) ZIndex() int { _ = "STUB: not implemented"; return 0 }

// SetSRID sets the SRID of an arbitrary geometry.
func SetSRID(g T, srid int) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// TransformInPlace replaces all coordinates in g using f.
func TransformInPlace(g T, f func(Coord)) T { _ = "STUB: not implemented"; return *new(T) }

// Must panics if err is not nil, otherwise it returns g.
func Must(g T, err error) T { _ = "STUB: not implemented"; return *new(T) }

var (
	errIncorrectEnd         = errors.New("geom: incorrect end")
	errLengthStrideMismatch = errors.New("geom: length/stride mismatch")
	errMisalignedEnd        = errors.New("geom: misaligned end")
	errNonEmptyEnds         = errors.New("geom: non-empty ends")
	errNonEmptyEndss        = errors.New("geom: non-empty endss")
	errNonEmptyFlatCoords   = errors.New("geom: non-empty flatCoords")
	errOutOfOrderEnd        = errors.New("geom: out-of-order end")
	errStrideLayoutMismatch = errors.New("geom: stride/layout mismatch")
)
