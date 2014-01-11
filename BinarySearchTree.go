package dyn4go

import (
	"math"
)

type BinarySearchTree struct {
	Root          *Node
	size          int
	selfBalancing bool
}

type Comparable interface {
	CompareTo(Comparable) int
}

type Node struct {
	Data   Comparable
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(comparable Comparable, parent, left, right *Node) *Node {
	if comparable == nil {
		return nil
	}
	n := new(Node)
	n.Data = comparable
	n.Parent = parent
	n.Left = left
	n.Right = right
	return n
}

func (n *Node) CompareTo(other *Node) int {
	return n.Data.CompareTo(other.Data)
}

func (n *Node) IsLeftChild() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Left == n
}

func (n *Node) CountNodesInTree() int {
	var count int = 1
	if n.Left != nil {
		count += n.Left.CountNodesInTree()
	}
	if n.Right != nil {
		count += n.Right.CountNodesInTree()
	}
	return count
}

type TreeIterator struct {
	NodeQueue   []*Node
	isAscending bool
}

func (n *Node) NewTreeIterator(ascending bool) *TreeIterator {
	iterator := new(TreeIterator)
	iterator.isAscending = ascending
	// count the Nodes in the tree
	iterator.NodeQueue = make([]*Node, 0)
	iterator.AssembleQueue(n, ascending)
	return iterator
}

func (iter *TreeIterator) AssembleQueue(n *Node, ascending bool) {
	if ascending {
		if n.Left != nil {
			iter.AssembleQueue(n.Left, ascending)
		}
	} else {
		if n.Right != nil {
			iter.AssembleQueue(n.Right, ascending)
		}
	}
	iter.NodeQueue = append(iter.NodeQueue, n)
	if ascending {
		if n.Right != nil {
			iter.AssembleQueue(n.Right, ascending)
		}
	} else {
		if n.Left != nil {
			iter.AssembleQueue(n.Left, ascending)
		}
	}
}

func (iter *TreeIterator) Next() Comparable {
	result := iter.NodeQueue[0].Data
	iter.NodeQueue = iter.NodeQueue[1:]
	return result
}

func (iter *TreeIterator) HasNext() bool {
	return len(iter.NodeQueue) > 0
}

func NewBinarySearchTree(selfBalancing bool) *BinarySearchTree {
	b := new(BinarySearchTree)
	b.selfBalancing = selfBalancing
	b.size = 0
	b.Root = nil
	return b
}

func CopyBinarySearchTree(oldTree *BinarySearchTree, selfBalancing bool) {
	b := new(BinarySearchTree)
	b.selfBalancing = selfBalancing
	b.insertSubBinaryTree(oldTree)
}

func (b *BinarySearchTree) IsSelfBalancing() bool {
	return b.selfBalancing
}

func (b *BinarySearchTree) SetSelfBalancing(selfBalancing bool) {
	if selfBalancing && !b.selfBalancing {
		b.selfBalancing = true
		if b.size > 2 {
			b.balanceTree()
		}
	}
	b.selfBalancing = selfBalancing
}

func (b *BinarySearchTree) Insert(comparable Comparable) bool {
	if comparable == nil {
		return false
	}
	n := NewNode(comparable, nil, nil, nil)
	b.insertNode(n)
	return true
}

func (b *BinarySearchTree) Remove(comparable Comparable) bool {
	if comparable == nil || b.Root == nil {
		return false
	}
	n := b.removeByComparable(comparable, b.Root)
	return n != nil
}

func (b *BinarySearchTree) RemoveMinimum() Comparable {
	if b.Root == nil {
		return nil
	}
	return b.removeMinimumNode(b.Root).Data
}

func (b *BinarySearchTree) RemoveMaximum() Comparable {
	if b.Root == nil {
		return nil
	}
	return b.removeMaximumNode(b.Root).Data
}

func (b *BinarySearchTree) GetMinimum() Comparable {
	if b.Root == nil {
		return nil
	}
	return b.getMinimumNode(b.Root).Data
}

func (b *BinarySearchTree) GetMaximum() Comparable {
	if b.Root == nil {
		return nil
	}
	return b.getMaximumNode(b.Root).Data
}

func (b *BinarySearchTree) Contains(comparable Comparable) bool {
	if comparable == nil || b.Root == nil {
		return false
	}
	Node := b.containsNodeComparable(b.Root, comparable)
	return Node != nil
}

func (b *BinarySearchTree) GetRoot() Comparable {
	if b.Root == nil {
		return nil
	}
	return b.Root.Data
}

func (b *BinarySearchTree) Clear() {
	b.Root = nil
	b.size = 0
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.Root == nil
}

func (b *BinarySearchTree) GetHeight() int {
	return b.getHeightNode(b.Root)
}

