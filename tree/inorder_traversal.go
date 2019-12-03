// 二叉树中序遍历
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历，先打印左节点值 => 中节点值 => 右节点值
func InorderTraversal(root *TreeNode) []int {
	order := []int{}
	if root == nil {
		return order
	}
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		// 遍历左子树直至左子树节点为nil，遍历的节点全部使用stack保存
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// 栈不为空则取栈顶元素
		for len(stack) > 0 {
			// 元素出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 输出节点值
			order = append(order, top.Val)

			// 当前节点存在右子树，则继续入栈处理右子树
			if top.Right != nil {
				node = top.Right
				break
			}
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
	order := InorderTraversal(newTestTree())
	fmt.Print(order)
}
