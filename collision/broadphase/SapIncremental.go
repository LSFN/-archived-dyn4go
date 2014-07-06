package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type sapIncrementalProxy struct {
	collidable collision.Collider
	aabb       *geometry.AABB
}

func (s *sapIncrementalProxy) CompareTo(o *sapIncrementalProxy) int {
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
			if s.collidable.GetID() > o.collidable.GetID() {
				return 1
			} else {
				return -1
			}
		}
	}
}

type sapIncrementalProxyList []*sapIncrementalProxy

func (s sapIncrementalProxyList) Len() int {
	return len(s)
}

func (s sapIncrementalProxyList) Less(i, j int) bool {
	return s[i].CompareTo(s[j]) < 0
}

func (s sapIncrementalProxyList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sapIncrementalProxyList) Search(item *sapIncrementalProxy) (int, bool) {
	for i, v := range s {
		compare := v.CompareTo(item)
		if compare == 0 {
			return i, true
		} else if compare > 0 {
			return i, false
		}
	}
	return 0, false
}

type sapIncrementalPairList struct {
	proxy      *sapIncrementalProxy
	potentials []*sapIncrementalProxy
}

func NewSapIncrementalPairList() *sapIncrementalPairList {
	p := new(sapIncrementalPairList)
	p.potentials = make([]*sapIncrementalProxy, 0, collision.GetEstimatedCollisions())
	return p
}

type SapIncremental struct {
	AbstractAABBDetector
	proxyList      sapIncrementalProxyList
	proxyMap       map[string]*sapIncrementalProxy
	potentialPairs []*sapIncrementalPairList
}

func NewSapIncremental() *SapIncremental {
	return NewSapIncrementalInt(64)
}

func NewSapIncrementalInt(initialCapacity int) *SapIncremental {
	s := new(SapIncremental)
	InitAbstractAABBDetector(&s.AbstractAABBDetector)
	s.proxyList = make([]*sapIncrementalProxy, 0, initialCapacity)
	s.proxyMap = make(map[string]*sapIncrementalProxy)
	s.potentialPairs = make([]*sapIncrementalPairList, 0, initialCapacity)
	return s
}

func (s *SapIncremental) Add(collidable collision.Collider) {
	id := collidable.GetID()
	aabb := collidable.CreateAABB()
	aabb.Expand(s.expansion)
	p := new(sapIncrementalProxy)
	p.collidable = collidable
	p.aabb = aabb
	index, _ := s.proxyList.Search(p)
	temp := s.proxyList[index:]
	s.proxyList = append(s.proxyList[:index], p)
	s.proxyList = append(s.proxyList, temp...)
	s.proxyMap[id] = p
}

func (s *SapIncremental) Remove(collidable collision.Collider) {
	index := -1
	for i, p := range s.proxyList {
		if p.collidable == collidable {
			index = i
			break
		}
	}
	if index != -1 {
		s.proxyList = append(s.proxyList[:index], s.proxyList[index+1:]...)
	}
	delete(s.proxyMap, collidable.GetID())
}

func (s *SapIncremental) Update(collidable collision.Collider) {
	p0 := s.proxyMap[collidable.GetID()]
	if p0 == nil {
		return
	}
	aabb := collidable.CreateAABB()
	if p0.aabb.ContainsAABB(aabb) {
		return
	} else {
		aabb.Expand(s.expansion)
	}
	index := -1
	for i, p := range s.proxyList {
		if p == p0 {
			index = i
			break
		}
	}
	if index != -1 {
		s.proxyList = append(s.proxyList[:index], s.proxyList[index+1:]...)
	}
	p0.aabb = aabb
	index, ok := s.proxyList.Search(p0)
	if ok {
		s.proxyList = append(s.proxyList, p0)
	}
}

func (s *SapIncremental) Clear() {
	s.proxyList = s.proxyList[0:0]
	for k := range s.proxyMap {
		delete(s.proxyMap, k)
	}
}

func (s *SapIncremental) GetAABB(collidable collision.Collider) *geometry.AABB {
	if proxy, ok := s.proxyMap[collidable.GetID()]; ok {
		return proxy.aabb
	}
	return nil
}

func (s *SapIncremental) Detect() []*BroadphasePair {
	size := len(s.proxyList)
	if size == 0 {
		return make([]*BroadphasePair, 0)
	}
	eSize := collision.GetEstimatedCollisionPairs(size)
	pairs := make([]*BroadphasePair, 0, eSize)
	s.potentialPairs = s.potentialPairs[0:0]

	pl := NewSapIncrementalPairList()
	for i := 0; i < size; i++ {
		current := s.proxyList[i]
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
			pl = NewSapIncrementalPairList()
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

func (s *SapIncremental) DetectAABB(aabb *geometry.AABB) []collision.Collider {
	size := len(s.proxyList)
	if size == 0 {
		return make([]collision.Collider, 0)
	}
	list := make([]collision.Collider, 0, collision.GetEstimatedCollisions())
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

func (s *SapIncremental) Raycast(ray *geometry.Ray, length float64) []collision.Collider {
	if len(s.proxyList) == 0 {
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
	y2 := st.Y + d.Y*l
	min := geometry.NewVector2FromXY(math.Min(x1, x2), math.Min(y1, y2))
	max := geometry.NewVector2FromXY(math.Max(x1, x2), math.Max(y1, y2))
	aabb := geometry.NewAABBFromVector2(min, max)
	return s.DetectAABB(aabb)
}

func (s *SapIncremental) ShiftCoordinates(shift *geometry.Vector2) {
	for _, proxy := range s.proxyList {
		proxy.aabb.Translate(shift)
	}
}
