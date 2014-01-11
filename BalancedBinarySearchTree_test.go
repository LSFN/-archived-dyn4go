package dyn4go

import (
	"testing"
	"math"
)

func setupTree2() *BinarySearchTree {
	tree := NewBinarySearchTree(true)
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

func getHeightLimit(size int) float64 {
	return math.Log((float64(size) + 2.0) - 1.0) / math.Log(math.Phi) 
}

func TestInsert2(t *testing.T) {
	tree := setupTree2()
	if float64(tree.GetHeight()) >= getHeightLimit(tree.GetSize()) {
		t.Error("Tree unbalanced beyond permissible unbalancedness")
	}
	if tree.Contains(ComparableInteger(5)) {
		t.Error("Tree contains an item that it should not")
	}
	if !tree.Insert(ComparableInteger(5)) {
		t.Error("Failed to insert item into tree")
	}
	if !tree.Contains(ComparableInteger(5)) {
		t.Error("Tree does not contain an item that it should")
	}
	t2 := NewBinarySearchTree(false)
	t2.Insert(ComparableInteger(14))
	t2.Insert(ComparableInteger(8))
	t2.Insert(ComparableInteger(16))
	t2.Insert(ComparableInteger(15))
	if tree.Contains(ComparableInteger(14)) ||
	tree.Contains(ComparableInteger(8)) ||
	tree.Contains(ComparableInteger(16)) ||
	tree.Contains(ComparableInteger(15)) {
		t.Error("Tree contains items from subtree when it should not")
	}
	tree.insertSubtree(t2.Root)
	if !tree.Contains(ComparableInteger(14)) ||
	!tree.Contains(ComparableInteger(8)) ||
	!tree.Contains(ComparableInteger(16)) ||
	!tree.Contains(ComparableInteger(15)) {
		t.Error("Tree does not contain items added from the subtree")
	}
	if float64(tree.GetHeight()) >= getHeightLimit(tree.GetSize()) {
		t.Error("Tree unbalanced beyond permissible unbalancedness after insertions")
	}
}

func TestRemove2(t *testing.T) {
	tree := setupTree2()
	if !tree.Remove(ComparableInteger(-3)) {
		t.Error("Removing item that should exist in tree returned false")
	}
	if tree.Contains(ComparableInteger(-3)) {
		t.Error("Tree contains item that was just removed")
	}
	if !tree.Contains(ComparableInteger(-4)) ||
	!tree.Contains(ComparableInteger(0)) ||
	!tree.Contains(ComparableInteger(1)) ||
	!tree.Contains(ComparableInteger(2)) ||
	!tree.Contains(ComparableInteger(3)) {
		t.Error("Tree does not contain items that it should after removal")
	}
	size := tree.GetSize()
	tree.RemoveMinimum()
	if tree.Contains(ComparableInteger(-4)) {
		t.Error("Tree contins minimum item which was supposed to have been removed")
	}
	if tree.GetSize() != size - 1 {
		t.Error("Tree has not decreased size by 1 from original size after removal")
	}
	n := tree.get(ComparableInteger(10))
	tree.removeMinimumNode(n)
	if tree.Contains(ComparableInteger(4)) {
		t.Error("Tree contains minimum of subnode that should have been removed")
	}
	if tree.GetSize() != size - 2 {
		t.Error("Tree has not decreased size by 2 from original size after removals")
	}
	tree.RemoveMaximum()
	if tree.Contains(ComparableInteger(19)) {
		t.Error("Tree contains maximum that should have been removed")
	}
	if tree.GetSize() != size - 3 {
		t.Error("Tree has not decreased size by 3 from original size after removals")
	}
	n = tree.get(ComparableInteger(0))
	tree.removeMaximumNode(n)
	if tree.Contains(ComparableInteger(2)) {
		t.Error("Tree contains maximum of subnode that should have been removed")
	}
	if tree.GetSize() != size - 4 {
		t.Error("Tree has not decreased size by 4 from original size after removals")
	}
	if float64(tree.GetHeight()) >= getHeightLimit(tree.GetSize()) {
		t.Error("Tree unbalanced beyond permissible unbalancedness after removals")
	}
	tree.removeSubtree(ComparableInteger(0))
	if tree.Contains(ComparableInteger(0)) ||
	tree.Contains(ComparableInteger(-1)) ||
	tree.Contains(ComparableInteger(1)) ||
	tree.Contains(ComparableInteger(2)) {
		t.Error("Tree contains items that should have been removed with the subtree")
	}
	if tree.GetSize() != 5 {
		t.Error("Tree is of wrong size after removals")
	}
	if float64(tree.GetHeight()) >= getHeightLimit(tree.GetSize()) {
		t.Error("Tree unbalanced beyond permissible unbalancedness after subtree removal")
	}
}

func TestRemoveNotFound2(t *testing.T) {
	tree := setupTree2()
	if tree.Remove(ComparableInteger(7)) {
		t.Error("Somehow succeeded in removing item that is not in the tree")
	}
	if tree.removeNodeBool(NewNode(ComparableInteger(-3), nil, nil, nil)) {
		t.Error("Somehow succeeded in removing node that is not part of the tree")
	}
}

func TestGetDepth2(t *testing.T) {
	tree := setupTree2()
	if tree.GetHeight() != 4 {
		t.Error("Tree not of correct depth")
	}
	if tree.getHeightNode(tree.get(ComparableInteger(-3))) != 2 {
		t.Error("Subtree not of correct depth")
	}
}

func TestGetMinumim2(t *testing.T) {
	tree := setupTree2()
	if tree.GetMinimum() != ComparableInteger(-4) {
		t.Error("Tree minimum is not what it should be")
	}
	if tree.getMinimumNode(tree.get(ComparableInteger(10))).Data != ComparableInteger(4) {
		t.Error("Minimum of subtree is not the value expected")
	}
	if tree.getMinimumNode(tree.get(ComparableInteger(1))).Data != ComparableInteger(1) {
		t.Error("Minimum of second subtree is not the value expected")
	}
}

func TestGetMaximum2(t *testing.T) {
	tree := setupTree2()
	if tree.GetMaximum() != ComparableInteger(19) {
		t.Error("Tree maximum is not what it should be")
	}
	if tree.getMaximumNode(tree.get(ComparableInteger(-3))).Data != ComparableInteger(-1) {
		t.Error("Maximum of subtree is not the value expected")
	}
	if tree.getMaximumNode(tree.get(ComparableInteger(6))).Data != ComparableInteger(9) {
		t.Error("Maximum of second subtree is not the value expected")
	}
}

func TestIsEmpty2(t *testing.T) {
	tree := setupTree2()
	if tree.IsEmpty() {
		t.Error("Tree is empty when it should not be")
	}
	tree2 := NewBinarySearchTree(false)
	if !tree2.IsEmpty() {
		t.Error("Tree is not empty when it should be")
	}
}

func TestClear2(t *testing.T) {
	tree := setupTree2()
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

func TestContains2(t *testing.T) {
	tree := setupTree2()
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

func TestGet2(t *testing.T) {
	tree := setupTree2()
	if tree.get(ComparableInteger(-3)) == nil {
		t.Error("Get did returned nil")
	}
	if tree.get(ComparableInteger(45)) != nil {
		t.Error("Get did not return nil")
	}
}

func TestSize2(t *testing.T) {
	tree := setupTree2()
	if tree.GetSize() != 13 {
		t.Error("Size is not correct")
	}
	n := tree.get(ComparableInteger(-3))
	if tree.getSizeNode(n) != 3 {
		t.Error("Size of tree node is not correct")
	}
}

func TestIterator2(t *testing.T) {
	tree := setupTree2()
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