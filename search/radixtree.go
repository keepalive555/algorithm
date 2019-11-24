// RadixTree实现
package main

import (
	"fmt"
)

type NodeType int

const (
	_ NodeType = iota
	Static
)

type Node struct {
	NodeType  NodeType    // 节点类型
	Path      string      // 路径
	Indicies  []byte      // 索引
	Childrens []*Node     // 节点
	Value     interface{} // 节点值
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	var i int
	n := min(len(a), len(b))
	for i = 0; i < n; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

func (node *Node) findLast(path string) (*Node, string) {
	current := node
	for {
		// 求当前节点Path与待查找Path的最长公共前缀
		i := longestCommonPrefix(node.Path, path)
		// 若无公共前缀
		if i == 0 {
			break
		}
	}
	return current, path
}

func main() {
	a := "/searchbox/user"
	b := "/searchbox/profile"
	i := longestCommonPrefix(a, b)
	fmt.Printf("prefix=%s\n", a[:i])
}
