// 二叉树广度优先遍历（BFS）
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先遍历，使用队列的数据结构
func Bfs(root *TreeNode) []int {
    order := []int{}
    if root == nil {
        return order
    }
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        // 队首出队
        head := queue[0]
        queue = queue[1:]
        // 输出元素值
        order = append(order, head.Val)
        // 孩子节点入队
        if head.Left != nil {
            queue = append(queue, head.Left)
        }
        if head.Right != nil {
            queue = append(queue, head.Right)
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
    order := Bfs(newTestTree())
    fmt.Println(order)
}
