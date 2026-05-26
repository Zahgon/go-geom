package geom

// A GeometryCollection is a collection of arbitrary geometries with the same
// SRID.
type GeometryCollection struct {
	layout Layout
	geoms  []T
	srid   int
}

// NewGeometryCollection returns a new empty GeometryCollection.
func NewGeometryCollection() *GeometryCollection { _ = "STUB: not implemented"; return nil }

// Geom returns the ith geometry in g.
func (g *GeometryCollection) Geom(i int) T {
	_ = "STUB: not implemented"

	// Geoms returns the geometries in g.
	return *new(T)
}

func (g *GeometryCollection) Geoms() []T {
	_ = "STUB: not implemented"

	// Layout returns the smallest layout that covers all of the layouts in g's
	// geometries.
	return nil
}

func (g *GeometryCollection) Layout() Layout { _ = "STUB: not implemented"; return *new(Layout) }

// NumGeoms returns the number of geometries in g.
func (g *GeometryCollection) NumGeoms() int { _ = "STUB: not implemented"; return 0 }

// Stride returns the stride of g's layout.
func (g *GeometryCollection) Stride() int { _ = "STUB: not implemented"; return 0 }

// Bounds returns the bounds of all the geometries in g.
func (g *GeometryCollection) Bounds() *Bounds {
	_ = "STUB: not implemented"
	// FIXME this needs work for mixing layouts, e.g. XYZ and XYM
	return nil
}

// Empty returns true if the collection is empty.
// This can return true if the GeometryCollection contains multiple Geometry objects
// which are all empty.
func (g *GeometryCollection) Empty() bool { _ = "STUB: not implemented"; return false }

// FlatCoords panics.
func (g *GeometryCollection) FlatCoords() []float64 { _ = "STUB: not implemented"; return nil }

// Ends panics.
func (g *GeometryCollection) Ends() []int { _ = "STUB: not implemented"; return nil }

// Endss panics.
func (g *GeometryCollection) Endss() [][]int { _ = "STUB: not implemented"; return nil }

// SRID returns g's SRID.
func (g *GeometryCollection) SRID() int {
	_ = "STUB: not implemented"

	// MustPush pushes gs to g. It panics on any error.
	return 0
}

func (g *GeometryCollection) MustPush(gs ...T) *GeometryCollection {
	_ = "STUB: not implemented"
	return nil
}

// CheckLayout checks all geometries in the collection match the given
// layout.
func (g *GeometryCollection) CheckLayout(layout Layout) error {
	_ = "STUB: not implemented"
	return nil
}

// MustSetLayout sets g's layout. It panics on any error.
func (g *GeometryCollection) MustSetLayout(layout Layout) *GeometryCollection {
	_ = "STUB: not implemented"
	return nil
}

// Push appends geometries.
func (g *GeometryCollection) Push(gs ...T) error { _ = "STUB: not implemented"; return nil }

// SetLayout sets g's layout.
func (g *GeometryCollection) SetLayout(layout Layout) error { _ = "STUB: not implemented"; return nil }

// SetSRID sets g's SRID and the SRID of all its elements.
func (g *GeometryCollection) SetSRID(srid int) *GeometryCollection {
	_ = "STUB: not implemented"
	return nil
}
