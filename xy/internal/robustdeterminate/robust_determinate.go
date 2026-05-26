// Package robustdeterminate implements an algorithm to compute the
// sign of a 2x2 determinant for double precision values robustly.
// It is a direct translation of code developed by Olivier Devillers.
//
// The original code carries the following copyright notice:
//
// Author : Olivier Devillers
// Olivier.Devillers@sophia.inria.fr
// http:/www.inria.fr:/prisme/personnel/devillers/anglais/determinant.html
//
// Relicensed under EDL and EPL with Permission from Olivier Devillers
// Copyright (c) 1995 by INRIA Prisme Project
// BP 93 06902 Sophia Antipolis Cedex, France.
// All rights reserved
package robustdeterminate

// Sign enumerates the different possible signs
type Sign int

const (
	// Negative indicates a negative determinate
	Negative Sign = iota - 1
	// Zero indicates the determinate is 0
	Zero
	// Positive indicates a positive determinate
	Positive
)

// SignOfDet2x2 computes the sign of the determinant of the 2x2 matrix
// with the given entries, in a robust way.
//
// return -1 if the determinant is negative,
// return  1 if the determinant is positive,
// return  0 if the determinant is 0.
func SignOfDet2x2(x1, y1, x2, y2 float64) Sign { _ = "STUB: not implemented"; return *new(Sign) }

/*
 *  testing null entries
 */

/*
 *  making y coordinates positive and permuting the entries
 */
/*
 *  so that y2 is the biggest one
 */

/*
 *  making x coordinates positive
 */
/*
 *  if |x2| < |x1| one can conclude
 */

/*
 *  all entries strictly positive   x1 <= x2 and y1 <= y2
 */

// MD - UNSAFE HACK for testing only!
//      k = (int) (x2 / x1);

/*
 *  testing if R (new U2) is in U1 rectangle
 */

/*
 *  finding R'
 */

/*
 *  exchange 1 and 2 role.
 */
// MD - UNSAFE HACK for testing only!
//      k = (int) (x1 / x2);

/*
 *  testing if R (new U1) is in U2 rectangle
 */

/*
 *  finding R'
 */
