package geom

// A MultiPolygon is a collection of Polygons.
type MultiPolygon struct {
	geom3
}

// NewMultiPolygon returns a new MultiPolygon with no Polygons.
func NewMultiPolygon(layout Layout) *MultiPolygon { _ = "STUB: not implemented"; return nil }

// NewMultiPolygonFlat returns a new MultiPolygon with the given flat coordinates.
func NewMultiPolygonFlat(layout Layout, flatCoords []float64, endss [][]int) *MultiPolygon {
	_ = "STUB: not implemented"
	return nil
}

// Area returns the sum of the area of the individual Polygons.
func (g *MultiPolygon) Area() float64 { _ = "STUB: not implemented"; return 0 }

// Clone returns a deep copy.
func (g *MultiPolygon) Clone() *MultiPolygon { _ = "STUB: not implemented"; return nil }

// Length returns the sum of the perimeters of the Polygons.
func (g *MultiPolygon) Length() float64 { _ = "STUB: not implemented"; return 0 }

// MustSetCoords sets the coordinates and panics on any error.
func (g *MultiPolygon) MustSetCoords(coords [][][]Coord) *MultiPolygon {
	_ = "STUB: not implemented"
	return nil
}

// NumPolygons returns the number of Polygons.
func (g *MultiPolygon) NumPolygons() int { _ = "STUB: not implemented"; return 0 }

// Polygon returns the ith Polygon.
func (g *MultiPolygon) Polygon(i int) *Polygon { _ = "STUB: not implemented"; return nil }

// Find the offset from the previous non-empty polygon element.

// Push appends a Polygon.
func (g *MultiPolygon) Push(p *Polygon) error { _ = "STUB: not implemented"; return nil }

// SetCoords sets the coordinates.
func (g *MultiPolygon) SetCoords(coords [][][]Coord) (*MultiPolygon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSRID sets the SRID of g.
func (g *MultiPolygon) SetSRID(srid int) *MultiPolygon { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *MultiPolygon) Swap(g2 *MultiPolygon) { _ = "STUB: not implemented"; return }
