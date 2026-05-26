package xy

import (
	"github.com/twpayne/go-geom"
)

// LinesCentroid computes the centroid of all the LineStrings provided as arguments.
//
// Algorithm: Compute the average of the midpoints of all line segments weighted by the segment length.
func LinesCentroid(line *geom.LineString, extraLines ...*geom.LineString) (centroid geom.Coord) {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// LinearRingsCentroid computes the centroid of all the LinearRings provided as arguments.
//
// Algorithm: Compute the average of the midpoints of all line segments weighted by the segment length.
func LinearRingsCentroid(line *geom.LinearRing, extraLines ...*geom.LinearRing) (centroid geom.Coord) {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// MultiLineCentroid computes the centroid of the MultiLineString string
//
// Algorithm: Compute the average of the midpoints of all line segments weighted by the segment length.
func MultiLineCentroid(line *geom.MultiLineString) (centroid geom.Coord) {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// LineCentroidCalculator is the data structure that contains the centroid calculation
// data.  This type cannot be used using its 0 values, it must be created
// using NewLineCentroid
type LineCentroidCalculator struct {
	layout      geom.Layout
	stride      int
	centSum     geom.Coord
	totalLength float64
}

// NewLineCentroidCalculator creates a new instance of the calculator.
// Once a calculator is created polygons, linestrings or linear rings can be added and the
// GetCentroid method can be used at any point to get the current centroid
// the centroid will naturally change each time a geometry is added
func NewLineCentroidCalculator(layout geom.Layout) *LineCentroidCalculator {
	_ = "STUB: not implemented"
	return nil
}

// GetCentroid obtains centroid currently calculated.  Returns a 0 coord if no geometries have been added
func (calc *LineCentroidCalculator) GetCentroid() geom.Coord {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// AddPolygon adds a Polygon to the calculation.
func (calc *LineCentroidCalculator) AddPolygon(polygon *geom.Polygon) *LineCentroidCalculator {
	_ = "STUB: not implemented"
	return nil
}

// AddLine adds a LineString to the current calculation
func (calc *LineCentroidCalculator) AddLine(line *geom.LineString) *LineCentroidCalculator {
	_ = "STUB: not implemented"
	return nil
}

// AddLinearRing adds a LinearRing to the current calculation
func (calc *LineCentroidCalculator) AddLinearRing(line *geom.LinearRing) *LineCentroidCalculator {
	_ = "STUB: not implemented"
	return nil
}

func (calc *LineCentroidCalculator) addLine(line []float64, startLine, endLine int) {
	_ = "STUB: not implemented"
	return
}
