package two_three_tree

import "testing"

func TestBSFPrint(t *testing.T) {
	root := &TwoThreeTreeNode{
		Vals: []int{1, 2},
	}

	for _, n := range BSFPrint(root) {
		if n == nil {
			t.Log("nil")
			continue
		}
		if len(n.Vals) >= 2 {
			t.Log(n.Vals[0], n.Vals[1])
		} else {
			t.Log(n.Vals[1])
		}
	}
}

func TestFind(t *testing.T) {
	root := &TwoThreeTreeNode{
		Vals: []int{1},
		nodes: []*TwoThreeTreeNode{
			{
				Vals: []int{0},
			},
		},
	}

	n := Find(root, 0)
	if n == nil {
		t.Log("nil")
	} else {
		t.Log(n.Vals, n.nodes)
	}
}

func TestAddNode(t *testing.T) {
	root := AddNode(nil, 4)
	root = AddNode(root, 5)
	root = AddNode(root, 3)
	root = AddNode(root, 2)
	root = AddNode(root, 1)
	for i := 0; i < 30; i++ {
		root = AddNode(root, i)
	}
	for _, n := range BSFPrint(root) {
		if n == nil {
			t.Log("nil")
		} else {
			t.Log(n.Vals, n.nodes)
		}
	}
}
