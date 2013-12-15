package dyn4go

import (
	"math"
)

type BinarySearchTree struct {
	root          *node
	size          int
	selfBalancing bool
}

type Comparable interface {
	CompareTo(Comparable) int
}

type node struct {
	comparable Comparable
	parent     *node
	left       *node
	right      *node
}

func newNode(comparable Comparable, parent, left, right *node) *node {
	if comparable == nil {
		return nil
	}
	n := new(node)
	n.comparable = comparable
	n.parent = parent
	n.left = left
	n.right = right
	return n
}

func (n *node) CompareTo(other *node) int {
	return n.comparable.CompareTo(other.comparable)
}

func (n *node) IsLeftChild() bool {
	if n.parent == nil {
		return false
	}
	return n.parent.left == n
}

func (n *node) countNodesInTree() int {
	var count int = 1
	if n.left != nil {
		count += n.left.countNodesInTree()
	}
	if n.right != nil {
		count += n.right.countNodesInTree()
	}
	return count
}

type TreeIterator struct {
	nodeQueue   []*node
	isAscending bool
}

func (n *node) NewTreeIterator(ascending bool) *TreeIterator {
	iterator := new(TreeIterator)
	iterator.isAscending = ascending
	// count the nodes in the tree
	numNodes := n.countNodesInTree()
	iterator.nodeQueue = make([]*node, numNodes)
	iterator.assembleQueue(n, ascending)
	return iterator
}

func (iter *TreeIterator) assembleQueue(n *node, ascending bool) {
	if ascending {
		if n.left != nil {
			iter.assembleQueue(n.left, ascending)
		}
		if n.right != nil {
			iter.assembleQueue(n.right, ascending)
		}
	} else {
		if n.right != nil {
			iter.assembleQueue(n.right, ascending)
		}
		if n.left != nil {
			iter.assembleQueue(n.left, ascending)
		}
	}
	iter.nodeQueue = append(iter.nodeQueue, n)
}

func (iter *TreeIterator) Next() Comparable {
	result := iter.nodeQueue[0].comparable
	iter.nodeQueue = iter.nodeQueue[1:]
	return result
}

func (iter *TreeIterator) HasNext() bool {
	return len(iter.nodeQueue) > 0
}

func newBinarySearchTree(selfBalancing bool) *BinarySearchTree {
	b := new(BinarySearchTree)
	b.selfBalancing = selfBalancing
	b.size = 0
	b.root = nil
	return b
}

func copyBinarySearchTree(oldTree *BinarySearchTree, selfBalancing bool) {
	b := new(BinarySearchTree)
	b.selfBalancing = selfBalancing
	b.insertSubBinaryTree(oldTree)
}

func (b *BinarySearchTree) IsSelfBalancing() bool {
	return b.selfBalancing
}

func (b *BinarySearchTree) SetSelfBalancing(selfBalancing bool) {
	if b.selfBalancing {
		b.selfBalancing = selfBalancing
	} else {
		if !selfBalancing {
			if b.size > 2 {
				b.balanceTree()
			}
		}
	}
}

func (b *BinarySearchTree) Insert(comparable Comparable) bool {
	if comparable == nil {
		return false
	}
	node := newNode(comparable, nil, nil, nil)
	b.insert(node)
	return true
}

func (b *BinarySearchTree) Remove(comparable Comparable) bool {
	if comparable == nil || b.root == nil {
		return false
	}
	n := b.removeByComparable(comparable, b.root)
	return n != nil
}

func (b *BinarySearchTree) RemoveMinimum() Comparable {
	if b.root == nil {
		return nil
	}
	return b.removeMinimum(b.root).comparable
}

func (b *BinarySearchTree) RemoveMaximum() Comparable {
	if b.root == nil {
		return nil
	}
	return b.removeMaximum(b.root).comparable
}

func (b *BinarySearchTree) GetMinimum() Comparable {
	if b.root == nil {
		return nil
	}
	return b.getMinimum(b.root).comparable
}

func (b *BinarySearchTree) GetMaximum() Comparable {
	if b.root == nil {
		return nil
	}
	return b.getMaximum(b.root).comparable
}

func (b *BinarySearchTree) Contains(comparable Comparable) bool {
	if comparable == nil || b.root == nil {
		return false
	}
	node := b.containsByComparable(b.root, comparable)
	return node != nil
}

func (b *BinarySearchTree) GetRoot() Comparable {
	if b.root == nil {
		return nil
	}
	return b.root.comparable
}

