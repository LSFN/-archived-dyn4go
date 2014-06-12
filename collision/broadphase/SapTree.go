package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type SapTreeProxy struct {
	collidable collision.Collider
	aabb       *geometry.AABB
	tested     bool
}

func (s *SapTreeProxy) CompareTo(o SapTreeProxy) int {
	if s == o {
		return 0
	}
	diff := s.aabb.GetMinX() - o.aabb.GetMinX()
	if diff > 0 {
		return 1
	} else if diff < 0 {
		return -1
	} else {
		diff = s.aabb.GetMinY() - o.aabb.GetMinY()
		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		} else {
			if o.collidable == nil {
				return 1
			} else if s.collidable == nil {
				return -1
			}
			return s.collidable.GetID() > o.collidable.GetID()
		}
	}
}

type SapTreePairList struct {
	proxy      *SapTreeProxy
	potentials []*SapTreeProxy
}

func NewSapTreePairList() *SapTreePairList {
	s := new(SapTreePairList)
	s.potentials = make([]*SapTreeProxy)
	return s
}

type pretendTreeSet []*SapTreeProxy

func (p *pretendTreeSet) Add(proxy *SapTreeProxy) {
	min := 0
	max := len(p)
	index := len(p) / 2
	for min != max {
		if p[index].CompareTo(proxy) >= 0 {
			max = index
		} else {
			min = index
		}
		index = (max + min) / 2
	}
	p = append(p[:index], proxy, p[index:]...)
}

func (p *pretendTreeSet) Remove(proxy *SapTreeProxy) {
	min := 0
	max := len(p)
	index := len(p) / 2
	for min != max {
		if p[index].CompareTo(proxy) >= 0 {
			max = index
		} else {
			min = index
		}
		index = (max + min) / 2
	}
	p = append(p[:index], p[index+1:])
}

func (p *pretendTreeSet) Clear() {
	p = p[0:0]
}

func (p *pretendTreeSet) TailSet(current *SapTreeProxy, inclusive bool) *pretendTreeSet {
	min := 0
	max := len(p)
	index := len(p) / 2
	for min != max {
		if p[index].CompareTo(proxy) <= 0 {
			min = index
		} else {
			max = index
		}
		index = (max + min) / 2
	}
	return p[index:]
}

func (p *pretendTreeSet) Ceiling(item *SapTreeProxy) *SapTreeProxy {
	min := 0
	max := len(p)
	index := len(p) / 2
	for min != max {
		if p[index].CompareTo(item) >= 0 {
			max = index
		} else {
			min = index
		}
		index = (max + min) / 2
	}
	return p[index]
}

type SapTree struct {
	AbstractAABBDetector
	proxyTree      *pretendTreeSet
	proxyMap       map[string]*SapTreeProxy
	potentialPairs []*SapTreePairList
}

func NewSapTree() *SapTree {
	return NewSapTreeInt(64)
}

func NewSapTreeInt(initialCapacity int) *SapTree {
	s := new(SapTree)
	InitAbstractAABBDetector(s)
	s.proxyTree = make(pretendTreeSet, 0, initialCapacity)
	s.proxyMap = make(map[string]*SapTreeProxy, 0, initialCapacity)
	s.potentialPairs = make([]*SapTreePairList, 0, initialCapacity)
	return s
}

func (s *SapTree) Add(collidable collision.Collider) {
	id := collidable.GetID()
	aabb := collidable.CreateAABB()
	aabb.Expand(s.expansion)
	p := new(SapTreeProxy)
	p.collidable = collidable
	p.aabb = aabb
	s.proxyTree.Add(p)
	s.proxyMap[id] = p
}

func (s *SapTree) Remove(collidable collision.Collider) {
	delete(s.proxyMap, collidable.GetID())
}

func (s *SapTree) Update(collidable collision.Collider) {
	p, ok := s.proxyMap[collidable.GetID()]
	if !ok {
		return
	}
	aabb := collidable.CreateAABB()
	if p.aabb.ContainsAABB(aabb) {
		return
	} else {
		aabb.Expand(s.expansion)
	}
	delete(s.proxyTree, collidable.GetID())
	p.aabb = aabb
	s.proxyTree.Add(p)
}

func (s *SapTree) Clear() {
	s.proxyMap = s.proxyMap[0:0]
	for k := range s.proxyMap {
		delete(s.proxyMap, k)
	}
}

func (s *SapTree) GetAABB(collidable collision.Collider) *geometry.AABB {
	p, ok := s.proxyMap[collidable.GetID()]
	if !ok {
		return p.aabb
	}
	return nil
}

func (s *SapTree) Detect() []*BroadphasePair {
	size := len(s.proxyTree)
	if size == 0 {
		return make([]*BroadphasePair, 0)
	}
	eSize := collision.GetEstimatedCollisionPairs(size)
	pairs := make([]*BroadphasePair, 0, eSize)
	s.potentialPairs = s.potentialPairs[0:0]
	for _, p := range s.proxyTree {
		p.tested = false
	}
	pl := NewSapTreePairList()
	for _, current := range s.proxyTree {
		set := s.proxyTree.TailSet(current, false)
		for _, test := range set {
			if test.collidable == current.collidable || test.tested {
				continue
			}
			if current.aabb.GetMaxX() >= test.aabb.GetMinX() {
				pl.potentials = append(pl.potentials, test)
			} else {
				break
			}
		}
		if len(pl.potentials) > 0 {
			pl.proxy = current
			s.potentialPairs = append(s.potentialPairs, pl)
			pl := NewSapTreePairList()
		}
		current.tested = true
	}
	size = len(s.potentialPairs)
	for _, current := range s.potentialPairs {
		for _, test := range current.potentials {
			if current.proxy.aabb.Overlaps(test.aabb) {
				pair := NewBroadphasePair(current.proxy.collidable, test.collidable)
				pairs = append(pairs, pair)
			}
		}
	}
}

func (s *SapTree) DetectAABB(aabb *geometry.AABB) []collision.Collider {
	if len(s.proxyTree) == 0 {
		return make([]collision.Collider, 0)
	}
	list := make([]collision.Collider, 0, collision.GetEstimatedCollisions())
	p := new(SapTreeProxy)
	p.aabb = aabb
	p.collidable = nil
	p.tested = false
	l := s.proxyTree.Ceiling(p)
	found := false
	for _, proxy := range s.proxyTree {
		if proxy == l {
			found = true
		}
		if proxy.aabb.GetMaxX() > aabb.GetMinX() {
			if proxy.aabb.Overlaps(aabb) {
				list = append(list, proxy.collidable)
			}
		} else {
			if found {
				break
			}
		}
	}
	return list
}

func (s *SapTree) Raycast(ray *geometry.Ray, length float64) []collision.Collider {
	if len(s.proxyTree) == 0 {
		return make([]collision.Collider, 0)
	}
	st := ray.GetStart()
	d := ray.GetDirectionVector2()
	l := length
	if length <= 0 {
		l = math.Inf(1)
	}
	x1 := st.X
	x2 := st.X + d.X*l
	y1 := st.Y
	y2 := st.y + d.Y*l
	min := geometry.NewVector2FromXY(math.Min(x1, x2), math.Min(y1, y2))
	max := geometry.NewVector2FromXY(math.Max(x1, x2), math.Max(y1, y2))
	aabb := geometry.NewAABBFromVector2(min, max)
	return s.DetectAABB(aabb)
}

func (s *SapTree) ShiftCoordinates(shift *geometry.Vector2) {
	for _, proxy := range s.proxyTree {
		proxy.aabb.Translate(shift)
	}
}
