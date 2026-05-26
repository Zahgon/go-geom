package wkt

type geomFlatCoordsRepr struct {
	flatCoords []float64
	ends       []int
}

func makeGeomFlatCoordsRepr(flatCoords []float64) geomFlatCoordsRepr {
	_ = "STUB: not implemented"
	return *new(geomFlatCoordsRepr)
}

func appendGeomFlatCoordsReprs(p1, p2 geomFlatCoordsRepr) geomFlatCoordsRepr {
	_ = "STUB: not implemented"
	return *new(geomFlatCoordsRepr)
}

type multiPolygonFlatCoordsRepr struct {
	flatCoords []float64
	endss      [][]int
}

func makeMultiPolygonFlatCoordsRepr(p geomFlatCoordsRepr) multiPolygonFlatCoordsRepr {
	_ = "STUB: not implemented"
	return *new(multiPolygonFlatCoordsRepr)
}

func appendMultiPolygonFlatCoordsRepr(
	p1, p2 multiPolygonFlatCoordsRepr,
) multiPolygonFlatCoordsRepr {
	_ = "STUB: not implemented"
	return *new(multiPolygonFlatCoordsRepr)
}