func (b *BinarySearchTree) Clear() {
	b.root = nil
	b.size = 0
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.root == nil
}

func (b *BinarySearchTree) GetHeight() int {
	return b.getHeight(b.root)
}

func (b *BinarySearchTree) GetSize() int {
	return b.size
}

func (b *BinarySearchTree) Iterator() *TreeIterator {
	return b.InOrderIterator()
}

func (b *BinarySearchTree) InOrderIterator() *TreeIterator {
	return b.root.NewTreeIterator(true)
}

func (b *BinarySearchTree) ReverseOrderIterator() *TreeIterator {
	return b.root.NewTreeIterator(false)
}

func (b *BinarySearchTree) getMinimum(n *node) *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (b *BinarySearchTree) getMaximum(n *node) *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (b *BinarySearchTree) getRootNode() *node {
	return b.root
}

func (b *BinarySearchTree) removeMinimum(n *node) *node {
	n = b.getMinimum(n)
	if n == nil {
		return nil
	}
	if n == b.root {
		b.root = n.right
	} else if n.parent.right == n {
		n.parent.right = n.right
	} else {
		n.parent.left = n.right
	}
	b.size--
	return n
}

func (b *BinarySearchTree) removeMaximum(n *node) *node {
	n = b.getMaximum(n)
	if n == nil {
		return nil
	}
	if n == b.root {
		b.root = n.left
	} else if n.parent.right == n {
		n.parent.right = n.left
	} else {
		n.parent.left = n.left
	}
	b.size--
	return n
}

func (b *BinarySearchTree) getHeight(n *node) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	return 1 + int(math.Max(float64(b.getHeight(n.left)), float64(b.getHeight(n.right))))
}

func (b *BinarySearchTree) getSize(n *node) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	return 1 + b.getSize(n.left) + b.getSize(n.right)
}

