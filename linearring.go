package geom

// A LinearRing is a linear ring.
type LinearRing struct {
	geom1
}

// NewLinearRing returns a new LinearRing with no coordinates.
func NewLinearRing(layout Layout) *LinearRing { _ = "STUB: not implemented"; return nil }

// NewLinearRingFlat returns a new LinearRing with the given flat coordinates.
func NewLinearRingFlat(layout Layout, flatCoords []float64) *LinearRing {
	_ = "STUB: not implemented"
	return nil
}

// Area returns the area.
func (g *LinearRing) Area() float64 { _ = "STUB: not implemented"; return 0 }

// Clone returns a deep copy.
func (g *LinearRing) Clone() *LinearRing { _ = "STUB: not implemented"; return nil }

// Length returns the length of the perimeter.
func (g *LinearRing) Length() float64 { _ = "STUB: not implemented"; return 0 }

// MustSetCoords sets the coordinates and panics if there is any error.
func (g *LinearRing) MustSetCoords(coords []Coord) *LinearRing {
	_ = "STUB: not implemented"
	return nil
}

// SetCoords sets the coordinates.
func (g *LinearRing) SetCoords(coords []Coord) (*LinearRing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSRID sets the SRID of g.
func (g *LinearRing) SetSRID(srid int) *LinearRing { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *LinearRing) Swap(g2 *LinearRing) { _ = "STUB: not implemented"; return }
