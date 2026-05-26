package hcoords

import (
	"github.com/twpayne/go-geom"
)

// GetIntersection Computes the (approximate) intersection point between two line segments
// using homogeneous coordinates.
//
// Note that this algorithm is not numerically stable; i.e. it can produce intersection points which
// lie outside the envelope of the line segments themselves.  In order to increase the precision of the calculation
// input points should be normalized before passing them to this routine.
func GetIntersection(line1End1, line1End2, line2End1, line2End2 geom.Coord) (geom.Coord, error) {
	_ = "STUB: not implemented"
	// unrolled computation
	return *new(geom.Coord), nil
}
