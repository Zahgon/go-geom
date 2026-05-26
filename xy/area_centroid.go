package xy

import (
	"github.com/twpayne/go-geom"
)

// PolygonsCentroid computes the centroid of an area geometry. (Polygon)
//
// Algorithm
// Based on the usual algorithm for calculating the centroid as a weighted sum of the centroids
// of a decomposition of the area into (possibly overlapping) triangles.
//
// The algorithm has been extended to handle holes and multi-polygons.
//
// See http://www.faqs.org/faqs/graphics/algorithms-faq/ for further details of the basic approach.
//
// The code has also be extended to handle degenerate (zero-area) polygons.
//
// In this case, the centroid of the line segments in the polygon will be returned.
func PolygonsCentroid(polygon *geom.Polygon, extraPolys ...*geom.Polygon) (centroid geom.Coord) {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// MultiPolygonCentroid computes the centroid of an area geometry. (MultiPolygon)
//
// Algorithm
// Based on the usual algorithm for calculating the centroid as a weighted sum of the centroids
// of a decomposition of the area into (possibly overlapping) triangles.
//
// The algorithm has been extended to handle holes and multi-polygons.
//
// See http://www.faqs.org/faqs/graphics/algorithms-faq/ for further details of the basic approach.
//
// The code has also be extended to handle degenerate (zero-area) polygons.
//
// In this case, the centroid of the line segments in the polygon will be returned.
func MultiPolygonCentroid(polygon *geom.MultiPolygon) (centroid geom.Coord) {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// AreaCentroidCalculator is the data structure that contains the centroid calculation
// data.  This type cannot be used using its 0 values, it must be created
// using NewAreaCentroid
type AreaCentroidCalculator struct {
	layout        geom.Layout
	stride        int
	basePt        geom.Coord
	triangleCent3 geom.Coord // temporary variable to hold centroid of triangle
	areasum2      float64    // Partial area sum
	cg3           geom.Coord // partial centroid sum

	centSum     geom.Coord // data for linear centroid computation, if needed
	totalLength float64
}

// NewAreaCentroidCalculator creates a new instance of the calculator.
// Once a calculator is created polygons can be added to it and the
// GetCentroid method can be used at any point to get the current centroid
// the centroid will naturally change each time a polygon is added
func NewAreaCentroidCalculator(layout geom.Layout) *AreaCentroidCalculator {
	_ = "STUB: not implemented"
	return nil
}

// GetCentroid obtains centroid currently calculated.  Returns a 0 coord if no geometries have been added
func (calc *AreaCentroidCalculator) GetCentroid() geom.Coord {
	_ = "STUB: not implemented"
	return *new(geom.Coord)
}

// if polygon was degenerate, compute linear centroid instead

// AddPolygon adds a polygon to the calculation.
func (calc *AreaCentroidCalculator) AddPolygon(polygon *geom.Polygon) {
	_ = "STUB: not implemented"
	return
}

func (calc *AreaCentroidCalculator) setBasePoint(basePt geom.Coord) {
	_ = "STUB: not implemented"
	return
}

func (calc *AreaCentroidCalculator) addShell(pts []float64) { _ = "STUB: not implemented"; return }

func (calc *AreaCentroidCalculator) addHole(pts []float64) { _ = "STUB: not implemented"; return }

func (calc *AreaCentroidCalculator) addTriangle(p0, p1, p2 geom.Coord, isPositiveArea bool) {
	_ = "STUB: not implemented"
	return
}

// Returns three times the centroid of the triangle p1-p2-p3.
// The factor of 3 is left in to permit division to be avoided until later.
func centroid3(p1, p2, p3, c geom.Coord) { _ = "STUB: not implemented"; return }

// Returns twice the signed area of the triangle p1-p2-p3,
// positive if a,b,c are oriented ccw, and negative if cw.
func area2(p1, p2, p3 geom.Coord) float64 { _ = "STUB: not implemented"; return 0 }

// Adds the linear segments defined by an array of coordinates
// to the linear centroid accumulators.
// This is done in case the polygon(s) have zero-area,
// in which case the linear centroid is computed instead.
//
// Param pts - an array of Coords
func (calc *AreaCentroidCalculator) addLinearSegments(pts []float64) {
	_ = "STUB: not implemented"
	return
}
