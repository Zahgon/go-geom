package wkt

import (
	"github.com/twpayne/go-geom"
)

// Constant expected by parser when lexer reaches EOF.
const eof = 0

// We define a base type geometry as a geometry type keyword without a type suffix.
// For example, POINT is a base type and POINTZ is not.
//
// The layout of the geometry is determined by the first geometry type keyword if it is a M, Z, or ZM variant.
// If it is a base type geometry, the layout is determined by the number of coordinates in the first point.
// If it is a geometrycollection, the type is the type of the first geometry in the collection.
//
// Edge cases involving geometrycollections:
// 1. GEOMETRYCOLLECTION (no type suffix) is allowed to be of type M. Normally a geometry without a type suffix
//    is only allowed to be XY, XYZ, or XYZM.
// 2. A base type empty geometry (e.g. POINT EMPTY) in a GEOMETRYCOLLECTIONM, GEOMETRYCOLLECTIONZ, GEOMETRYCOLLECTIONZM
//    is permitted and takes on the type of the collection. Normally, such a geometry is XY.
// 3. As a consequence of 1. and 2., special care must be given to parsing base geometry types inside a XYM
//    geometrycollection since a base geometry type is permitted inside a GEOMETRYCOLECTIONM only if it is empty.
//    For example, GEOMETRYCOLLECTION M (POINT EMPTY) should parse while GEOMETRYCOLLECTION M (POINT(0 0 0)) shouldn't.

// lexPos is a struct for keeping track of both the actual and human-readable lexed position in the string.
type lexPos struct {
	wktPos    int
	lineNum   int
	lineStart int
	linePos   int
}

// advanceOne advances a lexPos by one position on the same line.
func (lp *lexPos) advanceOne() { _ = "STUB: not implemented"; return }

// advanceLine advances a lexPos by a newline.
func (lp *lexPos) advanceLine() { _ = "STUB: not implemented"; return }

// wktLex is the lexer for lexing WKT tokens.
type wktLex struct {
	wkt      string
	curPos   lexPos
	lastPos  lexPos
	ret      geom.T
	lytStack layoutStack
	lastErr  error
}

// newWKTLex returns a pointer to a newly created wktLex.
func newWKTLex(wkt string) *wktLex { _ = "STUB: not implemented"; return nil }

// Lex lexes a token from the input.
func (l *wktLex) Lex(yylval *wktSymType) int {
	_ = "STUB: not implemented"
	// Skip leading spaces.
	return 0
}

// Lex a token.

// keyword lexes a string keyword.
func (l *wktLex) keyword() int { _ = "STUB: not implemented"; return 0 }

// Add the uppercase letter to the string builder.

// Check for extra dimensions for geometry types.

// num lexes a number.
func (l *wktLex) num(yylval *wktSymType) int { _ = "STUB: not implemented"; return 0 }

// peek returns the next rune to be read.
func (l *wktLex) peek() rune { _ = "STUB: not implemented"; return 0 }

// next returns the next rune to be read and advances the curPos counter.
func (l *wktLex) next() rune { _ = "STUB: not implemented"; return 0 }

// trimLeft increments the curPos counter until the next rune to be read is no longer a whitespace character.
func (l *wktLex) trimLeft() { _ = "STUB: not implemented"; return }

// validateStrideAndSetDefaultLayoutIfNoLayout validates whether a stride is consistent with the currently parsed
// layout and sets the layout with the default layout for that stride if no layout has been determined yet.
func (l *wktLex) validateStrideAndSetDefaultLayoutIfNoLayout(stride int) bool {
	_ = "STUB: not implemented"
	return false
}

// validateNonEmptyGeometryAllowed validates whether a non-empty geometry is allowed given the currently
// parsed layout. It is used to handle the edge case where a GEOMETRYCOLLECTIONM may have base type
// geometries only if they are empty.
func (l *wktLex) validateNonEmptyGeometryAllowed() bool { _ = "STUB: not implemented"; return false }

// validateAndSetLayoutIfNoLayout validates whether a newly parsed layout is compatible with the currently parsed
// layout and sets the layout if the current layout is unknown.
func (l *wktLex) validateAndSetLayoutIfNoLayout(layout geom.Layout) bool {
	_ = "STUB: not implemented"
	return false
}

