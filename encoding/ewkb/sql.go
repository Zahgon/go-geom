package ewkb

import (
	"database/sql/driver"

	"github.com/twpayne/go-geom"
)

// ErrExpectedByteSlice is returned when a []byte is expected.
type ErrExpectedByteSlice struct {
	Value any
}

func (e ErrExpectedByteSlice) Error() string { _ = "STUB: not implemented"; return "" }

// A Point is a EWKB-encoded Point that implements the sql.Scanner and
// driver.Value interfaces.
type Point struct {
	*geom.Point
}

// A LineString is a EWKB-encoded LineString that implements the
// sql.Scanner and driver.Value interfaces.
type LineString struct {
	*geom.LineString
}

// A Polygon is a EWKB-encoded Polygon that implements the sql.Scanner and
// driver.Value interfaces.
type Polygon struct {
	*geom.Polygon
}

// A MultiPoint is a EWKB-encoded MultiPoint that implements the
// sql.Scanner and driver.Value interfaces.
type MultiPoint struct {
	*geom.MultiPoint
}

// A MultiLineString is a EWKB-encoded MultiLineString that implements the
// sql.Scanner and driver.Value interfaces.
type MultiLineString struct {
	*geom.MultiLineString
}

// A MultiPolygon is a EWKB-encoded MultiPolygon that implements the
// sql.Scanner and driver.Value interfaces.
type MultiPolygon struct {
	*geom.MultiPolygon
}

// A GeometryCollection is a EWKB-encoded GeometryCollection that implements
// the sql.Scanner and driver.Value interfaces.
type GeometryCollection struct {
	*geom.GeometryCollection
}

// Scan scans from a []byte.
func (p *Point) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if p has a value.
func (p *Point) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of p.
func (p *Point) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (ls *LineString) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid return true if ls has a value.
func (ls *LineString) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of ls.
func (ls *LineString) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (p *Polygon) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if p has a value.
func (p *Polygon) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of p.
func (p *Polygon) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (mp *MultiPoint) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if mp has a value.
func (mp *MultiPoint) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of mp.
func (mp *MultiPoint) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (mls *MultiLineString) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if mls has a value.
func (mls *MultiLineString) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of mls.
func (mls *MultiLineString) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (mp *MultiPolygon) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if mp has a value.
func (mp *MultiPolygon) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of mp.
func (mp *MultiPolygon) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

// Scan scans from a []byte.
func (gc *GeometryCollection) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Valid returns true if gc has a value.
func (gc *GeometryCollection) Valid() bool { _ = "STUB: not implemented"; return false }

// Value returns the EWKB encoding of gc.
func (gc *GeometryCollection) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

//nolint:nilnil

func value(g geom.T) (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
