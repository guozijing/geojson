package geoJson

type StructGeo struct {
	Type string `json:"type"`
	Features []*StructFeature `json:"features"`
}

type StructFeature struct {
	Type string `json:"type"`
	Geometry StructGeometry `json:"geometry"`
}

type StructGeometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type StructPoint struct {
	lat float64
	lng float64
}
