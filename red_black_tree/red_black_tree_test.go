package red_black_tree

import (
	"math/rand"
	"testing"
	"time"
)

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
	root := &RBTreeNode{
		Val: 10,
	}
	for i := 0; i < 20; i++ {
		root = AddNode(root, i)
	}

	for _, n := range BFSPrint(root) {
		if n == nil {
			t.Log("nil")
			continue
		}
		p := n.parent
		if p == nil {
			t.Log(n.Val, n.Red, "nil")
		} else {
			t.Log(n.Val, n.Red, p.Val)
		}
	}
}

func TestDelNode(t *testing.T) {
	root := &RBTreeNode{
		Val: 10,
	}
	for i := 0; i < 20; i++ {
		root = AddNode(root, i)
	}
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		root = DelNode(root, rand.Intn(20))
	}

	for _, n := range BFSPrint(root) {
		if n == nil {
			t.Log("nil")
			continue
		}
		p := n.parent
		if p == nil {
			t.Log(n.Val, n.Red, "nil")
		} else {
			t.Log(n.Val, n.Red, p.Val)
		}

	}
	t.Log("done")
}
