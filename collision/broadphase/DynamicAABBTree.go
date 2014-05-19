package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type DATNode struct {
	left, right, parent *DATNode
	height              int
	collidable          collision.Collider
	aabb                *geometry.AABB
	tested              bool
}

func (d *DATNode) IsLeaf() bool {
	return d.left == nil
}

type DynamicAABBTree struct {
	AbstractAABBDetector
	root      *DATNode
	proxyList []*DATNode
	proxyMap  map[string]*DATNode
}

func NewDynamicAABBTree() *DynamicAABBTree {
	return NewDynamicAABBTreeInt(64)
}

func NewDynamicAABBTreeInt(initialCapacity int) *DynamicAABBTree {
	d := new(DynamicAABBTree)
	d.proxyList = make([]*DATNode, initialCapacity)
	d.proxyMap = make(map[string]*DATNode)
	InitAbstractAABBDetector(&d.AbstractAABBDetector)
	return d
}

func (d *DynamicAABBTree) Add(collidable collision.Collider) {
	aabb := collidable.CreateAABB()
	aabb.Expand(d.expansion)
	node := new(DATNode)
	node.collidable = collidable
	node.aabb = aabb
	d.proxyList = append(d.proxyList, node)
	d.proxyMap[collidable.GetID()] = node
	d.insert(node)
}

func (d *DynamicAABBTree) Remove(collidable collision.Collider) {
	node, ok := d.proxyMap[collidable.GetID()]
	if ok {
		delete(d, collidable.GetID())
		found := false
		for i, v := range d.proxyList {
			if v.collidable == collidable {
				found = true
				break
			}
		}
		if found {
			d.proxyList = append(d.proxyList[:i], d.proxyList[i+1:]...)
		}
		d.remove(collidable)
	}
}

func (d *DynamicAABBTree) Update(collidable collision.Collider) {
	node, ok := d.proxyMap[collidable.GetID()]
	if ok {
		aabb := collidable.CreateAABB()
		if node.aabb.ContainsAABB(aabb) {
			return
		}
		aabb.Expand(d.expansion)
		d.remove(node)
		node.aabb = aabb
		d.insert(node)
	}
}

func (d *DynamicAABBTree) Clear() {
	d.proxyList = d.proxyList[0:0]
	d.proxyMap = make(map[string]*DATNode)
	d.root = nil
}

func (d *DynamicAABBTree) GetAABB(collidable collision.Collider) *geometry.AABB {
	node, ok := d.proxyMap[collidable.GetID()]
	if ok {
		return node.aabb
	}
	return nil
}

func (d *DynamicAABBTree) Detect() []*BroadphasePair {
	size := len(d.proxyList)
	if size == 0 {
		return []*BroadphasePair{}
	}
	for _, node := range d.proxyList {
		node.tested = false
	}
	eSize := collision.GetEstimatedCollisionPairs(size)
	pairs := make([]*BroadphasePair, eSize)
	for _, node := range d.proxyList {
		detectNonRecursive(node, d.root, pairs)
		node.tested = true
	}
	return pairs
}

func (d *DynamicAABBTree) DetectAABB(aabb *geometry.AABB) []collision.Collider {
	return d.detectNonRecursive(aabb, d.root)
}

func (dyn *DynamicAABBTree) Raycast(ray *geometry.Ray, length float64) []collision.Collider {
	if len(d.proxyList) == 0 {
		return []collision.Collider{}
	}
	s := ray.GetStart()
	d := ray.GetDirectionVector2()
	l := length
	if l <= 0 {
		l = math.Inf(1)
	}
	x1 := s.X
	x2 := s.x + d.X*l
	y1 := s.Y
	y2 := s.Y + d.Y*l
	min := geometry.NewVector2FromXY(math.Min(x1, x2), math.Min(y1, y2))
	max := geometry.NewVector2FromXY(math.Max(x1, x2), math.Max(y1, y2))
	aabb := geometry.NewAABBFromVector2(min, max)
	return dyn.DetectAABB(aabb)
}

