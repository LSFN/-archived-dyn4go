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
