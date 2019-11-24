package main

import ()

const (
	defaultSlotsCount = 64
)

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type HashTable struct {
	Slots []*Node
}

type Map struct {
	ht             []*HashTable
	rehashing      bool
	rehashingIndex int
}
