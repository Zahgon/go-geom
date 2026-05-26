package geom

// PointEmptyCoordHex is the hex representation of a NaN that represents
// an empty coord in a shape.
const PointEmptyCoordHex = 0x7FF8000000000000

// PointEmptyCoord is the NaN float64 representation of the empty coordinate.
func PointEmptyCoord() float64 { _ = "STUB: not implemented"; return 0 }

// A Point represents a single point.
type Point struct {
	geom0
}

// NewPoint allocates a new Point with layout l and all values zero.
func NewPoint(l Layout) *Point { _ = "STUB: not implemented"; return nil }

// NewPointEmpty allocates a new Point with no coordinates.
func NewPointEmpty(l Layout) *Point { _ = "STUB: not implemented"; return nil }

// NewPointFlat allocates a new Point with layout l and flat coordinates flatCoords.
func NewPointFlat(l Layout, flatCoords []float64) *Point { _ = "STUB: not implemented"; return nil }

// NewPointFlatMaybeEmpty returns a new point, checking whether the point may be empty
// by checking wther all the points are NaN.
func NewPointFlatMaybeEmpty(layout Layout, flatCoords []float64) *Point {
	_ = "STUB: not implemented"
	return nil
}

// Area returns g's area, i.e. zero.
func (g *Point) Area() float64 {
	_ = "STUB: not implemented"

	// Clone returns a copy of g that does not alias g.
	return 0
}

func (g *Point) Clone() *Point { _ = "STUB: not implemented"; return nil }

// Length returns the length of g, i.e. zero.
func (g *Point) Length() float64 {
	_ = "STUB: not implemented"

	// MustSetCoords is like SetCoords but panics on any error.
	return 0
}

func (g *Point) MustSetCoords(coords Coord) *Point { _ = "STUB: not implemented"; return nil }

// SetCoords sets the coordinates of g.
func (g *Point) SetCoords(coords Coord) (*Point, error) { _ = "STUB: not implemented"; return nil, nil }

// SetSRID sets the SRID of g.
func (g *Point) SetSRID(srid int) *Point { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *Point) Swap(g2 *Point) {
	_ = "STUB: not implemented"

	// X returns g's X-coordinate.
	return
}

func (g *Point) X() float64 { _ = "STUB: not implemented"; return 0 }

// Y returns g's Y-coordinate.
func (g *Point) Y() float64 { _ = "STUB: not implemented"; return 0 }

// Z returns g's Z-coordinate, or zero if g has no Z-coordinate.
func (g *Point) Z() float64 { _ = "STUB: not implemented"; return 0 }

// M returns g's M-coordinate, or zero if g has no M-coordinate.
func (g *Point) M() float64 { _ = "STUB: not implemented"; return 0 }
