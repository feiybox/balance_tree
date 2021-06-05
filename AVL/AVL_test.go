package AVL

import "testing"

func TestFind(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		left: &TreeNode{
			Val: 3,
		},
		right: &TreeNode{
			Val: 7,
		},
	}

	val := 7
	node := Find(root, val)
	if node == nil {
		t.Errorf("fail, want=%d, get nil", val)
		return
	} else if node.Val != val {
		t.Errorf("fail, want=%d, get=%d", val, node.Val)
		return
	}
	t.Log("pass")
}

func TestGetHeight(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Dep: 2,
		left: &TreeNode{
			Val: 3,
			Dep: 1,
		},
	}

	height := GetHeight(root)
	wantHeight := 2
	if height != wantHeight {
		t.Errorf("fail, want=%d, get=%d", wantHeight, height)
		return
	}
	t.Log("pass")
}

func TestBFSPrint(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Dep: 3,
		left: &TreeNode{
			Val: 3,
			Dep: 2,
			left: &TreeNode{
				Val: 2,
				Dep: 1,
			},
		},
	}
	nodeList := BFSPrint(root)
	for _, n := range nodeList {
		if n == nil {
			t.Logf("nil")
			continue
		}
		t.Log(n.Val)
	}
}

func TestAddNode(t *testing.T) {
	root := AddNode(nil, 5)
	root = AddNode(root, 15)
	root = AddNode(root, 6)
	root = AddNode(root, 50)
	root = AddNode(root, 31)

	nodeList := BFSPrint(root)
	for _, n := range nodeList {
		if n == nil {
			t.Logf("nil")
			continue
		}
		t.Log(n.Val)
	}
	t.Log("pass")
}

func TestDelNode(t *testing.T) {
	root := AddNode(nil, 5)
	root = AddNode(root, 15)
	root = AddNode(root, 6)
	root = AddNode(root, 50)
	root = AddNode(root, 31)

	root = DelNode(root, 5)
	root = DelNode(root, 31)
	root = DelNode(root, 15)
	root = DelNode(root, 60)

	nodeList := BFSPrint(root)
	for _, n := range nodeList {
		if n == nil {
			t.Logf("nil")
			continue
		}
		t.Log(n.Val)
	}
	t.Log("pass")
}
