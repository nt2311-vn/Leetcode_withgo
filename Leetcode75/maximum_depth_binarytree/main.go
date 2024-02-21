package main

import "fmt"

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rootVal := TreeNode{
		val:   3,
		Left:  &TreeNode{val: 9, Right: nil, Left: nil},
		Right: &TreeNode{val: 20, Right: &TreeNode{val: 7}, Left: &TreeNode{val: 15}},
	}

	fmt.Println(maxDepth(&rootVal))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxVal := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	return 1 + maxVal(maxDepth(root.Left), maxDepth(root.Right))
}
