// 二叉树后序遍历
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	order := []int{}
	if root == nil {
		return order
	}
	var lastVisited *TreeNode
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		// 遍历左子树直至左子树节点为nil，遍历的节点全部使用stack保存
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// Peek栈顶元素（不出栈）
		node = stack[len(stack)-1]

		// 右子树不存在或已被访问过则输出元素，并出栈
		if node.Right == nil || node.Right == lastVisited {
			order = append(order, node.Val)
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]
			lastVisited = node
			node = nil
		} else {
			node = node.Right
		}
	}
	return order
}

// TODO: 需要根据数组建立二叉树的函数
func newTestTree() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	left := &TreeNode{
		Val: 2,
	}
	right := &TreeNode{
		Val: 3,
	}
	root.Left = left
	root.Right = right
	leftLeft := &TreeNode{
		Val: 4,
	}
	left.Left = leftLeft
	rightRight := &TreeNode{
		Val: 5,
	}
	right.Right = rightRight
	leftLeft.Right = &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	return root
}

func main() {
	order := PostorderTraversal(newTestTree())
	fmt.Print(order)
}
