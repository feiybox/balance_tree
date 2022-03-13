package red_black_tree

type RBTreeNode struct {
	Val    int
	Red    bool
	left   *RBTreeNode
	right  *RBTreeNode
	parent *RBTreeNode
}

// Find 查找树上的节点
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

// AddNode 增加节点
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
		node.left.parent = node
		son = node.left
		insertLeft = true
	} else {
		node.right = addNode(node.right, val)
		node.right.parent = node
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
	if leftNode == nil {
		return node
	}
	leftNode.parent = node.parent
	node.left = leftNode.right
	if leftNode.right != nil {
		leftNode.right.parent = node
	}
	leftNode.right = node
	node.parent = leftNode
	return leftNode
}

// 插入的节点是右孩子的右子树（左旋）
func rrRotate(node *RBTreeNode) *RBTreeNode {
	rightNode := node.right
	if rightNode == nil {
		return node
	}
	rightNode.parent = node.parent
	node.right = rightNode.left
	if rightNode.left != nil {
		rightNode.left.parent = node
	}
	rightNode.left = node
	node.parent = rightNode
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

func DelNode(root *RBTreeNode, val int) *RBTreeNode {
	root, _ = delNode(root, val)
	if root != nil && root.Red {
		root.Red = false
	}
	return root
}

// 删除对应数据的节点，返回最新二叉树子树节点，与是否需要重平衡
func delNode(root *RBTreeNode, val int) (*RBTreeNode, bool) {
	if root == nil {
		return nil, false
	}

	var depRe bool
	if root.Val > val {
		// 删除节点比当前节点值小
		root.left, depRe = delNode(root.left, val)
		if depRe {
			root, depRe = reBalance(root, true)
		}
		return root, depRe
	} else if root.Val < val {
		// 删除节点比当前节点值大
		root.right, depRe = delNode(root.right, val)
		if depRe {
			root, depRe = reBalance(root, false)
		}
		return root, depRe
	}

	// 删除节点无子节点
	if root.left == nil && root.right == nil {
		// 删除节点为红色
		if root.Red {
			return nil, false
		} else {
			// 删除节点为黑色
			// 如果删除节点为根节点，直接删除
			if root.parent == nil {
				return nil, false
			}
			// 删除节点是黑色节点，需要重新平衡
			return nil, true
		}
	} else if root.left != nil && root.right != nil {
		// 删除节点左右子节点存在
		rightTop := findMinVal(root.right)
		root.Val, rightTop.Val = rightTop.Val, root.Val
		root.right, depRe = delNode(root.right, val)
		if depRe {
			root, depRe = reBalance(root, false)
		}
		return root, depRe
	} else {
		// 删除节点有一个子节点
		if root.left != nil {
			root.left.Red = false
			root.left.parent = root.parent
			root = root.left
			return root, false
		} else if root.right != nil {
			root.right.parent = root.parent
			root.right.Red = false
			root = root.right
			return root, false
		}
		return root, false
	}
}

// 再平衡
func reBalance(root *RBTreeNode, delInLeft bool) (*RBTreeNode, bool) {
	brother := root.left
	if delInLeft {
		brother = root.right
	}

	if brother == nil {
		return root, false
	}

	if brother.Red {
		// 兄弟节点为红色
		if !delInLeft {
			// 兄在左
			root.Red, brother.Red = brother.Red, root.Red
			root = llRotate(root)
			brother = root.left
			// 进入兄节点为黑色的流程
		} else {
			// 兄在右
			root.Red, brother.Red = brother.Red, root.Red
			root = rrRotate(root)
			brother = root.right
			// 进入兄节点为黑色的流程
		}
	}

	// 兄弟节点为黑色
	if !brother.Red {
		// 兄弟节点的左右节点都是黑色
		if (brother.left == nil || !brother.left.Red) && (brother.right == nil || !brother.right.Red) {
			if root.Red {
				brother.Red = true
				root.Red = false
			} else {
				brother.Red = true
				// 递归重新平衡
				return root, true
			}
		} else {
			// 兄在左，兄左子红
			if !delInLeft && brother.left != nil && brother.left.Red {
				root.Red, brother.Red = brother.Red, root.Red
				brother.left.Red = false
				return llRotate(root), false
			} else if delInLeft && brother.right != nil && brother.right.Red {
				// 兄在右，兄右子红
				root.Red, brother.Red = brother.Red, root.Red
				brother.right.Red = false
				return rrRotate(root), false
			} else if !delInLeft && brother.right != nil && brother.right.Red {
				// 兄在左，兄右子红
				brother.Red, brother.right.Red = brother.right.Red, brother.Red
				brother = rrRotate(brother)
				root.left = brother

				// 处理：兄在左，兄左子红
				root.Red, brother.Red = brother.Red, root.Red
				return llRotate(root), false
			} else if delInLeft && brother.left != nil && brother.left.Red {
				// 兄在右，兄左子红
				brother.Red, brother.left.Red = brother.left.Red, brother.Red
				root.right = llRotate(brother)
				brother = root.right

				// 处理：兄在右，兄右子红
				return rrRotate(root), false
			}
		}
	}
	return root, false
}

// 找到树中val最小的节点
func findMinVal(root *RBTreeNode) *RBTreeNode {
	if root == nil {
		return nil
	}
	min := findMinVal(root.left)
	if min != nil {
		return min
	}
	return root
}
