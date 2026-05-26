// Package kml implements KML encoding.
package kml

import (
	"github.com/twpayne/go-kml/v3"

	"github.com/twpayne/go-geom"
)

// Encode encodes an arbitrary geometry.
func Encode(g geom.T) (kml.Element, error) {
	_ = "STUB: not implemented"
	return *new(kml.Element), nil
}

// EncodeLineString encodes a LineString.
func EncodeLineString(ls *geom.LineString) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodeLinearRing encodes a LinearRing.
func EncodeLinearRing(lr *geom.LinearRing) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodeMultiLineString encodes a MultiLineString.
func EncodeMultiLineString(mls *geom.MultiLineString) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodeMultiPoint encodes a MultiPoint.
func EncodeMultiPoint(mp *geom.MultiPoint) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodeMultiPolygon encodes a MultiPolygon.
func EncodeMultiPolygon(mp *geom.MultiPolygon) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodePoint encodes a Point.
func EncodePoint(p *geom.Point) kml.Element { _ = "STUB: not implemented"; return *new(kml.Element) }

// EncodePolygon encodes a Polygon.
func EncodePolygon(p *geom.Polygon) kml.Element {
	_ = "STUB: not implemented"
	return *new(kml.Element)
}

// EncodeGeometryCollection encodes a GeometryCollection.
func EncodeGeometryCollection(g *geom.GeometryCollection) (kml.Element, error) {
	_ = "STUB: not implemented"
	return *new(kml.Element), nil
}

func dim(l geom.Layout) int { _ = "STUB: not implemented"; return 0 }
