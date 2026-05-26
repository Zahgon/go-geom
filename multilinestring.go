package geom

// A MultiLineString is a collection of LineStrings.
type MultiLineString struct {
	geom2
}

// NewMultiLineString returns a new MultiLineString with no LineStrings.
func NewMultiLineString(layout Layout) *MultiLineString { _ = "STUB: not implemented"; return nil }

// NewMultiLineStringFlat returns a new MultiLineString with the given flat coordinates.
func NewMultiLineStringFlat(layout Layout, flatCoords []float64, ends []int) *MultiLineString {
	_ = "STUB: not implemented"
	return nil
}

// Area returns the area of g, i.e. 0.
func (g *MultiLineString) Area() float64 {
	_ = "STUB: not implemented"

	// Clone returns a deep copy.
	return 0
}

func (g *MultiLineString) Clone() *MultiLineString { _ = "STUB: not implemented"; return nil }

// Length returns the sum of the length of the LineStrings.
func (g *MultiLineString) Length() float64 { _ = "STUB: not implemented"; return 0 }

// LineString returns the ith LineString.
func (g *MultiLineString) LineString(i int) *LineString { _ = "STUB: not implemented"; return nil }

// MustSetCoords sets the coordinates and panics on any error.
func (g *MultiLineString) MustSetCoords(coords [][]Coord) *MultiLineString {
	_ = "STUB: not implemented"
	return nil
}

// NumLineStrings returns the number of LineStrings.
func (g *MultiLineString) NumLineStrings() int {
	_ = "STUB: not implemented"

	// Push appends a LineString.
	return 0
}

func (g *MultiLineString) Push(ls *LineString) error { _ = "STUB: not implemented"; return nil }

// SetCoords sets the coordinates.
func (g *MultiLineString) SetCoords(coords [][]Coord) (*MultiLineString, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSRID sets the SRID of g.
func (g *MultiLineString) SetSRID(srid int) *MultiLineString { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *MultiLineString) Swap(g2 *MultiLineString) { _ = "STUB: not implemented"; return }