func (b *BinarySearchTree) GetSize() int {
	return b.size
}

func (b *BinarySearchTree) Iterator() *TreeIterator {
	return b.InOrderIterator()
}

func (b *BinarySearchTree) InOrderIterator() *TreeIterator {
	return b.Root.NewTreeIterator(true)
}

func (b *BinarySearchTree) ReverseOrderIterator() *TreeIterator {
	return b.Root.NewTreeIterator(false)
}

func (b *BinarySearchTree) getMinimumNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (b *BinarySearchTree) getMaximumNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (b *BinarySearchTree) getRootNode() *Node {
	return b.Root
}

func (b *BinarySearchTree) removeMinimumNode(n *Node) *Node {
	n = b.getMinimumNode(n)
	if n == nil {
		return nil
	}
	if n == b.Root {
		b.Root = n.Right
	} else if n.Parent.Right == n {
		n.Parent.Right = n.Right
	} else {
		n.Parent.Left = n.Right
	}
	b.size--
	return n
}

func (b *BinarySearchTree) removeMaximumNode(n *Node) *Node {
	n = b.getMaximumNode(n)
	if n == nil {
		return nil
	}
	if n == b.Root {
		b.Root = n.Left
	} else if n.Parent.Right == n {
		n.Parent.Right = n.Left
	} else {
		n.Parent.Left = n.Left
	}
	b.size--
	return n
}

func (b *BinarySearchTree) getHeightNode(n *Node) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	return 1 + int(math.Max(float64(b.getHeightNode(n.Left)), float64(b.getHeightNode(n.Right))))
}

func (b *BinarySearchTree) getSizeNode(n *Node) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	return 1 + b.getSizeNode(n.Left) + b.getSizeNode(n.Right)
}

