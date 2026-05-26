package wkb

import (
	"database/sql/driver"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
)

// ErrExpectedByteSlice is returned when a []byte is expected.
type ErrExpectedByteSlice struct {
	Value any
}

func (e ErrExpectedByteSlice) Error() string { _ = "STUB: not implemented"; return "" }

// A Geom is a WKB-ecoded Geometry that implements the sql.Scanner and
// driver.Value interfaces.
// It can be used when the geometry shape is not defined.
type Geom struct {
	geom.T
	opts []wkbcommon.WKBOption
}

// A Point is a WKB-encoded Point that implements the sql.Scanner and
// driver.Valuer interfaces.
type Point struct {
	*geom.Point
	opts []wkbcommon.WKBOption
}

// A LineString is a WKB-encoded LineString that implements the sql.Scanner and
// driver.Valuer interfaces.
type LineString struct {
	*geom.LineString
	opts []wkbcommon.WKBOption
}

// A Polygon is a WKB-encoded Polygon that implements the sql.Scanner and
// driver.Valuer interfaces.
type Polygon struct {
	*geom.Polygon
	opts []wkbcommon.WKBOption
}

// A MultiPoint is a WKB-encoded MultiPoint that implements the sql.Scanner and
// driver.Valuer interfaces.
type MultiPoint struct {
	*geom.MultiPoint
	opts []wkbcommon.WKBOption
}

// A MultiLineString is a WKB-encoded MultiLineString that implements the
// sql.Scanner and driver.Valuer interfaces.
type MultiLineString struct {
	*geom.MultiLineString
	opts []wkbcommon.WKBOption
}

// A MultiPolygon is a WKB-encoded MultiPolygon that implements the sql.Scanner
// and driver.Valuer interfaces.
type MultiPolygon struct {
	*geom.MultiPolygon
	opts []wkbcommon.WKBOption
}

// A GeometryCollection is a WKB-encoded GeometryCollection that implements the
// sql.Scanner and driver.Valuer interfaces.
type GeometryCollection struct {
	*geom.GeometryCollection
	opts []wkbcommon.WKBOption
}

// Scan scans from a []byte.
func (g *Geom) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// NOTE(tb) other Scanners do not check the len of b, is it really useful ?

// Value returns the WKB encoding of g.
func (g *Geom) Value() (driver.Value, error) {
	_ = "STUB: not implemented"

	// Geom returns the underlying geom.T.
	return *new(driver.Value), nil
}

func (g *Geom) Geom() geom.T {
	_ = "STUB: not implemented"

	// Scan scans from a []byte.
	return *new(geom.T)
}

func (p *Point) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of p.
func (p *Point) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *

	// Scan scans from a []byte.
	new(driver.Value), nil
}

func (ls *LineString) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of ls.
func (ls *LineString) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scans from a []byte.
func (p *Polygon) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of p.
func (p *Polygon) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *

	// Scan scans from a []byte.
	new(driver.Value), nil
}

func (mp *MultiPoint) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of mp.
func (mp *MultiPoint) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scans from a []byte.
func (mls *MultiLineString) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of mls.
func (mls *MultiLineString) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scans from a []byte.
func (mp *MultiPolygon) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of mp.
func (mp *MultiPolygon) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scans from a []byte.
func (gc *GeometryCollection) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value returns the WKB encoding of gc.
func (gc *GeometryCollection) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func value(g geom.T) (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
