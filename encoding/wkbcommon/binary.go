package wkbcommon

import (
	"encoding/binary"
	"io"
)

func readFloat(buf []byte, byteOrder binary.ByteOrder) float64 { _ = "STUB: not implemented"; return 0 }

// ReadUInt32 reads a uint32 from r.
func ReadUInt32(r io.Reader, byteOrder binary.ByteOrder) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFloatArray reads a []float64 from r.
func ReadFloatArray(r io.Reader, byteOrder binary.ByteOrder, array []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert to an array of floats

// ReadByte reads a byte from r.
func ReadByte(r io.Reader) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func writeFloat(buf []byte, byteOrder binary.ByteOrder, value float64) {
	_ = "STUB: not implemented"
	return
}

// WriteFloatArray writes a []float64 to w.
func WriteFloatArray(w io.Writer, byteOrder binary.ByteOrder, array []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUInt32 writes a uint32 to w.
func WriteUInt32(w io.Writer, byteOrder binary.ByteOrder, value uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteByte wrties a byte to w.
func WriteByte(w io.Writer, value byte) error { _ = "STUB: not implemented"; return nil }

// WriteEmptyPointAsNaN outputs EmptyPoint as NaN values.
func WriteEmptyPointAsNaN(w io.Writer, byteOrder binary.ByteOrder, numCoords int) error {
	_ = "STUB: not implemented"
	return nil
}
