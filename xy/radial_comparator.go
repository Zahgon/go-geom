package xy

import (
	"sort"

	"github.com/twpayne/go-geom"
)

// NewRadialSorting creates an implementation sort.Interface which will sort the wrapped coordinate array
// radially around the focal point.  The comparison is based on the angle and distance
// from the focal point.
// First the angle is checked.
// Counter clockwise indicates a greater value and clockwise indicates a lesser value
// If co-linear then the coordinate nearer to the focalPoint is considered less.
func NewRadialSorting(layout geom.Layout, coordData []float64, focalPoint geom.Coord) sort.Interface {
	_ = "STUB: not implemented"
	return *new(sort.Interface)
}

// points are collinear - check distance