func (d *DynamicAABBTree) ShiftCoordinates(shift *geometry.Vector2) {
	node := d.root
	for node != nil {
		if node.left != nil {
			node = node.left
		} else if node.right != nil {
			node.aabb.Translate(shift)
			node = node.right
		} else {
			node.aabb.Translate(shift)
			nextNodeFound := false
			for node.parent != nil {
				if node == node.parent.left {
					if node.parent.right != nil {
						node.parent.aabb.Translate(shift)
						node = node.parent.right
						nextNodeFound = true
						break
					}
				}
				node = node.parent
			}
			if !nextNodeFound {
				break
			}
		}
	}
}

func (d *DynamicAABBTree) detect(node, root *DATNode, pairs *[]*BroadphasePair) {
	if root == nil || root.tested || node.collidable == root.collidable {
		return
	}
	if node.aabb.Overlaps(root.aabb) {
		if root.left == nil {
			pair := NewBroadphasePair(node.collidable, root.collidable)
			*pairs = append(*pairs, pair)
			return
		}
		if root.left != nil {
			d.detect(node, root.left, pairs)
		}
		if root.right != nil {
			d.detect(node, root.right, pairs)
		}
	}
}

func (d *DynamicAABBTree) detectNonRecursive(node, root *DATNode, pairs *[]*BroadphasePair) {
	n := root
	for n != nil {
		if n.aabb.Overlaps(node.aabb) {
			if n.left != nil {
				n = n.left
				continue
			} else {
				if !n.tested && n.collidable != node.collidable {
					pair := NewBroadphasePair(node.collidable, n.collidable)
					*pairs = append(*pairs, pair)
				}
			}
		}
		nextNodeFound := false
		for n.parent != nil {
			if n == n.parent.left {
				if n.parent.right != nil {
					n = n.parent.right
					nextNodeFound = true
					break
				}
			}
			n = n.parent
		}
		if !nextNodeFound {
			break
		}
	}
}

func (d *DynamicAABBTree) detectAABB(aabb *geometry.AABB, node *DATNode, list *[]collision.Collider) {
	if aabb.Overlaps(node.aabb) {
		if node.left == nil {
			*list = append(list, node.collidable)
			return
		}
		if node.left != nil {
			d.detectAABB(aabb, node.left, list)
		}
		if node.right != nil {
			d.detectAABB(aabb, node.right, list)
		}
	}
}

func (d *DynamicAABBTree) detectNonRecursiveAABB(aabb *geometry.AABB, node *DATNode) *[]collision.Collider {
	eSize := collision.GetEstimatedCollisions()
	list := make([]collision.Collider, eSize)
	for node != nil {
		if aabb.Overlaps(node.aabb) {
			if node.left != nil {
				node = node.left
				continue
			} else {
				list = append(list, node.collidable)
			}
		}
		nextNodeFound := false
		for node.parent != nil {
			if node == node.parent.left {
				if node.parent.right != nil {
					node = node.parent.right
					nextNodeFound = true
					break
				}
			}
			node = node.parent
		}
		if !nextNodeFound {
			break
		}
	}
	return list
}

