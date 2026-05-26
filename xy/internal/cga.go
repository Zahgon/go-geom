package internal

import (
	"github.com/twpayne/go-geom"
)

// IsPointWithinLineBounds calculates if the point p lays within the bounds of the line
// between end points lineEndpoint1 and lineEndpoint2
func IsPointWithinLineBounds(p, lineEndpoint1, lineEndpoint2 geom.Coord) bool {
	_ = "STUB: not implemented"
	return false
}

// DoLinesOverlap calculates if the bounding boxes of the two lines (line1End1, line1End2) and
// (line2End1, line2End2) overlap
func DoLinesOverlap(line1End1, line1End2, line2End1, line2End2 geom.Coord) bool {
	_ = "STUB: not implemented"
	return false
}

// Equal checks if the point starting at start one in array coords1 is equal to the
// point starting at start2 in the array coords2.
// Only x and y ordinates are compared and x is assumed to be the first ordinate and y as the second
// This is a utility method intended to be used only when performance is critical as it
// reduces readability.
func Equal(coords1 []float64, start1 int, coords2 []float64, start2 int) bool {
	_ = "STUB: not implemented"
	return false
}

// Distance2D calculates the 2d distance between the two coordinates
func Distance2D(c1, c2 geom.Coord) float64 { _ = "STUB: not implemented"; return 0 }