func (b *BinarySearchTree) contains(n *node) bool {
	if n == nil || b.root == nil {
		return false
	}
	if n == b.root {
		return true
	}
	curr := b.root
	for curr != nil {
		if curr == n {
			return true
		}
		diff := n.CompareTo(curr)
		if diff == 0 {
			return false
		} else if diff < 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (b *BinarySearchTree) get(comparable Comparable) *node {
	if comparable == nil || b.root == nil {
		return nil
	}
	return b.containsByComparable(b.root, comparable)
}

func (b *BinarySearchTree) insertSubtree(n *node) bool {
	if n == nil {
		return false
	}
	iter := n.NewTreeIterator(true)
	for iter.HasNext() {
		n2 := newNode(iter.Next(), nil, nil, nil)
		b.insert(n2)
	}
	return true
}

func (b *BinarySearchTree) insertSubBinaryTree(b2 *BinarySearchTree) bool {
	if b2 == nil {
		return false
	}
	if b2.root == nil {
		return true
	}
	iter := b.InOrderIterator()
	for iter.HasNext() {
		n2 := newNode(iter.Next(), nil, nil, nil)
		b.insert(n2)
	}
	return true
}

func (b *BinarySearchTree) removeSubtreeByComparable(comparable Comparable) bool {
	if comparable == nil || b.root == nil {
		return false
	}
	n := b.root
	for n != nil {
		diff := comparable.CompareTo(n.comparable)
		if diff < 0 {
			n = n.left
		} else if diff > 0 {
			n = n.right
		} else {
			if n.IsLeftChild() {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
			b.size -= b.getSize(n)
			if b.selfBalancing {
				b.balanceTreeByNode(n.parent)
			}
			return true
		}
	}
	return false
}

func (b *BinarySearchTree) removeSubtree(n *node) bool {
	if n == nil || b.root == nil {
		return false
	}
	if b.root == n {
		b.root = nil
	} else {
		if b.contains(n) {
			if n.IsLeftChild() {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
			b.size -= b.getSize(n)
			if b.selfBalancing {
				b.balanceTreeByNode(n.parent)
			}
			return true
		}
	}
	return false
}

func (b *BinarySearchTree) insert(n *node) bool {
	if b.root == nil {
		b.root = n
		b.size += 1
		return true
	} else {
		return b.insertAtNode(n, b.root)
	}
}

func (b *BinarySearchTree) insertAtNode(n *node, at *node) bool {
	if at == nil {
		return false
	}
	for at != nil {
		if n.CompareTo(at) < 0 {
			if at.left == nil {
				at.left = n
				n.parent = at
				break
			} else {
				at = at.left
			}
		} else {
			if at.right == nil {
				at.right = n
				n.parent = at
				break
			} else {
				at = at.right
			}
		}
	}
	b.size++
	if b.selfBalancing {
		b.balanceTreeByNode(at)
	}
	return true
}

func (b *BinarySearchTree) remove(n *node) bool {
	if n == nil || b.root == nil {
		return false
	}
	if b.contains(n) {
		b.remove(n)
		return true
	}
	return false
}

func (b *BinarySearchTree) removeByComparable(comparable Comparable, n *node) *node {
	for n != nil {
		diff := comparable.CompareTo(n.comparable)
		if diff < 0 {
			n = n.left
		} else if diff > 0 {
			n = n.right
		} else {
			b.removeNode(n)
			return n
		}
	}
	return nil
}

func (b *BinarySearchTree) removeNode(n *node) {
	isLeftChild := n.IsLeftChild()
	if n.left != nil && n.right != nil {
		min := b.getMinimum(n.right)
		if min != n.right {
			min.parent.left = min.right
			if min.right != nil {
				min.right.parent = min.parent
			}
			min.right = n.right
		}
		if n.right != nil {
			n.right.parent = min
		}
		if n.left != nil {
			n.left.parent = min
		}
		if n == b.root {
			b.root = min
		} else if isLeftChild {
			n.parent.left = min
		} else {
			n.parent.right = min
		}
		min.left = n.left
		min.parent = n.parent
		if b.selfBalancing {
			b.balanceTreeByNode(min.parent)
		}
	} else if n.left != nil {
		if n == b.root {
			b.root = n.left
		} else if isLeftChild {
			n.parent.left = n.left
		} else {
			n.parent.right = n.left
		}
		if n.left != nil {
			n.left.parent = n.parent
		}
	} else if n.right != nil {
		if n == b.root {
			b.root = n.right
		} else if isLeftChild {
			n.parent.left = n.right
		} else {
			n.parent.right = n.right
		}
		if n.right != nil {
			n.right.parent = n.parent
		}
	} else {
		if n == b.root {
			b.root = nil
		} else if isLeftChild {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	}
	b.size--
}

func (b *BinarySearchTree) containsByComparable(n *node, comparable Comparable) *node {
	for n != nil {
		nData := n.comparable
		diff := comparable.CompareTo(nData)
		if diff == 0 {
			return n
		} else if diff < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	return nil
}

func (b *BinarySearchTree) balanceTree() {
	root := b.root
	b.root = nil
	b.size = 0
	iter := root.NewTreeIterator(true)
	for iter.HasNext() {
		n := newNode(iter.Next(), nil, nil, nil)
		b.insert(n)
	}
}

func (b *BinarySearchTree) balanceTreeByNode(n *node) {
	for n != nil {
		n = b.balance(n)
		n = n.parent
	}
}

func (b *BinarySearchTree) balance(n *node) *node {
	if n == nil {
		return nil
	}
	if b.getHeight(n) < 2 {
		return n
	}
	p := n.parent
	l := n.left
	r := n.right
	lh := b.getHeight(l)
	rh := b.getHeight(r)
	balance := lh - rh
	if balance > 1 {
		lch := b.getHeight(l.right)
		if lch > 1 {
			c := l.right
			l.right = c.left
			if c.left != nil {
				c.left.parent = l
			}
			c.left = l
			l.parent = c
			n.left = c
			c.parent = n
		}
		c := n.left
		n.left = c.right
		if c.right != nil {
			c.right.parent = n
		}
		c.right = n
		c.parent = n.parent
		n.parent = c
		if p != nil {
			if p.left == n {
				p.left = c
			} else {
				p.right = c
			}
		} else {
			b.root = c
		}
		return c
	}
	if balance < -1 {
		rch := b.getHeight(r.left)
		if rch > 1 {
			d := r.left
			r.left = d.right
			if d.right != nil {
				d.right.parent = r
			}
			d.right = r
			r.parent = d
			n.right = d
			d.parent = n
		}
		d := n.right
		n.right = d.left
		if d.left != nil {
			d.left.parent = n
		}
		d.left = n
		d.parent = n.parent
		n.parent = d
		if p != nil {
			if p.left == n {
				p.left = d
			} else {
				p.right = d
			}
		} else {
			b.root = d
		}
		return d
	}

	return n
}