func (d *DynamicAABBTree) insert(item *DATNode) {
	if d.root == nil {
		d.root = item
		return
	}
	itemAABB := item.aabb
	node := d.root
	for !node.IsLeaf() {
		aabb := node.aabb
		perimeter := aabb.GetPerimeter()
		union := aabb.GetUnion(itemAABB)
		unionPerimeter := union.GetPerimeter()
		cost := 2 * unionPerimeter
		descendCost := 2 * (unionPerimeter - perimeter)
		left := node.left
		right := node.right
		costl := 0.0
		if left.IsLeaf() {
			u := left.aabb.GetUnion(itemAABB)
			costl = u.GetPerimeter() + descendCost
		} else {
			u := left.aabb.GetUnion(itemAABB)
			oldPerimeter := left.aabb.GetPerimeter()
			newPerimeter := u.GetPerimeter()
			costl = newPerimeter - oldPerimeter + descendCost
		}
		costr := 0.0
		if left.IsLeaf() {
			u := left.aabb.GetUnion(itemAABB)
			costr = u.GetPerimeter() + descendCost
		} else {
			u := right.aabb.GetUnion(itemAABB)
			oldPerimeter := right.aabb.GetPerimeter()
			newPerimeter := u.GetPerimeter()
			costr = newPerimeter - oldPerimeter + descendCost
		}
		if cost < costl && cost < costr {
			break
		}
		if costl < costr {
			node = left
		} else {
			node = right
		}
	}
	parent := node.parent
	newParent := new(DATNode)
	newParent.parent = node.parent
	newParent.aabb = node.aabb.GetUnion(itemAABB)
	newParent.height = node.height + 1
	if parent != nil {
		if parent.left == node {
			parent.left = newParent
		} else {
			parent.right = newParent
		}
		newParent.left = node
		newParent.right = item
		node.parent = newParent
		item.parent = newParent
	} else {
		newParent.left = node
		newParent.right = item
		node.parent = newParent
		item.parent = newParent
		d.root = newParent
	}
	node = item.parent
	for node != nil {
		node = d.balance(node)
		left := node.left
		right := node.right
		node.height = 1 + math.Max(left.height, right.height)
		node.aabb = left.aabb.GetUnion(right.aabb)
		node = node.parent
	}
}

func (d *DynamicAABBTree) remove(node *DATNode) {
	if d.root == nil {
		return
	}
	if node == d.root {
		d.root = nil
		return
	}
	parent := node.parent
	grandparent := parent.parent
	var other *DATNode
	if parent.left == node {
		other = parent.right
	} else {
		other = parent.left
	}
	if grandparent != nil {
		if grandparent.left == parent {
			grandparent.left = other
		} else {
			grandparent.right = other
		}
		other.parent = grandparent
		n := grandparent
		for n != nil {
			n := d.balance(n)
			left := n.left
			right := n.right
			n.height = 1 + math.Max(left.height, right.height)
			n.aabb = left.getUnion(right.aabb)
			n = n.parent
		}
	} else {
		d.root = other
		other.parent = nil
	}
}

func (d *DynamicAABBTree) balance(node *DATNode) *DATNode {
	a := node
	if a.IsLeaf() || a.height < 2 {
		return a
	}
	b := a.left
	c := a.right
	balance := c.height - b.height
	if balance > 1 {
		f := c.left
		g := c.right
		c.left = a
		c.parent = a.parent
		a.parent = c
		if c.parent != nil {
			if c.parent.left == a {
				c.parent.left = c
			} else {
				c.parent.right = c
			}
		} else {
			d.root = c
		}
		if f.height > g.height {
			c.right = f
			a.right = g
			g.parent = a
			a.aabb = b.aabb.GetUnion(g.aabb)
			c.aabb = a.aabb.GetUnion(f.aabb)
			a.height = 1 + math.Max(b.height, g.height)
			c.height = 1 + math.Max(a.height, f.height)
		} else {
			c.right = g
			a.right = f
			f.parent = a
			a.aabb = b.aabb.GetUnion(f.aabb)
			c.aabb = a.aabb.GetUnion(g.aabb)
			a.height = 1 + math.Max(b.height, g.height)
			c.height = 1 + math.Max(a.height, f.height)
		}
		return c
	}
	if balance < -1 {
		d2 := b.left
		e := b.right
		b.left = a
		b.parent = a.parent
		a.parent = b
		if b.parent != nil {
			if b.parent.left == a {
				b.parent.left = b
			} else {
				b.parent.right = b
			}
		} else {
			d.root = b
		}
		if d2.height > e.height {
			b.right = d2
			a.left = e
			e.parent = a
			a.aabb = c.aabb.GetUnion(e.aabb)
			b.aabb = a.aabb.GetUnion(d2.aabb)
			a.height = 1 + math.Max(c.height, e.height)
			b.height = 1 + math.Max(a.height, d2.height)
		} else {
			b.right = e
			a.left = d
			d2.parent = a
			a.aabb = c.aabb.GetUnion(d2.aabb)
			b.aabb = a.aabb.GetUnion(e.aabb)
			a.height = 1 + math.Max(c.height, d2.height)
			b.height = 1 + math.Max(a.height, e.height)
		}
		return b
	}
	return a
}
