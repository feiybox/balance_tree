package AVL

type TreeNode struct {
	Val   int
	Dep   int
	left  *TreeNode
	right *TreeNode
}

// Find 找到某个数
// 找到则返回对应的节点，如果找不到则返回nil
func Find(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	if val < root.Val {
		return Find(root.left, val)
	}
	return Find(root.right, val)
}

// GetHeight 获取二叉树的高度
func GetHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Dep
}

// BFSPrint 广度优先遍历
func BFSPrint(root *TreeNode) []*TreeNode {
	nodeList := make([]*TreeNode, 0, 10)
	nodeList = append(nodeList, root)
	var i int
	for ; i < len(nodeList); i++ {
		n := nodeList[i]
		if n == nil {
			continue
		}
		nodeList = append(nodeList, n.left)
		nodeList = append(nodeList, n.right)
	}
	return nodeList
}

// AddNode 将一个节点加入到AVL树中
// 返回新的跟节点
func AddNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
			Dep: 1,
		}
	}

	// 重复节点，忽略
	if val == root.Val {
		return root
	}

	if val < root.Val {
		root.left = AddNode(root.left, val)
	} else {
		root.right = AddNode(root.right, val)
	}

	// 再平衡
	root = balance(root)
	updateNodeDep(root)

	return root
}

// 使二叉树平衡
func balance(root *TreeNode) *TreeNode {
	if root == nil || root.Dep <= 2 {
		return root
	}
	// 判断是否需要再进行平衡
	l := GetHeight(root.left)
	r := GetHeight(root.right)
	if l > r+1 {
		// 左边高度大于右边+1，则需要右旋
		ll := GetHeight(root.left.left)
		lr := GetHeight(root.left.right)
		if ll > lr {
			return llRotate(root)
		}
		return lrRotate(root)
	} else if r > l+1 {
		// 右边高度大于左边+1，则需要右旋
		rl := GetHeight(root.right.left)
		rr := GetHeight(root.right.right)
		if rl > rr {
			return rlRotate(root)
		}
		return rrRotate(root)
	}
	return root
}

func updateNodeDep(root *TreeNode) {
	var dep int
	if root.left != nil {
		dep = root.left.Dep
	}
	if root.right != nil && dep < root.right.Dep {
		dep = root.right.Dep
	}
	root.Dep = dep + 1
}

// 插入的节点是左孩子的左子树（右旋）
func llRotate(node *TreeNode) *TreeNode {
	leftNode := node.left
	node.left = leftNode.right
	leftNode.right = node
	return leftNode
}

// 插入的节点是右孩子的右子树（左旋）
func rrRotate(node *TreeNode) *TreeNode {
	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node
	return rightNode
}

// 插入的节点是左孩子的右子树
func lrRotate(node *TreeNode) *TreeNode {
	node.left = rrRotate(node.left)
	return llRotate(node)
}

// 插入的节点是右孩子的左子树
func rlRotate(node *TreeNode) *TreeNode {
	node.right = llRotate(node.right)
	return rrRotate(node)
}

// DelNode 删除对应数据的节点，返回最新二叉树的根
func DelNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		if root.left != nil && root.right != nil {
			// 将左子树中的最大的数来替换当前节点，因此不需要在删除的时候进行再平衡
			if GetHeight(root.left) > GetHeight(root.right) {
				leftMax := findMaxVal(root.left)
				root.left = DelNode(root.left, leftMax.Val)
				root.Val = leftMax.Val
			} else {
				rightMin := findMinVal(root.right)
				root.right = DelNode(root.right, rightMin.Val)
				root.Val = rightMin.Val
			}
			updateNodeDep(root)
			return root
		} else if root.left != nil {
			// 删除当前节点
			return root.left
		} else {
			return root.right
		}
	}

	if root.Val > val {
		root.left = DelNode(root.left, val)
	} else {
		root.right = DelNode(root.right, val)
	}

	// 再平衡
	root = balance(root)
	updateNodeDep(root)

	return root
}

// 找到树中val最大的节点
func findMaxVal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	max := findMaxVal(root.right)
	if max != nil {
		return max
	}
	return root
}

// 找到树中val最小的节点
func findMinVal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	min := findMinVal(root.left)
	if min != nil {
		return min
	}
	return root
}
