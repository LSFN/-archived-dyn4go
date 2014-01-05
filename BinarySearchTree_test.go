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
	tree.InsertSubBinaryTree(tree2)
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
	n := tree.Get(ComparableInteger(0))
	tree.RemoveMinimumNode(n)
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
	n = tree.Get(ComparableInteger(3))
	tree.RemoveMaximumNode(n)
	if tree.Contains(ComparableInteger(9)) {
		t.Error("Removing maximum node did not remove maximum node")
	}
	if tree.GetSize() != s-4 {
		t.Error("Tree is of the wrong size after removing maximum node")
	}
	tree.RemoveSubtreeByComparable(ComparableInteger(3))
	if tree.Contains(ComparableInteger(3)) || tree.Contains(ComparableInteger(4)) || tree.Contains(ComparableInteger(6)) || tree.Contains(ComparableInteger(-1)) || tree.Contains(ComparableInteger(-4)) {
		t.Error("Tree contains supposedly removed elements")
	} 
	if tree.GetSize() != 2 {
		t.Error("Tree is of the wrong size")
	}
}