package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
	"math"
	"sort"
)

type sapBruteForceProxy struct {
	collidable collision.Collider
	aabb       *geometry.AABB
}

func (s *sapBruteForceProxy) CompareTo(o *sapBruteForceProxy) int {
	if s == o {
		return 0
	}
	diff := s.aabb.GetMinX() - o.aabb.GetMinX()
	if diff > 0 {
		return 1
	} else if diff < 0 {
		return -1
	} else {
		diff := s.aabb.GetMinY() - o.aabb.GetMinY()
		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		} else {
			if s.collidable.GetID() > o.collidable.GetID() {
				return 1
			} else {
				return -1
			}
		}
	}
}

type sapBruteForcePairList struct {
	proxy      *sapBruteForceProxy
	potentials []*sapBruteForceProxy
}

func NewsapBruteForcePairList() *sapBruteForcePairList {
	p := new(sapBruteForcePairList)
	p.potentials = make([]*sapBruteForceProxy, 0)
	return p
}

type sapBruteForceProxyList []*sapBruteForceProxy

func (p sapBruteForceProxyList) Len() int {
	return len(p)
}

func (p sapBruteForceProxyList) Less(i, j int) bool {
	return p[i].CompareTo(p[j]) < 0
}

func (p sapBruteForceProxyList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type SapBruteForce struct {
	AbstractAABBDetector
	proxyList      sapBruteForceProxyList
	proxyMap       map[string]*sapBruteForceProxy
	potentialPairs []*sapBruteForcePairList
	sort           bool
}

func NewSapBruteForce() *SapBruteForce {
	return NewSapBruteForceInt(64)
}

func NewSapBruteForceInt(initialCapacity int) *SapBruteForce {
	s := new(SapBruteForce)
	InitAbstractAABBDetector(&s.AbstractAABBDetector)
	s.sort = false
	s.proxyList = make([]*sapBruteForceProxy, 0, initialCapacity)
	s.proxyMap = make(map[string]*sapBruteForceProxy)
	s.potentialPairs = make([]*sapBruteForcePairList, 0, initialCapacity)
	return s
}

func (s *SapBruteForce) Add(collidable collision.Collider) {
	id := collidable.GetID()
	aabb := collidable.CreateAABB()
	aabb.Expand(s.expansion)
	p := new(sapBruteForceProxy)
	p.collidable = collidable
	p.aabb = aabb
	s.proxyList = append(s.proxyList, p)
	s.proxyMap[id] = p
	s.sort = true
}

func (s *SapBruteForce) Remove(collidable collision.Collider) {
	for i, p := range s.proxyList {
		if p.collidable == collidable {
			s.proxyList = append(s.proxyList[:i], s.proxyList[i+1:]...)
			break
		}
	}
	delete(s.proxyMap, collidable.GetID())
}

func (s *SapBruteForce) Update(collidable collision.Collider) {
	p0, ok := s.proxyMap[collidable.GetID()]
	if !ok {
		return
	}
	aabb := collidable.CreateAABB()
	if p0.aabb.ContainsAABB(aabb) {
		return
	} else {
		aabb.Expand(s.expansion)
	}
	p0.aabb = aabb
	s.sort = true
}

func (s *SapBruteForce) Clear() {
	s.proxyList = s.proxyList[0:0]
	for k := range s.proxyMap {
		delete(s.proxyMap, k)
	}
}

func (s *SapBruteForce) GetAABB(collidable collision.Collider) *geometry.AABB {
	proxy, ok := s.proxyMap[collidable.GetID()]
	if ok {
		return proxy.aabb
	}
	return nil
}

func (s *SapBruteForce) Detect() []*BroadphasePair {
	size := len(s.proxyList)
	if size == 0 {
		return make([]*BroadphasePair, 0)
	}
	eSize := collision.GetEstimatedCollisionPairs(size)
	pairs := make([]*BroadphasePair, 0, eSize)
	s.potentialPairs = make([]*sapBruteForcePairList, 0)
	if s.sort {
		sort.Sort(s.proxyList)
		s.sort = false
	}
	pl := NewsapBruteForcePairList()
	for i, current := range s.proxyList {
		for j := i + 1; j < size; j++ {
			test := s.proxyList[j]
			if current.aabb.GetMaxX() >= test.aabb.GetMinX() {
				pl.potentials = append(pl.potentials, test)
			} else {
				break
			}
		}
		if len(pl.potentials) > 0 {
			pl.proxy = current
			s.potentialPairs = append(s.potentialPairs, pl)
			pl = NewsapBruteForcePairList()
		}
	}
	for _, current := range s.potentialPairs {
		for _, test := range current.potentials {
			if current.proxy.aabb.Overlaps(test.aabb) {
				pair := NewBroadphasePair(current.proxy.collidable, test.collidable)
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}

func (s *SapBruteForce) DetectAABB(aabb *geometry.AABB) []collision.Collider {
	size := len(s.proxyList)
	if size == 0 {
		return make([]collision.Collider, 0)
	}
	list := make([]collision.Collider, collision.GetEstimatedCollisions())
	if s.sort {
		sort.Sort(s.proxyList)
		s.sort = false
	}
	index := size / 2
	max := size
	min := 0
	for true {
		p := s.proxyList[index]
		if p.aabb.GetMinX() < aabb.GetMinX() {
			min = index
		} else {
			max = index
		}
		if max-min == 1 {
			break
		}
		index = (min + max) / 2
	}
	for i, p := range s.proxyList {
		if p.aabb.GetMaxX() > aabb.GetMinX() {
			if p.aabb.Overlaps(aabb) {
				list = append(list, p.collidable)
			}
		} else {
			if i >= index {
				break
			}
		}
	}
	return list
}

func (s *SapBruteForce) Raycast(ray *geometry.Ray, length float64) []collision.Collider {
	if len(s.proxyList) == 0 {
		return make([]collision.Collider, 0)
	}
	st := ray.GetStart()
	d := ray.GetDirectionVector2()
	l := length
	if length <= 0.0 {
		l = math.Inf(1)
	}
	x1 := st.X
	x2 := st.X + d.X*l
	y1 := st.Y
	y2 := st.Y + d.Y*l
	min := geometry.NewVector2FromXY(math.Min(x1, x2), math.Min(y1, y2))
	max := geometry.NewVector2FromXY(math.Max(x1, x2), math.Max(y2, y2))
	aabb := geometry.NewAABBFromVector2(min, max)
	return s.DetectAABB(aabb)
}

func (s *SapBruteForce) ShiftCoordinates(shift *geometry.Vector2) {
	for _, proxy := range s.proxyList {
		proxy.aabb.Translate(shift)
	}
}
