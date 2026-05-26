// Package geojson implements GeoJSON encoding and decoding.
package geojson

import (
	"encoding/json"
	"reflect"

	geom "github.com/twpayne/go-geom"
)

var nullGeometry = []byte("null")

// DefaultLayout is the default layout for empty geometries.
// FIXME This should be Codec-specific, not global.
var DefaultLayout = geom.XY

// ErrDimensionalityTooLow is returned when the dimensionality is too low.
type ErrDimensionalityTooLow int

func (e ErrDimensionalityTooLow) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUnsupportedType is returned when the type is unsupported.
type ErrUnsupportedType string

func (e ErrUnsupportedType) Error() string { _ = "STUB: not implemented"; return "" }

// CRS is a deprecated field but still populated in some programs (e.g. PostGIS).
// See https://geojson.org/geojson-spec for original specification of CRS.
type CRS struct {
	Type       string         `json:"type"`
	Properties map[string]any `json:"properties"`
}

// A Geometry is a geometry in GeoJSON format.
type Geometry struct {
	Type        string           `json:"type"`
	BBox        *json.RawMessage `json:"bbox,omitempty"`
	CRS         *CRS             `json:"crs,omitempty"`
	Coordinates *json.RawMessage `json:"coordinates,omitempty"`
	Geometries  *json.RawMessage `json:"geometries,omitempty"`
}

// A Feature is a GeoJSON Feature.
type Feature struct {
	ID         string
	BBox       *geom.Bounds
	Geometry   geom.T
	Properties map[string]any
}

type geojsonFeature struct {
	Type       string         `json:"type"`
	ID         any            `json:"id,omitempty"`
	BBox       []float64      `json:"bbox,omitempty"`
	Geometry   *Geometry      `json:"geometry"`
	Properties map[string]any `json:"properties"`
}

// A FeatureCollection is a GeoJSON FeatureCollection.
type FeatureCollection struct {
	BBox     *geom.Bounds
	Features []*Feature
}

type geojsonFeatureCollection struct {
	Type     string     `json:"type"`
	BBox     []float64  `json:"bbox,omitempty"`
	Features []*Feature `json:"features"`
}

func guessLayout0(coords0 []float64) (geom.Layout, error) {
	_ = "STUB: not implemented"
	return *new(geom.Layout), nil
}

func guessLayout1(coords1 []geom.Coord) (geom.Layout, error) {
	_ = "STUB: not implemented"
	return *new(geom.Layout), nil
}

func guessLayout2(coords2 [][]geom.Coord) (geom.Layout, error) {
	_ = "STUB: not implemented"
	return *new(geom.Layout), nil
}

func guessLayout3(coords3 [][][]geom.Coord) (geom.Layout, error) {
	_ = "STUB: not implemented"
	return *new(geom.Layout), nil
}

// Decode decodes g to a geometry.
func (g *Geometry) Decode() (geom.T, error) { _ = "STUB: not implemented"; return *new(geom.T), nil }

//nolint:nilnil

// EncodeGeometryOption applies extra metadata to the Geometry GeoJSON encoding.
type EncodeGeometryOption struct {
	onGeometryHandler func(*Geometry, geom.T, ...EncodeGeometryOption) error
	onFloat64Handler  func(any) any
}

// nestedFloat64WithMaxDecimalDigits is a wrapper around any nested array
// of float64s that will marshal into JSON with the maximum JSON digits.
type nestedFloat64WithMaxDecimalDigits struct {
	obj              any
	maxDecimalDigits int
}

// MarshalJSON implements the json.Marshaller interface.
func (c *nestedFloat64WithMaxDecimalDigits) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// marshalJSON is a helper routine that recurses down slices of float64s,
// appending float64 to a JSON list structure.
func (c *nestedFloat64WithMaxDecimalDigits) marshalJSON(
	buf []byte, val reflect.Value,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

// encodeJSONFloat64WithMaxDecimalDigits is an option implementation that converts slices of float64s
// to round to the maxDecimalDigits if necessary.
func encodeJSONFloat64WithMaxDecimalDigits(maxDecimalDigits int) func(any) any {
	_ = "STUB: not implemented"
	return nil
}

// EncodeGeometryWithBBox adds a bbox field to the Geometry GeoJSON encoding.
func EncodeGeometryWithBBox() EncodeGeometryOption {
	_ = "STUB: not implemented"
	return *new(EncodeGeometryOption)
}

// EncodeGeometryWithCRS adds the crs field to the Geometry GeoJSON encoding.
func EncodeGeometryWithCRS(crs *CRS) EncodeGeometryOption {
	_ = "STUB: not implemented"
	return *new(EncodeGeometryOption)
}

// EncodeGeometryWithMaxDecimalDigits encodes the Geometry with maximum decimal digits
// in the JSON representation.
func EncodeGeometryWithMaxDecimalDigits(maxDecimalDigits int) EncodeGeometryOption {
	_ = "STUB: not implemented"
	return *new(EncodeGeometryOption)
}

// Encode encodes g as a GeoJSON geometry.
func Encode(g geom.T, opts ...EncodeGeometryOption) (*Geometry, error) {
	_ = "STUB: not implemented"
	return nil, nil

	//nolint:nilnil
}

// encode encodes the geometry assuming it is not nil.
func encode(g geom.T, opts ...EncodeGeometryOption) (*Geometry, error) {
	_ = "STUB: not implemented"
	return nil, nil

	//nolint:nilnil
}

// Marshal marshals an arbitrary geometry to a []byte.
func Marshal(g geom.T, opts ...EncodeGeometryOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal unmarshalls a []byte to an arbitrary geometry.
func Unmarshal(data []byte, g *geom.T) error { _ = "STUB: not implemented"; return nil }

// decodeBBox decodes bb into a Bounds.
func decodeBBox(bb []float64) (*geom.Bounds, error) { _ = "STUB: not implemented"; return nil, nil }

// encodeBBox encodes b as a GeoJson Bounding Box.
func encodeBBox(b *geom.Bounds) ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalJSON implements json.Marshaler.MarshalJSON.
func (f *Feature) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Empty ID should be nil. Not a nil interface{} value.

// UnmarshalJSON implements json.Unmarshaler.UnmarshalJSON.
func (f *Feature) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.MarshalJSON.
func (fc *FeatureCollection) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON implements json.Unmarshaler.UnmarshalJSON.
func (fc *FeatureCollection) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
