package raycrossing

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy/location"
)

// LocatePointInRing determine where the point is with regards to the ring
func LocatePointInRing(layout geom.Layout, p geom.Coord, ring []float64) location.Type {
	_ = "STUB: not implemented"
	return *new(location.Type)
}

type rayCrossingCounter struct {
	p             geom.Coord
	crossingCount int
	// true if the test point lies on an input segment
	isPointOnSegment bool
}

// Gets the {@link Location} of the point relative to
// the ring, polygon
//  or multipolygon from which the processed segments were provided.
//
// This method only determines the correct location
// if <b>all</b> relevant segments must have been processed.
//
// return the Location of the point

func (counter *rayCrossingCounter) getLocation() location.Type {
	_ = "STUB: not implemented"
	return *new(location.Type)
}

// The point is in the interior of the ring if the number of X-crossings is
// odd.

/**
 * Counts a segment
 *
 * @param p1 an endpoint of the segment
 * @param p2 another endpoint of the segment
 */
func (counter *rayCrossingCounter) countSegment(p1, p2 geom.Coord) {
	_ = "STUB: not implemented"
	/**
	 * For each segment, check if it crosses
	 * a horizontal ray running from the test point in the positive x direction.
	 */return
}

// check if the segment is strictly to the left of the test point

// check if the point is equal to the current ring vertex

/**
 * For horizontal segments, check if the point is on the segment.
 * Otherwise, horizontal segments are not counted.
 */

/**
 * Evaluate all non-horizontal segments which cross a horizontal ray to the
 * right of the test pt. To avoid double-counting shared vertices, we use the
 * convention that
 * <ul>
 * <li>an upward edge includes its starting endpoint, and excludes its
 * final endpoint
 * <li>a downward edge excludes its starting endpoint, and includes its
 * final endpoint
 * </ul>
 */

// translate the segment so that the test point lies on the origin

/**
 * The translated segment straddles the x-axis. Compute the sign of the
 * ordinate of intersection with the x-axis. (y2 != y1, so denominator
 * will never be 0.0)
 */
// double xIntSign = RobustDeterminant.signOfDet2x2(x1, y1, x2, y2) / (y2
// - y1);
// MD - faster & more robust computation?

// xsave = xInt;

// System.out.println("xIntSign(" + x1 + ", " + y1 + ", " + x2 + ", " + y2 + " = " + xIntSign);
// The segment crosses the ray if the sign is strictly positive.
