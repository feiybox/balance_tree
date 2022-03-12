package two_three_tree

type TwoThreeTreeNode struct {
	Vals   []int
	parent *TwoThreeTreeNode
	nodes  []*TwoThreeTreeNode // 指向和val对应的子树的节点
}

func BSFPrint(root *TwoThreeTreeNode) []*TwoThreeTreeNode {
	nodeList := make([]*TwoThreeTreeNode, 0, 10)
	nodeList = append(nodeList, root)
	for i := 0; i < len(nodeList); i++ {
		n := nodeList[i]
		if n == nil {
			continue
		}
		nodeList = append(nodeList, n.nodes...)
	}
	return nodeList
}

// 查找2-3树节点
func Find(root *TwoThreeTreeNode, val int) *TwoThreeTreeNode {
	if root == nil {
		return nil
	}

	flag := 0
	for ; flag < len(root.Vals); flag++ {
		if root.Vals[flag] == val {
			return root
		}
		if root.Vals[flag] > val {
			// 处理2-节点，nodes为空case
			if len(root.nodes) <= flag {
				return nil
			}
			return Find(root.nodes[flag], val)
		}
	}
	// 处理2-节点，nodes为空case
	if len(root.nodes) <= flag {
		return nil
	}
	return Find(root.nodes[flag], val)
}

// 向2-3树中添加节点
func AddNode(root *TwoThreeTreeNode, val int) *TwoThreeTreeNode {
	if root == nil {
		return &TwoThreeTreeNode{
			Vals: []int{val},
		}
	}

	addNode(root, val)
	if len(root.Vals) < 3 {
		return root
	}

	left := &TwoThreeTreeNode{
		Vals: []int{root.Vals[0]},
	}
	right := &TwoThreeTreeNode{
		Vals: []int{root.Vals[2]},
	}
	if len(root.nodes) > 0 {
		left.nodes = root.nodes[:2]
		right.nodes = root.nodes[2:]
	}

	// 分裂根节点
	return &TwoThreeTreeNode{
		Vals:  []int{root.Vals[1]},
		nodes: []*TwoThreeTreeNode{left, right},
	}
}

func addNode(root *TwoThreeTreeNode, val int) {
	// 找到第一个比当前数大的数
	flag := 0
	for ; flag < len(root.Vals); flag++ {
		if root.Vals[flag] == val {
			return
		}
		if root.Vals[flag] > val {
			break
		}
	}

	// 找到第一个不需要往下查找的节点
	if len(root.nodes) == 0 {
		// 将当前数据插入到vals中
		root.Vals = append(root.Vals[:flag], append([]int{val}, root.Vals[flag:]...)...)
		return
	}

	// 继续向子节点查找查找
	addNode(root.nodes[flag], val)

	// 判断是否需要进行分裂，flag代表分裂节点
	needSplit := false
	splitFlag := 0
	for ; splitFlag < len(root.nodes); splitFlag++ {
		// 如果节点是4节点，则将中间数抽到这层
		if len(root.nodes[splitFlag].Vals) == 3 {
			needSplit = true
			break
		}
	}
	if needSplit {
		// 分裂节点
		sonNode := root.nodes[splitFlag]
		left := &TwoThreeTreeNode{Vals: []int{sonNode.Vals[0]}}
		right := &TwoThreeTreeNode{Vals: []int{sonNode.Vals[2]}}
		if len(sonNode.nodes) > 0 {
			left.nodes = sonNode.nodes[:2]
			right.nodes = sonNode.nodes[2:]
		}
		root.Vals = append(root.Vals[:splitFlag], append([]int{sonNode.Vals[1]}, root.Vals[splitFlag:]...)...)
		root.nodes = append(root.nodes[:splitFlag], append([]*TwoThreeTreeNode{left, right}, root.nodes[splitFlag+1:]...)...)
	}
	return
}

// 删除2-3树中的节点
func DelNode(root *TwoThreeTreeNode, val int) *TwoThreeTreeNode {
	// 暂未实现
	return nil
}
