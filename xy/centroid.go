package xy

import (
	"github.com/twpayne/go-geom"
)

// Centroid calculates the centroid of the geometry.  The centroid may be outside of the geometry depending
// on the topology of the geometry
func Centroid(geometry geom.T) (centroid geom.Coord, err error) {
	_ = "STUB: not implemented"
	return *new(geom.Coord), nil
}
