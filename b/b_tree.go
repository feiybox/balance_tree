package b

// M B树的阶
const M int = 7

type BNode struct {
	Vals  []int
	nodes []*BNode
}

func BFSPrint(root *BNode) [][]*BNode {
	if root == nil {
		return nil
	}
	res := make([][]*BNode, 0, M+1)
	res = append(res, []*BNode{root})
	arr := root.nodes
	if arr == nil || len(arr) == 0 {
		return res
	}
	for len(arr) > 0 {
		res = append(res, arr)
		nextArr := make([]*BNode, 0, len(arr)*2)
		for _, node := range arr {
			if node.nodes == nil {
				continue
			}
			nextArr = append(nextArr, node.nodes...)
		}
		arr = nextArr
	}
	return res
}

// Find B树查找
func Find(root *BNode, val int) *BNode {
	i := 0
	for ; i < len(root.Vals) && val > root.Vals[i]; i++ {
	}
	if i < len(root.Vals) && val == root.Vals[i] {
		return root
	}
	if len(root.nodes) > 0 {
		return Find(root.nodes[i], val)
	}
	return nil
}

// AddNode 向B树中添加node
func AddNode(root *BNode, val int) *BNode {
	// 没有node时，则直接构建node
	if root == nil {
		return &BNode{
			Vals: []int{val},
		}
	}
	// 找到和val值最相近的叶节点
	root = addNode(root, val)
	if len(root.Vals) > M {
		// 分裂节点，并且层高+1
		flag := len(root.Vals) / 2
		var leftNodes, rightNodes []*BNode = nil, nil
		if root.nodes != nil && len(root.nodes) > 0 {
			leftNodes = root.nodes[:flag+1]
			rightNodes = root.nodes[flag+1:]
		}
		node := &BNode{
			Vals: []int{root.Vals[flag]},
			nodes: []*BNode{
				{
					Vals:  root.Vals[:flag],
					nodes: leftNodes,
				}, {
					Vals:  root.Vals[flag+1:],
					nodes: rightNodes,
				},
			},
		}
		root = node
	}
	return root
}

func addNode(root *BNode, val int) *BNode {
	if root == nil {
		return nil
	}

	i := 0
	for ; i < len(root.Vals) && val > root.Vals[i]; i++ {
	}
	if i < len(root.Vals) && val == root.Vals[i] {
		return root
	}
	if len(root.nodes) > 0 {
		root.nodes[i] = addNode(root.nodes[i], val)
		if len(root.nodes[i].Vals) > M {
			root = splitSubNode(root, i)
		}
		return root
	}
	// 将val加入到当前节点val里
	endVals := root.Vals[i:]
	root.Vals = append(append(append(make([]int, 0, len(root.Vals)+1), root.Vals[:i]...), val), endVals...)

	return root
}

// 分裂某个子节点
func splitSubNode(root *BNode, index int) *BNode {
	subNode := root.nodes[index]
	flag := len(subNode.Vals) / 2
	endVals := root.Vals[index:]
	endNodes := root.nodes[index+1:]
	var subLeftNodes, subRightNodes []*BNode = nil, nil
	if subNode.nodes != nil && len(subNode.nodes) > 0 {
		subLeftNodes = subNode.nodes[:flag+1]
		subRightNodes = subNode.nodes[flag+1:]
	}
	root.Vals = append(append(append(make([]int, 0, len(root.Vals)+1), root.Vals[:index]...), subNode.Vals[flag]), endVals...)
	root.nodes = append(append(append(make([]*BNode, 0, len(root.nodes)+1), root.nodes[:index]...), &BNode{
		Vals:  subNode.Vals[:flag],
		nodes: subLeftNodes,
	}, &BNode{
		Vals:  subNode.Vals[flag+1:],
		nodes: subRightNodes,
	}), endNodes...)
	return root
}

func DelNode(root *BNode, val int) *BNode {

	return nil
}
