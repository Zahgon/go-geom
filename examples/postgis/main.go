package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	dsn = flag.String("dsn", "postgres://localhost/geomtest?binary_parameters=yes&sslmode=disable", "data source name")

	create   = flag.Bool("create", false, "create database schema")
	populate = flag.Bool("populate", false, "populate waypoints")
	read     = flag.Bool("read", false, "import waypoint from stdin in GeoJSON format")
	write    = flag.Bool("write", false, "write waypoints to stdout in GeoJSON format")
)

// A Waypoint is a location with an identifier and a name.
type Waypoint struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Geometry json.RawMessage `json:"geometry"`
}

// createDB demonstrates create a PostgreSQL/PostGIS database with a table with
// a geometry column.
func createDB(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

// populateDB demonstrates populating a PostgreSQL/PostGIS database using
// pq.CopyIn for fast imports.
func populateDB(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

// readGeoJSON demonstrates reading data in GeoJSON format and inserting it
// into a database in EWKB format.
func readGeoJSON(db *sql.DB, r io.Reader) error { _ = "STUB: not implemented"; return nil }

// writeGeoJSON demonstrates reading data from a database in EWKB format and
// writing it as GeoJSON.
func writeGeoJSON(db *sql.DB, w io.Writer) error { _ = "STUB: not implemented"; return nil }

func run() error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
