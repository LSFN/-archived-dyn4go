package dyn4go

import (
	"testing"
)

type ComparableInteger int

func (c ComparableInteger) CompareTo(c2 Comparable) int {
	return int(c - c2.(ComparableInteger))
}

func setupTree() *BinarySearchTree {
	tree := NewBinarySearchTree(false)
	tree.Insert(ComparableInteger(10))
	tree.Insert(ComparableInteger(3))
	tree.Insert(ComparableInteger(-3))
	tree.Insert(ComparableInteger(4))
	tree.Insert(ComparableInteger(0))
	tree.Insert(ComparableInteger(1))
	tree.Insert(ComparableInteger(11))
	tree.Insert(ComparableInteger(19))
	tree.Insert(ComparableInteger(6))
	tree.Insert(ComparableInteger(-1))
	tree.Insert(ComparableInteger(2))
	tree.Insert(ComparableInteger(9))
	tree.Insert(ComparableInteger(-4))
	return tree
}

func TestInsert(t *testing.T) {
	tree := setupTree()
	if tree.Contains(ComparableInteger(5)) {
		t.Error("Binary tree contains a value that it should not")
	}
	if !tree.Insert(ComparableInteger(5)) {
		t.Error("Failed to insert value into binary tree")
	}
	if !tree.Contains(ComparableInteger(5)) {
		t.Error("Binary tree does not contain a value that it should")
	}
	tree2 := NewBinarySearchTree(false)
	tree2.Insert(ComparableInteger(14))
	tree2.Insert(ComparableInteger(8))
	tree2.Insert(ComparableInteger(16))
	tree2.Insert(ComparableInteger(15))

	if tree.Contains(ComparableInteger(14)) || tree.Contains(ComparableInteger(8)) || tree.Contains(ComparableInteger(16)) || tree.Contains(ComparableInteger(15)) {
		t.Error("Binary tree contains multiple values that it should not")
	}
	tree.insertSubBinaryTree(tree2)
	if !(tree.Contains(ComparableInteger(14)) && tree.Contains(ComparableInteger(8)) && tree.Contains(ComparableInteger(16)) && tree.Contains(ComparableInteger(15))) {
		t.Error("Binary tree does not contain multiple values that it should")
	}
}

func TestRemove(t *testing.T) {
	tree := setupTree()
	if !tree.Remove(ComparableInteger(-3)) {
		t.Error("Element not removed from tree")
	}
	if tree.Contains(ComparableInteger(-3)) {
		t.Error("Tree still contains removed value")
	}
	if !(tree.Contains(ComparableInteger(-4)) && tree.Contains(ComparableInteger(0)) && tree.Contains(ComparableInteger(1)) && tree.Contains(ComparableInteger(2)) && tree.Contains(ComparableInteger(3))) {
		t.Error("Tree no longer contains removed node's surrounding values")
	}
	s := tree.GetSize()
	tree.RemoveMinimum()
	if tree.Contains(ComparableInteger(-4)) {
		t.Error("Removing minimum did not remove minimum")
	}
	if tree.GetSize() != s-1 {
		t.Error("Tree is of the wrong size after removing minimum")
	}
	n := tree.get(ComparableInteger(0))
	tree.removeMinimumNode(n)
	if tree.Contains(ComparableInteger(0)) {
		t.Error("Removing minimum node did not remove minimum node")
	}
	if tree.GetSize() != s-2 {
		t.Error("Tree is of the wrong size after removing minimum node")
	}
	tree.RemoveMaximum()
	if tree.Contains(ComparableInteger(19)) {
		t.Error("Removing maximum did not remove maximum")
	}
	if tree.GetSize() != s-3 {
		t.Error("Tree is of the wrong size after removing maximum")
	}
	n = tree.get(ComparableInteger(3))
	tree.removeMaximumNode(n)
	if tree.Contains(ComparableInteger(9)) {
		t.Error("Removing maximum node did not remove maximum node")
	}
	if tree.GetSize() != s-4 {
		t.Error("Tree is of the wrong size after removing maximum node")
	}
	tree.removeSubtree(ComparableInteger(3))
	if tree.Contains(ComparableInteger(3)) || tree.Contains(ComparableInteger(4)) || tree.Contains(ComparableInteger(6)) || tree.Contains(ComparableInteger(-1)) || tree.Contains(ComparableInteger(-4)) {
		t.Error("Tree contains supposedly removed elements")
	} 
	if tree.GetSize() != 2 {
		t.Error("Tree is of the wrong size")
	}
}

func TestRemoveNotFound(t *testing.T) {
	tree := setupTree()
	if tree.Remove(ComparableInteger(7)) {
		t.Error("Somehow managed to remove a non-existent node")
	}
	n := tree.get(ComparableInteger(-3))
	if tree.getHeightNode(n) != 4 {
		t.Error("Tree is not of correct height after removal operation")
	}
}

