package class03

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ts := NewTreeSet()
	ts.Add(5)
	ts.Add(3)
	ts.Add(8)
	ts.Add(1)
	ts.Add(4)
	ts.Add(7)
	ts.Add(9)

	fmt.Println("Size:", ts.Size())
	fmt.Println("Values:", ts.Values())
	fmt.Println("Contains 4:", ts.Contains(4))
	fmt.Println("Contains 6:", ts.Contains(6))

	ts.Remove(4)
	fmt.Println("Size after removal:", ts.Size())
	fmt.Println("Values after removal:", ts.Values())
	fmt.Println("Contains 4 after removal:", ts.Contains(4))
	return "Hello World!", nil
}

// 红黑树实现
type TreeSet struct {
	tree *redBlackTree
}

// NewTreeSet creates a new TreeSet.
func NewTreeSet() *TreeSet {
	return &TreeSet{tree: &redBlackTree{}}
}

// Add adds an element to the set.
func (ts *TreeSet) Add(value int) {
	ts.tree.Insert(value)
}

// Remove removes an element from the set.
func (ts *TreeSet) Remove(value int) {
	ts.tree.Delete(value)
}

// Contains checks if the set contains the given element.
func (ts *TreeSet) Contains(value int) bool {
	return ts.tree.Search(value) != nil
}

// Size returns the number of elements in the set.
func (ts *TreeSet) Size() int {
	return ts.tree.Size()
}

// Values returns all elements in the set in sorted order.
func (ts *TreeSet) Values() []int {
	return ts.tree.InOrderTraversal()
}

// redBlackTree is a red-black tree implementation.
type redBlackTree struct {
	root *node
	size int
}

// node is a node in the red-black tree.
type node struct {
	value       int
	left, right *node
	isRed       bool
}

// Insert inserts a value into the tree.
func (t *redBlackTree) Insert(value int) {
	t.root = t.insert(t.root, value)
	t.root.isRed = false
}

func (t *redBlackTree) insert(n *node, value int) *node {
	if n == nil {
		t.size++
		return &node{value: value, isRed: true}
	}

	if value < n.value {
		n.left = t.insert(n.left, value)
	} else if value > n.value {
		n.right = t.insert(n.right, value)
	}

	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}

	return n
}

// Delete deletes a value from the tree.
func (t *redBlackTree) Delete(value int) {
	if t.Search(value) == nil {
		return
	}

	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.isRed = true
	}

	t.root = t.delete(t.root, value)
	if t.root != nil {
		t.root.isRed = false
	}
}

func (t *redBlackTree) delete(n *node, value int) *node {
	if value < n.value {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = t.delete(n.left, value)
	} else {
		if isRed(n.left) {
			n = rotateRight(n)
		}
		if value == n.value && n.right == nil {
			return nil
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		if value == n.value {
			x := min(n.right)
			n.value = x.value
			n.right = t.deleteMin(n.right)
		} else {
			n.right = t.delete(n.right, value)
		}
	}
	return balance(n)
}

// Search searches for a value in the tree.
func (t *redBlackTree) Search(value int) *node {
	return t.search(t.root, value)
}

func (t *redBlackTree) search(n *node, value int) *node {
	if n == nil {
		return nil
	}
	if value < n.value {
		return t.search(n.left, value)
	} else if value > n.value {
		return t.search(n.right, value)
	} else {
		return n
	}
}

// Size returns the number of elements in the tree.
func (t *redBlackTree) Size() int {
	return t.size
}

// InOrderTraversal returns all elements in the tree in sorted order.
func (t *redBlackTree) InOrderTraversal() []int {
	var result []int
	t.inOrderTraversal(t.root, &result)
	return result
}

func (t *redBlackTree) inOrderTraversal(n *node, result *[]int) {
	if n != nil {
		t.inOrderTraversal(n.left, result)
		*result = append(*result, n.value)
		t.inOrderTraversal(n.right, result)
	}
}

// Helper functions for red-black tree operations

func isRed(n *node) bool {
	if n == nil {
		return false
	}
	return n.isRed
}

func rotateLeft(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n
	x.isRed = n.isRed
	n.isRed = true
	return x
}

func rotateRight(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n
	x.isRed = n.isRed
	n.isRed = true
	return x
}

func flipColors(n *node) {
	n.isRed = !n.isRed
	n.left.isRed = !n.left.isRed
	n.right.isRed = !n.right.isRed
}

func moveRedLeft(n *node) *node {
	flipColors(n)
	if isRed(n.right.left) {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		flipColors(n)
	}
	return n
}

func moveRedRight(n *node) *node {
	flipColors(n)
	if isRed(n.left.left) {
		n = rotateRight(n)
		flipColors(n)
	}
	return n
}

func balance(n *node) *node {
	if isRed(n.right) {
		n = rotateLeft(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}
	return n
}

func min(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *redBlackTree) deleteMin(n *node) *node {
	if n.left == nil {
		return nil
	}
	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}
	n.left = t.deleteMin(n.left)
	return balance(n)
}
