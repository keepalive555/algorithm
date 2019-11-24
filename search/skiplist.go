package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	MaxLevel int = 32 // 最大层高
)

var (
	P float32 = 0.25 // 概率（即：掷4次骰子，只有1次机会可以使level++）
)

var (
	InvalidArguments = errors.New(`invalid arguments`)
)

type KeyType = int   // 键类型
type ValueType = int // 值类型

// 节点值
type Node struct {
	Key     KeyType
	Value   ValueType
	Forward []*Node // 前进指针
}

// 创建节点
func createNode(key KeyType, value ValueType, level int) (*Node, error) {
	if level > MaxLevel {
		return nil, InvalidArguments
	}
	node := &Node{
		Key:     key,
		Value:   value,
		Forward: make([]*Node, level, level),
	}
	return node, nil
}

// SkipList数据结构的精髓，节点Level由randomLevel()随机生成，最大MaxLevel层。
// 此函数，经由大数据样本测试，从试验效果验证了函数符合预期。
func randomLevel() int {
	level := 1
	for (rand.Int() & 0xffff) < int(P*0xffff) {
		level++
		if level > MaxLevel {
			level = MaxLevel
			break
		}
	}
	return level
}

// SkipList数据结构定义
type SkipList struct {
	head  *Node // 头指针
	level int   // 最大Level，[1, MaxLevel]
	size  int   // 元素个数
}

func (skipList *SkipList) get(key KeyType) *Node {
	if skipList == nil || skipList.size == 0 {
		return nil
	}
	node := skipList.head
	for i := skipList.level - 1; i >= 0; i-- {
		for node.Forward[i] != nil && node.Forward[i].Key < key {
			node = node.Forward[i]
		}
		if node.Forward[i] != nil && node.Forward[i].Key == key {
			return node.Forward[i]
		} else if i == 0 { // 遍历至最后一层
			return nil
		}
	}
	return nil
}

func (skipList *SkipList) Get(key KeyType) (ValueType, bool) {
	node := skipList.get(key)
	if node != nil {
		return node.Value, true
	}
	return 0, false
}

// Golang风格的迭代器
func (skipList *SkipList) Range(callback func(KeyType, ValueType)) {
	// 临界检查
	if skipList == nil || skipList.size == 0 {
		return
	}
	// 遍历链表（跳过首节点）
	node := skipList.head.Forward[0]
	for node != nil {
		callback(node.Key, node.Value)
		node = node.Forward[0]
	}
}

// 删除节点
func (skipList *SkipList) Delete(key KeyType) {
	// 从上至下、从左至右遍历节点
	node := skipList.head
	last := make([]*Node, MaxLevel, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		last[i] = skipList.head
	}
	for i := skipList.level - 1; i >= 0; i-- {
		for node.Forward[i] != nil && node.Forward[i].Key < key {
			node = node.Forward[i]
		}
		last[i] = node
	}
	next := node.Forward[0]
	if next != nil {
		// 删除节点
		for i := 0; i < MaxLevel; i++ {
			if last[i].Forward[i] == next {
				last[i].Forward[i] = next.Forward[i]
			}
		}
		next = nil
		// 校正跳表level
		for skipList.level > 0 && skipList.head.Forward[skipList.level-1] == nil {
			skipList.level--
		}
		skipList.size--
	}
}

func (skipList *SkipList) Set(key KeyType, value ValueType) {
	// 从上至下、从左至右遍历节点
	node := skipList.head
	last := make([]*Node, MaxLevel, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		last[i] = skipList.head
	}
	for i := skipList.level - 1; i >= 0; i-- {
		for node.Forward[i] != nil && node.Forward[i].Key < key {
			node = node.Forward[i]
		}
		// 更新节点
		if node.Forward[i] != nil && node.Forward[i].Key == key {
			node.Forward[i].Value = value
			return
		}
		last[i] = node
	}
	// 新建level并校正跳表level
	level := randomLevel()
	if level > skipList.level {
		for j := skipList.level; j < level; j++ {
			last[j] = skipList.head
		}
		skipList.level = level
	}
	// 创建节点
	_new, _ := createNode(key, value, level)
	for i := 0; i < level; i++ {
		_new.Forward[i] = last[i].Forward[i]
		last[i].Forward[i] = _new
	}
	skipList.size++
}

func NewSkipList() *SkipList {
	node, _ := createNode(0, 0, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		node.Forward[i] = nil
	}
	skipList := &SkipList{
		head:  node,
		level: 1,
		size:  0,
	}
	return skipList
}

func main() {
	skipList := NewSkipList()
	skipList.Set(1, 100)
	skipList.Set(2, 200)
	skipList.Set(8, 800)
	skipList.Set(7, 700)
	skipList.Set(3, 300)
	skipList.Set(4, 400)
	// 获取Key
	v, ok := skipList.Get(8)
	fmt.Printf("<key:%d>=<%d, %v>\n", 8, v, ok)
	// 删除Key
	skipList.Delete(8)
	v, ok = skipList.Get(8)
	fmt.Printf("<key:%d>=<%d, %v>\n", 8, v, ok)
	// 遍历跳表
	skipList.Range(func(key KeyType, value ValueType) {
		fmt.Printf("key=%d, value=%d\n", key, value)
	})
}