func TestGetDepth(t *testing.T) {
	tree := setupTree()
	if tree.GetHeight() != 6 {
		t.Error("Tree is not of correct depth")
	}
	n := tree.get(ComparableInteger(-3))
	if tree.getHeightNode(n) != 4 {
		t.Error("Tree is not of correct depth for a given node")
	}
}

func TestGetMinimum(t *testing.T) {
	tree := setupTree()
	if tree.GetMinimum() != ComparableInteger(-4) {
		t.Error("Incorrect minimum value retrieved")
	}
	n := tree.get(ComparableInteger(4))
	if tree.getMinimumNode(n).Data != ComparableInteger(4) {
		t.Error("Data does not have the correct minimum value")
	}
	n = tree.get(ComparableInteger(0))
	if tree.getMinimumNode(n).Data != ComparableInteger(-1) {
		t.Error("Data does not have correct minimum value as child of given node")
	}
}

func TestGetMaximum(t *testing.T) {
	tree := setupTree()
	if tree.GetMaximum() != ComparableInteger(19) {
		t.Error("Incorrect maximum value retrieved")
	}
	n := tree.get(ComparableInteger(-3))
	if tree.getMaximumNode(n).Data != ComparableInteger(2) {
		t.Error("Data does not have the correct maximum value")
	}
	n = tree.get(ComparableInteger(11))
	if tree.getMaximumNode(n).Data != ComparableInteger(19) {
		t.Error("Data does not have correct maximum value as child of given node")
	}
}

func TestIsEmpty(t *testing.T) {
	tree := setupTree()
	if tree.IsEmpty() {
		t.Error("Tree is empty when it shouldn't be")
	}
	tree2 := NewBinarySearchTree(false);
	if !tree2.IsEmpty() {
		t.Error("Tree contains items when it shouldn't")
	}
}

func TestClear(t *testing.T) {
	tree := setupTree()
	if tree.IsEmpty() {
		t.Error("Tree is empty before clearing")
	}
	if tree.GetSize() != 13 {
		t.Error("Tree is of the wrong size")
	}
	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Tree is not empty after it has been cleared")
	}
	if tree.GetSize() != 0 {
		t.Error("Tree is not of size 0 when it should be")
	}
	if tree.GetHeight() != 0 {
		t.Error("Tree is not of height 0 when it should be")
	}
	if tree.GetRoot() != nil {
		t.Error("Tree root is not null as it should be")
	}
}

func TestContains(t *testing.T) {
	tree := setupTree()
	if !tree.Contains(ComparableInteger(9)) {
		t.Error("Tree does not contain item it should")
	}
	if tree.Contains(ComparableInteger(14)) {
		t.Error("Tree contains an item that it shouldn't")
	}
	n := tree.get(ComparableInteger(-3))
	if !tree.containsNode(n) {
		t.Error("Tree doesn't contain a node that it should")
	}
	n2 := NewNode(ComparableInteger(-3), nil, nil, nil)
	if tree.containsNode(n2) {
		t.Error("Tree contains node that it shouldn't")
	}
}

func TestGet(t *testing.T) {
	tree := setupTree()
	if tree.get(ComparableInteger(-3)) == nil {
		t.Error("Get did returned nil")
	}
	if tree.get(ComparableInteger(45)) != nil {
		t.Error("Get did not return nil")
	}
}

func TestSize(t *testing.T) {
	tree := setupTree()
	if tree.GetSize() != 13 {
		t.Error("Size is not correct")
	}
	n := tree.get(ComparableInteger(-3))
	if tree.getSizeNode(n) != 6 {
		t.Error("Size of tree node is not correct")
	}
}

func TestIterator(t *testing.T) {
	tree := setupTree()
	it := tree.InOrderIterator();
	last := ComparableInteger(-9999)
	for it.HasNext() {
		i := it.Next()
		if i.CompareTo(last) < 0 {
			t.Error("Min to max values out of order")
		}
		last = i.(ComparableInteger)
	}
	it = tree.ReverseOrderIterator();
	last = ComparableInteger(9999)
	for it.HasNext() {
		i := it.Next()
		if i.CompareTo(last) > 0 {
			t.Error("Max to min alues out of order")
		}
		last = i.(ComparableInteger)
	}
}

func TestBalance(t *testing.T) {
	tree := setupTree()
	if tree.GetHeight() != 6 {
		t.Error("Tree not correctly unbalanced")
	}
	tree.SetSelfBalancing(true)
	if tree.GetHeight() != 4 {
		t.Log(tree.GetHeight())
		t.Error("Tree not correctly balanced")
	}
}