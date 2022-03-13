package b

import (
	"math/rand"
	"testing"
)

func TestFind(t *testing.T) {
	root := AddNode(nil, 1)
	for i := 0; i < 10; i++ {
		root = AddNode(root, i)
	}

	node := Find(root, 9)
	println(node != nil)
}

func TestAddNode(t *testing.T) {
	root := AddNode(nil, 10)
	for i := 1; i < 200; i++ {
		root = AddNode(root, rand.Intn(i))
	}

	for _, nodes := range BFSPrint(root) {
		for _, node := range nodes {
			if node == nil {
				continue
			}
			for _, val := range node.Vals {
				print(val)
				print(",")
			}
			print(" ; ")
		}
		println()
	}
}
