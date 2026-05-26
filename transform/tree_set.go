package transform

import (
	"github.com/twpayne/go-geom"
)

// Compare compares two coordinates for equality and magnitude
type Compare interface {
	IsEquals(x, y geom.Coord) bool
	IsLess(x, y geom.Coord) bool
}

type tree struct {
	left  *tree
	value geom.Coord
	right *tree
}

// TreeSet sorts the coordinates according to the Compare strategy and removes duplicates as
// dicated by the Equals function of the Compare strategy
type TreeSet struct {
	compare Compare
	tree    *tree
	size    int
	layout  geom.Layout
	stride  int
}

// NewTreeSet creates a new TreeSet instance
func NewTreeSet(layout geom.Layout, compare Compare) *TreeSet {
	_ = "STUB: not implemented"
	return nil
}

// Insert adds a new coordinate to the tree set
// the coordinate must be the same size as the Stride of the layout provided
// when constructing the TreeSet
// Returns true if the coordinate was added, false if it was already in the tree
func (set *TreeSet) Insert(coord geom.Coord) bool { _ = "STUB: not implemented"; return false }

// ToFlatArray returns an array of floats containing all the coordinates in the TreeSet
func (set *TreeSet) ToFlatArray() []float64 { _ = "STUB: not implemented"; return nil }

func (set *TreeSet) walk(t *tree, visitor func([]float64)) { _ = "STUB: not implemented"; return }

func (set *TreeSet) insertImpl(t *tree, v []float64) (*tree, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
