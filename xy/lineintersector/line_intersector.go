package lineintersector

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy/lineintersection"
)

// Strategy is the line intersection implementation
type Strategy interface {
	computePointOnLineIntersection(data *lineIntersectorData, p, lineEndpoint1, lineEndpoint2 geom.Coord)
	computeLineOnLineIntersection(data *lineIntersectorData, line1End1, line1End2, line2End1, line2End2 geom.Coord)
}

// PointIntersectsLine tests if point intersects the line
func PointIntersectsLine(strategy Strategy, point, lineStart, lineEnd geom.Coord) (hasIntersection bool) {
	_ = "STUB: not implemented"
	return false
}

// LineIntersectsLine tests if the first line (line1Start,line1End) intersects the second line (line2Start, line2End)
// and returns a data structure that indicates if there was an intersection, the type of intersection and where the intersection
// was.  See lineintersection.Result for a more detailed explanation of the result object
func LineIntersectsLine(strategy Strategy, line1Start, line1End, line2Start, line2End geom.Coord) lineintersection.Result {
	_ = "STUB: not implemented"
	return *new(lineintersection.Result)
}

// An internal data structure for containing the data during calculations
type lineIntersectorData struct {
	// new Coordinate[2][2];
	inputLines [2][2]geom.Coord

	// if only a point intersection then 0 index coord will contain the intersection point
	// if co-linear (lines overlay each other) the two coordinates represent the start and end points of the overlapping lines.
	intersectionPoints [2]geom.Coord
	intersectionType   lineintersection.Type

	// The indexes of the endpoints of the intersection lines, in order along
	// the corresponding line
	isProper bool
	pa, pb   geom.Coord
	strategy Strategy
}

/**
 *  RParameter computes the parameter for the point p
 *  in the parameterized equation
 *  of the line from p1 to p2.
 *  This is equal to the 'distance' of p along p1-p2
 */
func rParameter(p1, p2, p geom.Coord) float64 {
	_ = "STUB: not implemented"

	// compute maximum delta, for numerical stability
	// also handle case of p1-p2 being vertical or horizontal
	return 0
}
