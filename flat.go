package geom

type geom0 struct {
	layout     Layout
	stride     int
	flatCoords []float64
	srid       int
}

type geom1 struct {
	geom0
}

type geom2 struct {
	geom1
	ends []int
}

type geom3 struct {
	geom1
	endss [][]int
}

// Bounds returns the bounds of g.
func (g *geom0) Bounds() *Bounds { _ = "STUB: not implemented"; return nil }

// Coords returns all the coordinates in g, i.e. a single coordinate.
func (g *geom0) Coords() Coord { _ = "STUB: not implemented"; return *new(Coord) }

// Empty returns true if g contains no coordinates.
func (g *geom0) Empty() bool { _ = "STUB: not implemented"; return false }

// Ends returns the end indexes of sub-structures of g, i.e. an empty slice.
func (g *geom0) Ends() []int {
	_ = "STUB: not implemented"

	// Endss returns the end indexes of sub-sub-structures of g, i.e. an empty
	// slice.
	return nil
}

func (g *geom0) Endss() [][]int {
	_ = "STUB: not implemented"

	// FlatCoords returns the flat coordinates of g.
	return nil
}

func (g *geom0) FlatCoords() []float64 { _ = "STUB: not implemented"; return nil }

// Layout returns g's layout.
func (g *geom0) Layout() Layout {
	_ = "STUB: not implemented"

	// NumCoords returns the number of coordinates in g, i.e. 1.
	return *new(Layout)
}

func (g *geom0) NumCoords() int {
	_ = "STUB: not implemented"

	// Reserve reserves space in g for n coordinates.
	return 0
}

func (g *geom0) Reserve(n int) { _ = "STUB: not implemented"; return }

// SRID returns g's SRID.
func (g *geom0) SRID() int { _ = "STUB: not implemented"; return 0 }

func (g *geom0) setCoords(coords0 []float64) error { _ = "STUB: not implemented"; return nil }

// Stride returns g's stride.
func (g *geom0) Stride() int { _ = "STUB: not implemented"; return 0 }

func (g *geom0) verify() error { _ = "STUB: not implemented"; return nil }

// Coord returns the ith coord of g.
func (g *geom1) Coord(i int) Coord { _ = "STUB: not implemented"; return *new(Coord) }

// Coords unpacks and returns all of g's coordinates.
func (g *geom1) Coords() []Coord { _ = "STUB: not implemented"; return nil }

// NumCoords returns the number of coordinates in g.
func (g *geom1) NumCoords() int { _ = "STUB: not implemented"; return 0 }

// Reverse reverses the order of g's coordinates.
func (g *geom1) Reverse() { _ = "STUB: not implemented"; return }

func (g *geom1) setCoords(coords1 []Coord) error { _ = "STUB: not implemented"; return nil }

func (g *geom1) verify() error { _ = "STUB: not implemented"; return nil }

// Coords returns all of g's coordinates.
func (g *geom2) Coords() [][]Coord { _ = "STUB: not implemented"; return nil }

// Ends returns the end indexes of all sub-structures in g.
func (g *geom2) Ends() []int {
	_ = "STUB: not implemented"

	// Reverse reverses the order of coordinates for each sub-structure in g.
	return nil
}

func (g *geom2) Reverse() { _ = "STUB: not implemented"; return }

func (g *geom2) setCoords(coords2 [][]Coord) error { _ = "STUB: not implemented"; return nil }

func (g *geom2) verify() error { _ = "STUB: not implemented"; return nil }

// Coords returns all the coordinates in g.
func (g *geom3) Coords() [][][]Coord { _ = "STUB: not implemented"; return nil }

// Endss returns a list of all the sub-sub-structures in g.
func (g *geom3) Endss() [][]int {
	_ = "STUB: not implemented"

	// Reverse reverses the order of coordinates for each sub-sub-structure in g.
	return nil
}

func (g *geom3) Reverse() { _ = "STUB: not implemented"; return }

func (g *geom3) setCoords(coords3 [][][]Coord) error { _ = "STUB: not implemented"; return nil }

func (g *geom3) verify() error { _ = "STUB: not implemented"; return nil }

func doubleArea1(flatCoords []float64, offset, end, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func doubleArea2(flatCoords []float64, offset int, ends []int, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func doubleArea3(flatCoords []float64, offset int, endss [][]int, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func deflate0(flatCoords []float64, c Coord, stride int) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deflate1(flatCoords []float64, coords1 []Coord, stride int) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deflate2(
	flatCoords []float64, ends []int, coords2 [][]Coord, stride int,
) ([]float64, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func deflate3(
	flatCoords []float64, endss [][]int, coords3 [][][]Coord, stride int,
) ([]float64, [][]int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func inflate0(flatCoords []float64, offset, end, stride int) Coord {
	_ = "STUB: not implemented"
	return *new(Coord)
}

func inflate1(flatCoords []float64, offset, end, stride int) []Coord {
	_ = "STUB: not implemented"
	return nil
}

func inflate2(flatCoords []float64, offset int, ends []int, stride int) [][]Coord {
	_ = "STUB: not implemented"
	return nil
}

func inflate3(flatCoords []float64, offset int, endss [][]int, stride int) [][][]Coord {
	_ = "STUB: not implemented"
	return nil
}

func length1(flatCoords []float64, offset, end, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func length2(flatCoords []float64, offset int, ends []int, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func length3(flatCoords []float64, offset int, endss [][]int, stride int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func reverse1(flatCoords []float64, offset, end, stride int) { _ = "STUB: not implemented"; return }

func reverse2(flatCoords []float64, offset int, ends []int, stride int) {
	_ = "STUB: not implemented"
	return
}

func reverse3(flatCoords []float64, offset int, endss [][]int, stride int) {
	_ = "STUB: not implemented"
	return
}