// validateBaseGeometryTypeAllowed validates whether a base geometry type is permitted based on the parsed layout.
func (l *wktLex) validateBaseGeometryTypeAllowed() bool {
	_ = "STUB: not implemented"
	// Base type geometry are permitted in GEOMETRYCOLLECTIONM, GEOMETRYCOLLECTIONZ, GEOMETRYCOLLECTIONZM.
	// The stride of the coordinates/whether EMPTY is allowed will be validated later.
	return false
}

// A base type is only permitted in a GEOMETRYCOLLECTIONM if it is EMPTY. We require an EMPTY instead of
// coordinates follow this base type keyword.

// At the top level, a base geometry type is permitted. In a base type GEOMETRYCOLLECTION, a base type geometry
// is only not permitted if the parsed layout is XYM.

// validateBaseTypeEmptyAllowed validates whether a base type EMPTY is permitted based on the parsed layout.
func (l *wktLex) validateBaseTypeEmptyAllowed() bool {
	_ = "STUB: not implemented"
	// EMPTY is always permitted in a non-base type collection.
	return false
}

// A base type EMPTY geometry is the only permitted base type geometry in a GEOMETRYCOLLECTIONM
// and we have now finished reading one.

// In a base type collection (or at the top level), EMPTY can only be XY.

// validateAndPushLayoutStackFrame validates that a given layout is valid and pushes a frame to the layout stack.
func (l *wktLex) validateAndPushLayoutStackFrame(layout geom.Layout) bool {
	_ = "STUB: not implemented"
	// Check that the new layout is compatible with the previous one.
	// Note a base type GEOMETRYCOLLECTION is permitted inside every layout.
	return false
}

// validateAndPopLayoutStackFrame pops a frame from the layout stack and validates that the type is valid.
func (l *wktLex) validateAndPopLayoutStackFrame() bool { _ = "STUB: not implemented"; return false }

// Update the outer context with the type we parsed in the inner context.

// This should never happen. Any layout incompatibility should error at the point it's discovered.

// validateLayoutStackAtEnd returns whether the layout stack is in the expected state at the end of parsing.
func (l *wktLex) validateLayoutStackAtEnd() bool { _ = "STUB: not implemented"; return false }

func (l *wktLex) isValidPoint(flatCoords []float64) bool { _ = "STUB: not implemented"; return false }

func (l *wktLex) isValidLineString(flatCoords []float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *wktLex) isValidPolygonRing(flatCoords []float64) bool {
	_ = "STUB: not implemented"
	return false
}

// setLayoutIfNoLayout sets the parsed layout if no layout has been determined yet.
func (l *wktLex) setLayoutIfNoLayout(layout geom.Layout) { _ = "STUB: not implemented"; return }

// setIncorrectUsageOfBaseTypeInsteadOfMVariantInGeometryCollectionError sets the error when a
// base type geometry is used in a base type GEOMETRYCOLLECTION when the parsed layout is XYM.
func (l *wktLex) setIncorrectUsageOfBaseTypeInsteadOfMVariantInGeometryCollectionError() {
	_ = "STUB: not implemented"
	return
}

// setIncorrectStrideError sets the error when a newly parsed stride doesn't match the currently parsed layout.
func (l *wktLex) setIncorrectStrideError(incorrectStride int, hint string) {
	_ = "STUB: not implemented"
	return
}

// setIncorrectLayoutError sets the error when a newly parsed layout doesn't match the currently parsed layout.
func (l *wktLex) setIncorrectLayoutError(incorrectLayout geom.Layout, hint string) {
	_ = "STUB: not implemented"
	return
}

// curLayout returns the currently parsed layout.
func (l *wktLex) curLayout() geom.Layout { _ = "STUB: not implemented"; return *new(geom.Layout) }

// currentlyInBaseTypeCollection returns whether we are currently scanning inside a base type GEOMETRYCOLLECTION.
func (l *wktLex) currentlyInBaseTypeCollection() bool { _ = "STUB: not implemented"; return false }

// nextScannedPointMustBeEmpty returns whether the next scanned point must be empty.
func (l *wktLex) nextScannedPointMustBeEmpty() bool { _ = "STUB: not implemented"; return false }

// setLexError is called by Lex when a lexing (tokenizing) error is detected.
func (l *wktLex) setLexError(expectedTokType string) { _ = "STUB: not implemented"; return }

