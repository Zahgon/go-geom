package igc

import (
	"io"

	"github.com/twpayne/go-geom"
)

// An Encoder is an IGC encoder.
type Encoder struct {
	a string
	w io.Writer
}

// An EncoderOption sets an option on an Encoder.
type EncoderOption func(*Encoder)

func clamp(x, minValue, maxValue int) int { _ = "STUB: not implemented"; return 0 }

// NewEncoder returns a new Encoder that writes to w.
func NewEncoder(w io.Writer, options ...EncoderOption) *Encoder {
	_ = "STUB: not implemented"
	return nil
}

// Encode encodes a LineString.
func (enc *Encoder) Encode(ls *geom.LineString) error { _ = "STUB: not implemented"; return nil }

// A sets the A-record text.
func A(a string) EncoderOption { _ = "STUB: not implemented"; return *new(EncoderOption) }
