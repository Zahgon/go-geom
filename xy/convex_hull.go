package xy

import (
	"github.com/twpayne/go-geom"
)

type convexHullCalculator struct {
	layout   geom.Layout
	stride   int
	inputPts []float64
}

// ConvexHull computes the convex hull of the geometry.
// A convex hull is the smallest convex geometry that contains
// all the points in the input geometry
// Uses the Graham Scan algorithm
func ConvexHull(geometry geom.T) geom.T {
	_ = "STUB: not implemented"
	// copy coords because the algorithm reorders them
	return *new(geom.T)
}

// ConvexHullFlat computes the convex hull of the geometry.
// A convex hull is the smallest convex geometry that contains
// all the points in the input coordinates
// Uses the Graham Scan algorithm
func ConvexHullFlat(layout geom.Layout, coords []float64) geom.T {
	_ = "STUB: not implemented"
	return *new(geom.T)
}

func (calc *convexHullCalculator) getConvexHull() geom.T {
	_ = "STUB: not implemented"
	return *new(geom.T)
}

// use heuristic to reduce points, if large

// sort points for Graham scan.

// Use Graham scan to find convex hull.

// Convert array to appropriate output geometry.

func (calc *convexHullCalculator) lineOrPolygon(coordinates []float64) geom.T {
	_ = "STUB: not implemented"
	return *new(geom.T)
}

func (calc *convexHullCalculator) cleanRing(original []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (calc *convexHullCalculator) isBetween(c1, c2, c3 []float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (calc *convexHullCalculator) grahamScan(coordData []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// check for empty stack to guard against robustness problems

func (calc *convexHullCalculator) preSort(pts []float64) {
	_ = "STUB: not implemented"
	// find the lowest point in the set. If two or more points have
	// the same minimum y coordinate choose the one with the minimu x.
	// This focal point is put in array location pts[0].
	return
}

// sort the points radially around the focal point.

// Uses a heuristic to reduce the number of points scanned
// to compute the hull.
// The heuristic is to find a polygon guaranteed to
// be in (or on) the hull, and eliminate all points inside it.
// A quadrilateral defined by the extremal points
// in the four orthogonal directions
// can be used, but even more inclusive is
// to use an octilateral defined by the points in the 8 cardinal directions.
//
// Note that even if the method used to determine the polygon vertices
// is not 100% robust, this does not affect the robustness of the convex hull.
//
// To satisfy the requirements of the Graham Scan algorithm,
// the returned array has at least 3 entries.
func (calc *convexHullCalculator) reduce(inputPts []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// add points defining polygon

/**
 * Add all unique points not in the interior poly.
 * CGAlgorithms.isPointInRing is not defined for points actually on the ring,
 * but this doesn't matter since the points of the interior polygon
 * are forced to be in the reduced set.
 */

// ensure that computed array has at least 3 points (not necessarily unique)

func (calc *convexHullCalculator) padArray3(pts []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (calc *convexHullCalculator) computeOctRing(inputPts []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// Dedup adjacent points, only keep ones that are different from previous.

// Need at least 3 unique points (a triangle) to exclude anything inside.

func (calc *convexHullCalculator) computeOctPts(inputPts []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type comparator struct{}

func (c comparator) IsEquals(x, y geom.Coord) bool { _ = "STUB: not implemented"; return false }

func (c comparator) IsLess(x, y geom.Coord) bool { _ = "STUB: not implemented"; return false }
