// parse-igc parsers all IGC files passed on the command line and prints any
// errors encountered.
package main

import (
	"fmt"
	"os"

	"github.com/twpayne/go-geom/encoding/igc"
)

func parseIGC(filename string) (*igc.T, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:unparam
func run() error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
