package red_black_tree

type RBTreeNode struct {
	Val   int
	Red   bool
	left  *RBTreeNode
	right *RBTreeNode
}

// 查找树上的节点
func Find(root *RBTreeNode, val int) *RBTreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return Find(root.left, val)
	}
	return Find(root.right, val)
}

func BFSPrint(root *RBTreeNode) []*RBTreeNode {
	nodeList := make([]*RBTreeNode, 0, 10)
	nodeList = append(nodeList, root)
	for i := 0; i < len(nodeList); i++ {
		n := nodeList[i]
		if n == nil {
			continue
		}
		nodeList = append(nodeList, n.left)
		nodeList = append(nodeList, n.right)
	}
	return nodeList
}

// 增加节点
func AddNode(root *RBTreeNode, val int) *RBTreeNode {
	// if root is nil, init root and set color black
	if root == nil {
		return &RBTreeNode{
			Val: val,
			Red: false,
		}
	}
	root = addNode(root, val)
	if root.Red {
		root.Red = false
	}
	return root
}

// add new node
func addNode(node *RBTreeNode, val int) *RBTreeNode {
	if node == nil {
		return &RBTreeNode{
			Val: val,
			Red: true,
		}
	}

	// 插入节点
	if node.Val == val {
		return node
	}

	var son *RBTreeNode
	var insertLeft bool
	if node.Val > val {
		node.left = addNode(node.left, val)
		son = node.left
		insertLeft = true
	} else {
		node.right = addNode(node.right, val)
		son = node.right
	}

	// rebalance and recolor
	if son.Red && ((son.left != nil && son.left.Red) || (son.right != nil && son.right.Red)) {
		if node.left != nil && node.left.Red && node.right != nil && node.right.Red {
			node.Red = true
			node.left.Red = false
			node.right.Red = false
		} else if son.left != nil && son.left.Red {
			if insertLeft {
				node = llRotate(node)
			} else {
				node = rlRotate(node)
			}
			node.Red = false
			node.left.Red = true
			node.right.Red = true
		} else {
			if insertLeft {
				node = lrRotate(node)
			} else {
				node = rrRotate(node)
			}
			node.Red = false
			node.left.Red = true
			node.right.Red = true
		}
	}

	return node
}

// 插入的节点是左孩子的左子树（右旋）
func llRotate(node *RBTreeNode) *RBTreeNode {
	leftNode := node.left
	node.left = leftNode.right
	leftNode.right = node
	return leftNode
}

// 插入的节点是右孩子的右子树（左旋）
func rrRotate(node *RBTreeNode) *RBTreeNode {
	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node
	return rightNode
}

// 插入的节点是左孩子的右子树
func lrRotate(node *RBTreeNode) *RBTreeNode {
	node.left = rrRotate(node.left)
	return llRotate(node)
}

// 插入的节点是右孩子的左子树
func rlRotate(node *RBTreeNode) *RBTreeNode {
	node.right = llRotate(node.right)
	return rrRotate(node)
}
