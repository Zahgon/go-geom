// Package wkbcommon contains code common to WKB and EWKB encoding.
package wkbcommon

import (
	"encoding/binary"
	"io"
)

// Byte order IDs.
const (
	XDRID = 0
	NDRID = 1
)

// Byte orders.
var (
	XDR = binary.BigEndian
	NDR = binary.LittleEndian
)

// An ErrUnknownByteOrder is returned when an unknown byte order is encountered.
type ErrUnknownByteOrder byte

func (e ErrUnknownByteOrder) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrUnsupportedByteOrder is returned when an unsupported byte order is encountered.
type ErrUnsupportedByteOrder struct{}

func (e ErrUnsupportedByteOrder) Error() string { _ = "STUB: not implemented"; return "" }

// A Type is a WKB code.
type Type uint32

// An ErrUnknownType is returned when an unknown type is encountered.
type ErrUnknownType Type

func (e ErrUnknownType) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrUnsupportedType is returned when an unsupported type is encountered.
type ErrUnsupportedType Type

func (e ErrUnsupportedType) Error() string { _ = "STUB: not implemented"; return "" }

// An ErrUnexpectedType is returned when an unexpected type is encountered.
type ErrUnexpectedType struct {
	Got  any
	Want any
}

func (e ErrUnexpectedType) Error() string { _ = "STUB: not implemented"; return "" }

// MaxGeometryElements is the maximum number of elements that will be decoded
// at different levels. Its primary purpose is to prevent corrupt inputs from
// causing excessive memory allocations (which could be used as a denial of
// service attack).
//
// This is a variable, so you can override it in your application code by
// importing the `github.com/twpayne/go-geom/encoding/wkbcommon` module and
// setting the value of `wkbcommon.MaxGeometryElements`.
//
// FIXME This should be Codec-specific, not global.
// FIXME Consider overall per-geometry limit rather than per-level limit.
var MaxGeometryElements = [4]int{
	0,  // Unused
	-1, // LineString, LinearRing, and MultiPoint
	-1, // MultiLineString and Polygon
	-1, // MultiPolygon
}

// An ErrGeometryTooLarge is returned when the geometry is too large.
type ErrGeometryTooLarge struct {
	Level int
	N     int
	Limit int
}

func (e ErrGeometryTooLarge) Error() string { _ = "STUB: not implemented"; return "" }

// Geometry type IDs.
const (
	PointID              = 1
	LineStringID         = 2
	PolygonID            = 3
	MultiPointID         = 4
	MultiLineStringID    = 5
	MultiPolygonID       = 6
	GeometryCollectionID = 7
	PolyhedralSurfaceID  = 15
	TINID                = 16
	TriangleID           = 17
)

// ReadFlatCoords0 reads flat coordinates 0.
func ReadFlatCoords0(r io.Reader, byteOrder binary.ByteOrder, stride int) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFlatCoords1 reads flat coordinates 1.
func ReadFlatCoords1(r io.Reader, byteOrder binary.ByteOrder, stride int) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFlatCoords2 reads flat coordinates 2.
func ReadFlatCoords2(r io.Reader, byteOrder binary.ByteOrder, stride int) ([]float64, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// WriteFlatCoords0 writes flat coordinates 0.
func WriteFlatCoords0(w io.Writer, byteOrder binary.ByteOrder, coord []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFlatCoords1 writes flat coordinates 1.
func WriteFlatCoords1(w io.Writer, byteOrder binary.ByteOrder, coords []float64, stride int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFlatCoords2 writes flat coordinates 2.
func WriteFlatCoords2(w io.Writer, byteOrder binary.ByteOrder, flatCoords []float64, ends []int, stride int) error {
	_ = "STUB: not implemented"
	return nil
}
