package geometry

type Feature struct {
	featureType int
}

func (f *Feature) IsEdge() bool {
	return f.featureType == FEATURE_EDGE
}

func (f *Feature) IsVertex() bool {
	return f.featureType == FEATURE_VERTEX
}