// setParseError is called when a context-sensitive error is detected during parsing.
// The generated wktParse function can only catch context-free errors.
func (l *wktLex) setParseError(problem, hint string) { _ = "STUB: not implemented"; return }

// Error is called by wktParse if an error is encountered during parsing (takes place after lexing).
func (l *wktLex) Error(s string) { _ = "STUB: not implemented"; return }

// setSyntaxError is called when a syntax error occurs.
func (l *wktLex) setSyntaxError(problem, hint string) { _ = "STUB: not implemented"; return }

// setError sets the lastErr field of the wktLex object with the given error.
func (l *wktLex) setError(err error) {
	_ = "STUB: not implemented"
	// Lex errors take precedence.
	return
}

// isValidFirstNumRune returns whether a rune is valid as the first rune in a number (coordinate).
func isValidFirstNumRune(r rune) bool {
	_ = "STUB: not implemented"

	// PostGIS doesn't seem to accept numbers with a leading '+'.
	return false
}

// Scientific notation number must have a number before the e.
// Checking this case explicitly helps disambiguate between a number and a keyword.

// isNumRune returns whether a rune could potentially be a part of a number (coordinate).
func isNumRune(r rune) bool { _ = "STUB: not implemented"; return false }

// keywordsMap defines a map from strings to tokens.
var keywordsMap = map[string]int{
	"EMPTY": EMPTY,
	"POINT": POINT, "POINTM": POINTM, "POINTZ": POINTZ, "POINTZM": POINTZM,
	"LINESTRING": LINESTRING, "LINESTRINGM": LINESTRINGM, "LINESTRINGZ": LINESTRINGZ, "LINESTRINGZM": LINESTRINGZM,
	"POLYGON": POLYGON, "POLYGONM": POLYGONM, "POLYGONZ": POLYGONZ, "POLYGONZM": POLYGONZM,
	"MULTIPOINT": MULTIPOINT, "MULTIPOINTM": MULTIPOINTM, "MULTIPOINTZ": MULTIPOINTZ, "MULTIPOINTZM": MULTIPOINTZM,
	"MULTILINESTRING": MULTILINESTRING, "MULTILINESTRINGM": MULTILINESTRINGM,
	"MULTILINESTRINGZ": MULTILINESTRINGZ, "MULTILINESTRINGZM": MULTILINESTRINGZM,
	"MULTIPOLYGON": MULTIPOLYGON, "MULTIPOLYGONM": MULTIPOLYGONM,
	"MULTIPOLYGONZ": MULTIPOLYGONZ, "MULTIPOLYGONZM": MULTIPOLYGONZM,
	"GEOMETRYCOLLECTION": GEOMETRYCOLLECTION, "GEOMETRYCOLLECTIONM": GEOMETRYCOLLECTIONM,
	"GEOMETRYCOLLECTIONZ": GEOMETRYCOLLECTIONZ, "GEOMETRYCOLLECTIONZM": GEOMETRYCOLLECTIONZM,
}

// keywordToken returns the yacc token for a WKT keyword.
func keywordToken(tokStr string) int { _ = "STUB: not implemented"; return 0 }

// isValidStrideForLayout returns whether a stride is consistent with a parsed layout.
// It is used for ensuring points have the right number of coordinates for the parsed layout.
func isValidStrideForLayout(stride int, layout geom.Layout) bool {
	_ = "STUB: not implemented"
	return false
}

// This should never happen.

// defaultLayoutForStride returns the default layout for a base type geometry with the given stride.
func defaultLayoutForStride(stride int) geom.Layout {
	_ = "STUB: not implemented"
	return *new(geom.Layout)
}

// This should never happen.

// isCompatibleLayout returns whether a second layout is compatible with the first layout.
// It is used for ensuring the layout of each nested geometry is consistent with the previously parsed layout.
func isCompatibleLayout(outerLayout, innerLayout geom.Layout) bool {
	_ = "STUB: not implemented"
	return false
}

// layoutName returns the string representation of each layout.
func layoutName(layout geom.Layout) string {
	_ = "STUB: not implemented"

	// geom.NoLayout is used when a base type geometry is read.
	return ""
}

// This should never happen.

// assertValidLayout asserts that a given layout is valid and panics if it is not.
func assertValidLayout(layout geom.Layout) { _ = "STUB: not implemented"; return }
