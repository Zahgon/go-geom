package sorting

import "github.com/twpayne/go-geom"

// FlatCoord is a sort.Interface implementation that will result in sorting the
// wrapped coords based on the comparator function
//
// Note: this data struct cannot be used with its 0 values.  it must be
// constructed using NewFlatCoordSorting
type FlatCoord struct {
	isLess IsLess
	coords []float64
	layout geom.Layout
	stride int
}

// IsLess the function used by FlatCoord to sort the coordinate array
// returns true is v1 is less than v2
type IsLess func(v1, v2 []float64) bool

// IsLess2D is a comparator that compares based on the size of the x and y coords.
//
// First the x coordinates are compared.
// if x coords are equal then the y coords are compared
func IsLess2D(v1, v2 []float64) bool { _ = "STUB: not implemented"; return false }

// NewFlatCoordSorting2D creates a Compare2D based sort.Interface implementation
func NewFlatCoordSorting2D(layout geom.Layout, coordData []float64) FlatCoord {
	_ = "STUB: not implemented"
	return *new(FlatCoord)
}

// NewFlatCoordSorting creates a sort.Interface implementation based on the Comparator function
func NewFlatCoordSorting(layout geom.Layout, coordData []float64, comparator IsLess) FlatCoord {
	_ = "STUB: not implemented"
	return *new(FlatCoord)
}

func (s FlatCoord) Len() int { _ = "STUB: not implemented"; return 0 }

func (s FlatCoord) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s FlatCoord) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