func (b *BinarySearchTree) containsNode(n *Node) bool {
	if n == nil || b.Root == nil {
		return false
	}
	if n == b.Root {
		return true
	}
	curr := b.Root
	for curr != nil {
		if curr == n {
			return true
		}
		diff := n.CompareTo(curr)
		if diff == 0 {
			return false
		} else if diff < 0 {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

func (b *BinarySearchTree) get(comparable Comparable) *Node {
	if comparable == nil || b.Root == nil {
		return nil
	}
	return b.containsNodeComparable(b.Root, comparable)
}

func (b *BinarySearchTree) insertSubtree(n *Node) bool {
	if n == nil {
		return false
	}
	iter := n.NewTreeIterator(true)
	for iter.HasNext() {
		n2 := NewNode(iter.Next(), nil, nil, nil)
		b.insertNode(n2)
	}
	return true
}

func (b *BinarySearchTree) insertSubBinaryTree(b2 *BinarySearchTree) bool {
	if b2 == nil {
		return false
	}
	if b2.Root == nil {
		return true
	}
	iter := b2.InOrderIterator()
	for iter.HasNext() {
		n2 := NewNode(iter.Next(), nil, nil, nil)
		b.insertNode(n2)
	}
	return true
}

func (b *BinarySearchTree) removeSubtree(comparable Comparable) bool {
	if comparable == nil || b.Root == nil {
		return false
	}
	n := b.Root
	for n != nil {
		diff := comparable.CompareTo(n.Data)
		if diff < 0 {
			n = n.Left
		} else if diff > 0 {
			n = n.Right
		} else {
			if n.IsLeftChild() {
				n.Parent.Left = nil
			} else {
				n.Parent.Right = nil
			}
			b.size -= b.getSizeNode(n)
			if b.selfBalancing {
				b.balanceTreeByNode(n.Parent)
			}
			return true
		}
	}
	return false
}

func (b *BinarySearchTree) removeSubtreeNode(n *Node) bool {
	if n == nil || b.Root == nil {
		return false
	}
	if b.Root == n {
		b.Root = nil
	} else {
		if b.containsNode(n) {
			if n.IsLeftChild() {
				n.Parent.Left = nil
			} else {
				n.Parent.Right = nil
			}
			b.size -= b.getSizeNode(n)
			if b.selfBalancing {
				b.balanceTreeByNode(n.Parent)
			}
			return true
		}
	}
	return false
}

func (b *BinarySearchTree) insertNode(n *Node) bool {
	if b.Root == nil {
		b.Root = n
		b.size += 1
		return true
	} else {
		return b.insertAtNode(n, b.Root)
	}
}

func (b *BinarySearchTree) insertAtNode(n *Node, at *Node) bool {
	if at == nil {
		return false
	}
	for at != nil {
		if n.CompareTo(at) < 0 {
			if at.Left == nil {
				at.Left = n
				n.Parent = at
				break
			} else {
				at = at.Left
			}
		} else {
			if at.Right == nil {
				at.Right = n
				n.Parent = at
				break
			} else {
				at = at.Right
			}
		}
	}
	b.size++
	if b.selfBalancing {
		b.balanceTreeByNode(at)
	}
	return true
}

func (b *BinarySearchTree) removeNodeBool(n *Node) bool {
	if n == nil || b.Root == nil {
		return false
	}
	if b.containsNode(n) {
		b.removeNode(n)
		return true
	}
	return false
}

func (b *BinarySearchTree) removeByComparable(comparable Comparable, n *Node) *Node {
	for n != nil {
		diff := comparable.CompareTo(n.Data)
		if diff < 0 {
			n = n.Left
		} else if diff > 0 {
			n = n.Right
		} else {
			b.removeNode(n)
			return n
		}
	}
	return nil
}

func (b *BinarySearchTree) removeNode(n *Node) {
	isLeftChild := n.IsLeftChild()
	if n.Left != nil && n.Right != nil {
		min := b.getMinimumNode(n.Right)
		if min != n.Right {
			min.Parent.Left = min.Right
			if min.Right != nil {
				min.Right.Parent = min.Parent
			}
			min.Right = n.Right
		}
		if n.Right != nil {
			n.Right.Parent = min
		}
		if n.Left != nil {
			n.Left.Parent = min
		}
		if n == b.Root {
			b.Root = min
		} else if isLeftChild {
			n.Parent.Left = min
		} else {
			n.Parent.Right = min
		}
		min.Left = n.Left
		min.Parent = n.Parent
		if b.selfBalancing {
			b.balanceTreeByNode(min.Parent)
		}
	} else if n.Left != nil {
		if n == b.Root {
			b.Root = n.Left
		} else if isLeftChild {
			n.Parent.Left = n.Left
		} else {
			n.Parent.Right = n.Left
		}
		if n.Left != nil {
			n.Left.Parent = n.Parent
		}
	} else if n.Right != nil {
		if n == b.Root {
			b.Root = n.Right
		} else if isLeftChild {
			n.Parent.Left = n.Right
		} else {
			n.Parent.Right = n.Right
		}
		if n.Right != nil {
			n.Right.Parent = n.Parent
		}
	} else {
		if n == b.Root {
			b.Root = nil
		} else if isLeftChild {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
	}
	b.size--
}

func (b *BinarySearchTree) containsNodeComparable(n *Node, comparable Comparable) *Node {
	for n != nil {
		nData := n.Data
		diff := comparable.CompareTo(nData)
		if diff == 0 {
			return n
		} else if diff < 0 {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return nil
}

func (b *BinarySearchTree) balanceTree() {
	Root := b.Root
	b.Root = nil
	b.size = 0
	iter := Root.NewTreeIterator(true)
	for iter.HasNext() {
		n := NewNode(iter.Next(), nil, nil, nil)
		b.insertNode(n)
	}
}

func (b *BinarySearchTree) balanceTreeByNode(n *Node) {
	for n != nil {
		n = b.balance(n)
		n = n.Parent
	}
}

func (b *BinarySearchTree) balance(n *Node) *Node {
	if n == nil {
		return nil
	}
	if b.getHeightNode(n) < 2 {
		return n
	}
	p := n.Parent
	l := n.Left
	r := n.Right
	lh := b.getHeightNode(l)
	rh := b.getHeightNode(r)
	balance := lh - rh
	if balance > 1 {
		lch := b.getHeightNode(l.Right)
		if lch > 1 {
			c := l.Right
			l.Right = c.Left
			if c.Left != nil {
				c.Left.Parent = l
			}
			c.Left = l
			l.Parent = c
			n.Left = c
			c.Parent = n
		}
		c := n.Left
		n.Left = c.Right
		if c.Right != nil {
			c.Right.Parent = n
		}
		c.Right = n
		c.Parent = n.Parent
		n.Parent = c
		if p != nil {
			if p.Left == n {
				p.Left = c
			} else {
				p.Right = c
			}
		} else {
			b.Root = c
		}
		return c
	}
	if balance < -1 {
		rch := b.getHeightNode(r.Left)
		if rch > 1 {
			d := r.Left
			r.Left = d.Right
			if d.Right != nil {
				d.Right.Parent = r
			}
			d.Right = r
			r.Parent = d
			n.Right = d
			d.Parent = n
		}
		d := n.Right
		n.Right = d.Left
		if d.Left != nil {
			d.Left.Parent = n
		}
		d.Left = n
		d.Parent = n.Parent
		n.Parent = d
		if p != nil {
			if p.Left == n {
				p.Left = d
			} else {
				p.Right = d
			}
		} else {
			b.Root = d
		}
		return d
	}

	return n
}
