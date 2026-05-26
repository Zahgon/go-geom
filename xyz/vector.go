package xyz

import (
	"github.com/twpayne/go-geom"
)

// VectorDot calculates the dot product of two vectors
func VectorDot(v1Start, v1End, v2Start, v2End geom.Coord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// VectorNormalize creates a coordinate that is the normalized vector from 0,0,0 to vector
func VectorNormalize(vector geom.Coord) geom.Coord {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// VectorLength calculates the length of the vector from 0,0,0 to vector
func VectorLength(vector geom.Coord) float64 { _ = "STUB: not implemented"; return 0 }
