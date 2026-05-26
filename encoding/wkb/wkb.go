// Package wkb implements Well Known Binary encoding and decoding.
//
// If you are encoding geometries in WKB to send to PostgreSQL/PostGIS, then
// you must specify binary_parameters=yes in the data source name that you pass
// to sql.Open.
package wkb

import (
	"encoding/binary"
	"io"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
)

var (
	// XDR is big endian.
	XDR = wkbcommon.XDR
	// NDR is little endian.
	NDR = wkbcommon.NDR
)

const (
	wkbXYID   = 0
	wkbXYZID  = 1000
	wkbXYMID  = 2000
	wkbXYZMID = 3000
)

// Read reads an arbitrary geometry from r.
func Read(r io.Reader, opts ...wkbcommon.WKBOption) (geom.T, error) {
	_ = "STUB: not implemented"
	return *new(geom.T), nil
}

// If EMPTY, mark the collection with a fixed layout to differentiate
// GEOMETRYCOLLECTION EMPTY between 2D/Z/M/ZM.

// Unmarshal unmrshals an arbitrary geometry from a []byte.
func Unmarshal(data []byte, opts ...wkbcommon.WKBOption) (geom.T, error) {
	_ = "STUB: not implemented"
	return *new(geom.T), nil
}

// Write writes an arbitrary geometry to w.
func Write(w io.Writer, byteOrder binary.ByteOrder, g geom.T, opts ...wkbcommon.WKBOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Special case for empty GeometryCollections

// Marshal marshals an arbitrary geometry to a []byte.
func Marshal(g geom.T, byteOrder binary.ByteOrder, opts ...wkbcommon.WKBOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
