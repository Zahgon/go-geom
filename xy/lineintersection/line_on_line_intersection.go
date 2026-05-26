package lineintersection

import "github.com/twpayne/go-geom"

// Type enumerates the types of intersection two lines can have
type Type int

const (
	// NoIntersection indicates that the lines do not intersect
	NoIntersection Type = iota
	// PointIntersection indicates that the lines intersect at a point
	PointIntersection
	// CollinearIntersection indicates that the lines overlap each other
	CollinearIntersection
)

var labels = [3]string{"NoIntersection", "PointIntersection", "CollinearIntersection"}

func (t Type) String() string {
	_ = "STUB: not implemented"

	// Result the results from LineIntersectsLine function.
	// It contains the intersection point(s) and indicates what type of
	// intersection there was (or if there was no intersection)
	return ""
}

type Result struct {
	intersectionType Type
	intersection     []geom.Coord
}

// NewResult create a new result object
func NewResult(intersectionType Type, intersection []geom.Coord) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

// HasIntersection returns true if the lines have an intersection
func (i *Result) HasIntersection() bool { _ = "STUB: not implemented"; return false }

// Type returns the type of intersection between the two lines
func (i *Result) Type() Type {
	_ = "STUB: not implemented"
	return *

	// Intersection returns an array of Coords which are the intersection points.
	// If the type is PointIntersection then there will only be a single Coordinate (the first coord).
	// If the type is CollinearIntersection then there will two Coordinates the start and end points of the line
	// that represents the intersection
	new(Type)
}

func (i *Result) Intersection() []geom.Coord { _ = "STUB: not implemented"; return nil }
