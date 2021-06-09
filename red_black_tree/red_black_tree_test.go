package red_black_tree

import "testing"

func TestFind(t *testing.T) {
	root := &RBTreeNode{
		Val: 5,
		left: &RBTreeNode{
			Val: 3,
			Red: true,
		},
	}
	n := Find(root, 3)
	if n == nil {
		t.Log("nil")
	} else {
		t.Log(n.Val)
	}
}

func TestBFSPrint(t *testing.T) {
	root := &RBTreeNode{
		Val: 0,
	}
	for _, n := range BFSPrint(root) {
		if n == nil {
			t.Log("nil")
			continue
		}
		t.Log(n.Val)
	}
}

func TestAddNode(t *testing.T) {
	root := AddNode(nil, 0)
	for i := 1; i <= 400; i++ {
		root = AddNode(root, i)
	}
	for _, n := range BFSPrint(root) {
		if n == nil {
			t.Log("nil")
			continue
		}
		t.Log(n.Val, n.Red)
	}
}
