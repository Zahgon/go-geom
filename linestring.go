package geom

// A LineString represents a single, unbroken line, linearly interpreted
// between zero or more control points.
type LineString struct {
	geom1
}

// NewLineString returns a new LineString with layout l and no control points.
func NewLineString(l Layout) *LineString { _ = "STUB: not implemented"; return nil }

// NewLineStringFlat returns a new LineString with layout l and control points
// flatCoords.
func NewLineStringFlat(layout Layout, flatCoords []float64) *LineString {
	_ = "STUB: not implemented"
	return nil
}

// Area returns the area of g, i.e. zero.
func (g *LineString) Area() float64 {
	_ = "STUB: not implemented"

	// Clone returns a copy of g that does not alias g.
	return 0
}

func (g *LineString) Clone() *LineString { _ = "STUB: not implemented"; return nil }

// Interpolate returns the index and delta of val in dimension dim.
func (g *LineString) Interpolate(val float64, dim int) (int, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Length returns the length of g.
func (g *LineString) Length() float64 { _ = "STUB: not implemented"; return 0 }

// MustSetCoords is like SetCoords but it panics on any error.
func (g *LineString) MustSetCoords(coords []Coord) *LineString {
	_ = "STUB: not implemented"
	return nil
}

// SetCoords sets the coordinates of g.
func (g *LineString) SetCoords(coords []Coord) (*LineString, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSRID sets the SRID of g.
func (g *LineString) SetSRID(srid int) *LineString { _ = "STUB: not implemented"; return nil }

// SubLineString returns a LineString from starts at index start and stops at
// index stop of g. The returned LineString aliases g.
func (g *LineString) SubLineString(start, stop int) *LineString {
	_ = "STUB: not implemented"
	return nil
}

// Swap swaps the values of g and g2.
func (g *LineString) Swap(g2 *LineString) { _ = "STUB: not implemented"; return }
