package geoJson

type StructGeo struct {
	Type string `json:"type"`
	Features []*StructFeature `json:"features"`
}

type StructFeature struct {
	Type string `json:"type"`
	Properties StructProperties `json:"properties"`
	Geometry StructGeometry `json:"geometry"`
}

type StructProperties struct {
	Index int `json:"index"`
}

type StructGeometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type StructPoint struct {
	lat float64
	lng float64
}
