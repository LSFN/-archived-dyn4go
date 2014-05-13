package manifold

type IndexedManifoldPointId struct {
	referenceEdge, incidentEdge, incidentVertex int
	flipped                                     bool
}

func NewIndexedManifoldPointID(referenceEdge, incidentEdge, incidentVertex int) *IndexedManifoldPointId {
	return &IndexedManifoldPointId{referenceEdge, incidentEdge, incidentVertex, false}
}

func NewIndexedManifoldPointIDBool(referenceEdge, incidentEdge, incidentVertex int, flipped bool) *IndexedManifoldPointId {
	return &IndexedManifoldPointId{referenceEdge, incidentEdge, incidentVertex, flipped}
}

func (i *IndexedManifoldPointId) GetReferenceEdge() int {
	return i.referenceEdge
}

func (i *IndexedManifoldPointId) GetIncidentEdge() int {
	return i.incidentEdge
}

func (i *IndexedManifoldPointId) GetIncidentVertex() int {
	return i.incidentVertex
}

func (i *IndexedManifoldPointId) IsFlipped() bool {
	return i.flipped
}
