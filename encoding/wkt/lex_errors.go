package wkt

// SyntaxError is an error that occurs during parsing of a WKT string.
type SyntaxError struct {
	wkt       string
	problem   string
	lineNum   int
	lineStart int
	linePos   int
	hint      string
}

// Error generates a detailed syntax error message with line and pos numbers as well as a snippet of
// the erroneous input.
func (e *SyntaxError) Error() string {
	_ = "STUB: not implemented"
	// These constants define the maximum number of characters of the line to show on each side of the cursor.
	return ""
}

// Print the problem along with line and pos number.

// Find the position of the end of the line.

// Prepend the line with the line number.

// Trim the start and end of the line as needed.

// Print a cursor pointing to the token where the problem occurred.

// Print a hint, if applicable.
