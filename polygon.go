package geom

// A Polygon represents a polygon as a collection of LinearRings. The first
// LinearRing is the outer boundary. Subsequent LinearRings are inner
// boundaries (holes).
type Polygon struct {
	geom2
}

// NewPolygon returns a new, empty, Polygon.
func NewPolygon(layout Layout) *Polygon { _ = "STUB: not implemented"; return nil }

// NewPolygonFlat returns a new Polygon with the given flat coordinates.
func NewPolygonFlat(layout Layout, flatCoords []float64, ends []int) *Polygon {
	_ = "STUB: not implemented"
	return nil
}

// Area returns the area.
func (g *Polygon) Area() float64 { _ = "STUB: not implemented"; return 0 }

// Clone returns a deep copy.
func (g *Polygon) Clone() *Polygon { _ = "STUB: not implemented"; return nil }

// Length returns the perimter.
func (g *Polygon) Length() float64 { _ = "STUB: not implemented"; return 0 }

// LinearRing returns the ith LinearRing.
func (g *Polygon) LinearRing(i int) *LinearRing { _ = "STUB: not implemented"; return nil }

// MustSetCoords sets the coordinates and panics on any error.
func (g *Polygon) MustSetCoords(coords [][]Coord) *Polygon { _ = "STUB: not implemented"; return nil }

// NumLinearRings returns the number of LinearRings.
func (g *Polygon) NumLinearRings() int {
	_ = "STUB: not implemented"

	// Push appends a LinearRing.
	return 0
}

func (g *Polygon) Push(lr *LinearRing) error { _ = "STUB: not implemented"; return nil }

// SetCoords sets the coordinates.
func (g *Polygon) SetCoords(coords [][]Coord) (*Polygon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSRID sets the SRID of g.
func (g *Polygon) SetSRID(srid int) *Polygon { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *Polygon) Swap(g2 *Polygon) { _ = "STUB: not implemented"; return }
