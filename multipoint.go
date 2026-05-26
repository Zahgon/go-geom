package geom

// A MultiPoint is a collection of Points.
type MultiPoint struct {
	// To represent an MultiPoint that allows EMPTY elements, e.g.
	// MULTIPOINT ( EMPTY, POINT(1.0 1.0), EMPTY), we have to allow
	// record ends. If there is an empty point, ends[i] == ends[i-1].
	geom2
}

// NewMultiPoint returns a new, empty, MultiPoint.
func NewMultiPoint(layout Layout) *MultiPoint { _ = "STUB: not implemented"; return nil }

// NewMultiPointFlatOption represents an option that can be passed into
// NewMultiPointFlat.
type NewMultiPointFlatOption func(*MultiPoint)

// NewMultiPointFlatOptionWithEnds allows passing ends to NewMultiPointFlat,
// which allows the representation of empty points.
func NewMultiPointFlatOptionWithEnds(ends []int) NewMultiPointFlatOption {
	_ = "STUB: not implemented"
	return *new(NewMultiPointFlatOption)
}

// NewMultiPointFlat returns a new MultiPoint with the given flat coordinates.
// Assumes no points are empty by default. Use `NewMultiPointFlatOptionWithEnds`
// to specify empty points.
func NewMultiPointFlat(
	layout Layout, flatCoords []float64, opts ...NewMultiPointFlatOption,
) *MultiPoint {
	_ = "STUB: not implemented"
	return nil
}

// If no ends are provided, assume all points are non empty.

// Area returns the area of g, i.e. zero.
func (g *MultiPoint) Area() float64 {
	_ = "STUB: not implemented"

	// Clone returns a deep copy.
	return 0
}

func (g *MultiPoint) Clone() *MultiPoint { _ = "STUB: not implemented"; return nil }

// Length returns zero.
func (g *MultiPoint) Length() float64 {
	_ = "STUB: not implemented"

	// MustSetCoords sets the coordinates and panics on any error.
	return 0
}

func (g *MultiPoint) MustSetCoords(coords []Coord) *MultiPoint {
	_ = "STUB: not implemented"
	return nil
}

// Coord returns the ith coord of g.
func (g *MultiPoint) Coord(i int) Coord { _ = "STUB: not implemented"; return *new(Coord) }

// SetCoords sets the coordinates.
func (g *MultiPoint) SetCoords(coords []Coord) (*MultiPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Coords unpacks and returns all of g's coordinates.
func (g *MultiPoint) Coords() []Coord { _ = "STUB: not implemented"; return nil }

// NumCoords returns the number of coordinates in g.
func (g *MultiPoint) NumCoords() int {
	_ = "STUB: not implemented"

	// SetSRID sets the SRID of g.
	return 0
}

func (g *MultiPoint) SetSRID(srid int) *MultiPoint { _ = "STUB: not implemented"; return nil }

// NumPoints returns the number of Points.
func (g *MultiPoint) NumPoints() int {
	_ = "STUB: not implemented"

	// Point returns the ith Point.
	return 0
}

func (g *MultiPoint) Point(i int) *Point { _ = "STUB: not implemented"; return nil }

// Push appends a point.
func (g *MultiPoint) Push(p *Point) error { _ = "STUB: not implemented"; return nil }

// Swap swaps the values of g and g2.
func (g *MultiPoint) Swap(g2 *MultiPoint) { _ = "STUB: not implemented"; return }
