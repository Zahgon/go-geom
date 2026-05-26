package wkt

import (
	"strings"

	"github.com/twpayne/go-geom"
)

// Encode translates a geometry to the corresponding WKT.
func (e *Encoder) Encode(g geom.T) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *Encoder) write(sb *strings.Builder, g geom.T) error { _ = "STUB: not implemented"; return nil }

// Special case for empty GeometryCollections

func (e *Encoder) writeCoord(sb *strings.Builder, coord []float64) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:interfacer
func (e *Encoder) writeEMPTY(sb *strings.Builder) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) writeFlatCoords0(sb *strings.Builder, flatCoords []float64, stride int) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) writeFlatCoords1(sb *strings.Builder, flatCoords []float64, stride int) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) writeFlatCoords1Ends(
	sb *strings.Builder, flatCoords []float64, start int, ends []int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) writeFlatCoords2(
	sb *strings.Builder, flatCoords []float64, start int, ends []int, stride int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) writeFlatCoords3(
	sb *strings.Builder, flatCoords []float64, endss [][]int, stride int,
) error {
	_ = "STUB: not implemented"
	return nil
}
