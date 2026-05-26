package centralendpoint

import (
	geom "github.com/twpayne/go-geom"
)

// GetIntersection computes an approximate intersection of two line segments by taking the most central of the endpoints of the segments.
//
// This is effective in cases where the segments are nearly parallel and should intersect at an endpoint.
// It is also a reasonable strategy for cases where the endpoint of one segment lies on or almost on the interior of another one.
// Taking the most central endpoint ensures that the computed intersection point lies in the envelope of the segments.
//
// Also, by always returning one of the input points, this should result  in reducing segment fragmentation.
// Intended to be used as a last resort for  computing ill-conditioned intersection situations which cause other methods to fail.
func GetIntersection(line1End1, line1End2, line2End1, line2End2 geom.Coord) geom.Coord {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

type centralEndpointIntersector struct {
	line1End1, line1End2, line2End1, line2End2, intersectionPoint geom.Coord
}

func (intersector *centralEndpointIntersector) compute() { _ = "STUB: not implemented"; return }

func average(pts [4]geom.Coord) geom.Coord { _ = "STUB: not implemented"; return *new(geom.Coord) }

func findNearestPoint(p geom.Coord, pts [4]geom.Coord) geom.Coord {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// always initialize the result
