package geom

// A Bounds represents a multi-dimensional bounding box.
type Bounds struct {
	layout Layout
	min    Coord
	max    Coord
}

// NewBounds creates a new Bounds.
func NewBounds(layout Layout) *Bounds { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of b.
func (b *Bounds) Clone() *Bounds { _ = "STUB: not implemented"; return nil }

// Extend extends b to include geometry g.
func (b *Bounds) Extend(g T) *Bounds { _ = "STUB: not implemented"; return nil }

// IsEmpty returns true if b is empty.
func (b *Bounds) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Layout returns b's layout.
func (b *Bounds) Layout() Layout {
	_ = "STUB: not implemented"

	// Max returns the maximum value in dimension dim.
	return *new(Layout)
}

func (b *Bounds) Max(dim int) float64 {
	_ = "STUB: not implemented"

	// Min returns the minimum value in dimension dim.
	return 0
}

func (b *Bounds) Min(dim int) float64 {
	_ = "STUB: not implemented"

	// Overlaps returns true if b overlaps b2 in layout.
	return 0
}

func (b *Bounds) Overlaps(layout Layout, b2 *Bounds) bool { _ = "STUB: not implemented"; return false }

// Polygon returns b as a two-dimensional Polygon.
func (b *Bounds) Polygon() *Polygon { _ = "STUB: not implemented"; return nil }

// Set sets the minimum and maximum values. args must be an even number of
// values: the first half are the minimum values for each dimension and the
// second half are the maximum values for each dimension. If necessary, the
// layout of b will be extended to cover all the supplied dimensions implied by
// args.
func (b *Bounds) Set(args ...float64) *Bounds { _ = "STUB: not implemented"; return nil }

// SetCoords sets the minimum and maximum values of the Bounds.
func (b *Bounds) SetCoords(minCoord, maxCoord Coord) *Bounds { _ = "STUB: not implemented"; return nil }

// OverlapsPoint determines if the bounding box overlaps the point (point is
// within or on the border of the bounds).
func (b *Bounds) OverlapsPoint(layout Layout, point Coord) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *Bounds) extendFlatCoords(flatCoords []float64, offset, end, stride int) *Bounds {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bounds) extendLayout(layout Layout) { _ = "STUB: not implemented"; return }

func (b *Bounds) extendStride(stride int) { _ = "STUB: not implemented"; return }

func (b *Bounds) extendXYZMFlatCoordsWithXYM(flatCoords []float64, offset, end int) *Bounds {
	_ = "STUB: not implemented"
	return nil
}
