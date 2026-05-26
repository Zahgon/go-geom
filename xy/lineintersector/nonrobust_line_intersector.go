package lineintersector

import (
	geom "github.com/twpayne/go-geom"
)

// NonRobustLineIntersector is a performant but non robust line intersection implementation.
type NonRobustLineIntersector struct{}

func (li NonRobustLineIntersector) computePointOnLineIntersection(data *lineIntersectorData, p, lineStart, lineEnd geom.Coord) {
	_ = "STUB: not implemented"
	/*
	 *  Coefficients of line eqns.
	 */return
}

/*
 *  'Sign' values
 */

/*
 *  Compute a1, b1, c1, where line joining points 1 and 2
 *  is "a1 x  +  b1 y  +  c1  =  0".
 */

/*
 *  Compute r3 and r4.
 */

// if r != 0 the point does not lie on the line

// Point lies on line - check to see whether it lies in line segment.

func (li NonRobustLineIntersector) computeLineOnLineIntersection(data *lineIntersectorData, line1Start, line1End, line2Start, line2End geom.Coord) {
	_ = "STUB: not implemented"
	/*
	 *  Coefficients of line eqns.
	 */return
}

/*
 *  Coefficients of line eqns.
 */

/*
 *  Coefficients of line eqns.
 */

/*
 *  'Sign' values
 */
// double denom, offset, num;     /* Intermediate values */

/*
 *  Compute a1, b1, c1, where line joining points 1 and 2
 *  is "a1 x  +  b1 y  +  c1  =  0".
 */

/*
 *  Compute r3 and r4.
 */

/*
 *  Check signs of r3 and r4.  If both point 3 and point 4 lie on
 *  same side of line 1, the line segments do not intersect.
 */

/*
 *  Compute a2, b2, c2
 */

/*
 *  Compute r1 and r2
 */

/*
 *  Check signs of r1 and r2.  If both point 1 and point 2 lie
 *  on same side of second line segment, the line segments do
 *  not intersect.
 */

/**
 *  Line segments intersect: compute intersection point.
 */

// check if this is a proper intersection BEFORE truncating values,
// to avoid spurious equality comparisons with endpoints

func (li NonRobustLineIntersector) computeCollinearIntersection(data *lineIntersectorData, line1Start, line1End, line2Start, line2End geom.Coord) {
	_ = "STUB: not implemented"
	return
}

// make sure p3-p4 is in same direction as p1-p2

// intersection MUST be a segment - compute endpoints
