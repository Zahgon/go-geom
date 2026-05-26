// Package xyz contains operations in 3d coordinate space.  Each
// layout must have 3 ordinates (and thus each coordinate)
// it is assumed that x,y,z are ordinates with indexes 0,1,2 respectively
package xyz

import (
	geom "github.com/twpayne/go-geom"
)

// Distance calculates the distance between the two coordinates in 3d space.
func Distance(point1, point2 geom.Coord) float64 {
	_ = "STUB: not implemented"
	// default to 2D distance if either Z is not set
	return 0
}

// DistancePointToLine calculates the distance from point to a point on the line
func DistancePointToLine(point, lineStart, lineEnd geom.Coord) float64 {
	_ = "STUB: not implemented"
	// if start = end, then just compute distance to one of the endpoints
	return 0
}

// otherwise use comp.graphics.algorithms Frequently Asked Questions method
/*
 * (1) r = AC dot AB
 *         ---------
 *         ||AB||^2
 *
 * r has the following meaning:
 *   r=0 P = A
 *   r=1 P = B
 *   r<0 P is on the backward extension of AB
 *   r>1 P is on the forward extension of AB
 *   0<r<1 P is interior to AB
 */

// compute closest point q on line segment

// result is distance from p to q

// Equals determines if the two coordinates have equal in 3d space
func Equals(point1, other geom.Coord) bool { _ = "STUB: not implemented"; return false }

// DistanceLineToLine computes the distance between two 3D segments
func DistanceLineToLine(line1Start, line1End, line2Start, line2End geom.Coord) float64 {
	_ = "STUB: not implemented"
	/**
	 * This calculation is susceptible to roundoff errors when
	 * passed large ordinate values.
	 * It may be possible to improve this by using {@link DD} arithmetic.
	 */return 0
}

/**
 * Algorithm derived from http://softsurfer.com/Archive/algorithm_0106/algorithm_0106.htm
 */

/**
 * The lines are parallel.
 * In this case solve for the parameters s and t by assuming s is 0.
 */

// choose largest denominator for optimal numeric conditioning

/**
 * The closest points are in interiors of segments,
 * so compute them directly
 */

// length (p1-p2)
