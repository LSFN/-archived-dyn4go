package manifold

import (
	"github.com/LSFN/dyn4go/collision/narrowphase"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type ClippingManifoldSolver struct{}

func (c *ClippingManifoldSolver) GetManifold(penetration *narrowphase.Penetration, convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, manifold *Manifold) bool {
	manifold.Clear()
	n := penetration.GetNormal()
	feature1 := convex1.GetFarthestFeature(n, transform1)
	if feature1.IsVertex() {
		vertex := feature1.(*geometry.Vertex)
		mp := NewManifoldPointInterfaceVector2Float64(DISTANCE, vertex.GetPoint(), penetration.GetDepth())
		manifold.points = append(manifold.points, mp)
		manifold.normal = n.Negate()
		return true
	}
	feature2 := convex2.GetFarthestFeature(n, transform2)
	if feature2.IsVertex() {
		vertex := feature2.(*geometry.Vertex)
		mp := NewManifoldPointInterfaceVector2Float64(DISTANCE, vertex.GetPoint(), penetration.GetDepth())
		manifold.points = append(manifold.points, mp)
		manifold.normal = n.Negate()
		return true
	}
	reference := feature1.(*geometry.Edge)
	incident := feature2.(*geometry.Edge)
	flipped := false
	e1 := reference.GetEdge()
	e2 := incident.GetEdge()
	if math.Abs(e1.DotVector2(n)) > math.Abs(e2.DotVector2(n)) {
		e := reference
		reference = incident
		incident = e
		flipped = true
	}
	refev := reference.GetEdge()
	refev.Normalize()
	offset1 := -refev.DotVector2(reference.GetVertex1().GetPoint())
	offset2 := refev.DotVector2(reference.GetVertex2().GetPoint())
	clip1 := c.clip(incident.GetVertex1(), incident.GetVertex2(), refev.GetNegative(), offset1)
	if len(clip1) < 2 {
		return false
	}
	clip2 := c.clip(clip1[0], clip1[1], refev, offset2)
	if len(clip2) < 2 {
		return false
	}
	frontNormal := refev.CrossZ(1)
	frontOffset := frontNormal.DotVector2(reference.GetMaximum().GetPoint())
	manifold.normal = frontNormal
	if flipped {
		manifold.normal = frontNormal.GetNegative()
	}
	for _, vertex := range clip2 {
		point := vertex.GetPoint()
		depth := frontNormal.DotVector2(point) - frontOffset
		if depth >= 0 {
			id := NewIndexedManifoldPointIDBool(reference.GetIndex(), incident.GetIndex(), vertex.GetIndex(), flipped)
			mp := NewManifoldPointInterfaceVector2Float64(id, point, depth)
			manifold.points = append(manifold.points, mp)
		}
	}
	return len(manifold.points) != 0
}

func (c *ClippingManifoldSolver) clip(v1, v2 *geometry.Vertex, n *geometry.Vector2, offset float64) []*geometry.Vertex {
	points := make([]*geometry.Vertex, 2)
	p1 := v1.GetPoint()
	p2 := v2.GetPoint()
	d1 := n.DotVector2(p1) - offset
	d2 := n.DotVector2(p2) - offset
	if d1 <= 0 {
		points = append(points, v1)
	}
	if d2 <= 0 {
		points = append(points, v2)
	}
	if d1*d2 < 0 {
		e := p1.HereToVector2(p2)
		u := d1 / (d1 - d2)
		e.Multiply(u)
		e.AddVector2(p1)
		if d1 > 0 {
			points = append(points, geometry.NewVertexVector2Int(e, v1.GetIndex()))
		} else {
			points = append(points, geometry.NewVertexVector2Int(e, v2.GetIndex()))
		}
	}
	return points
}
