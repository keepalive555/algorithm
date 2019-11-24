package main

import (
	"fmt"
)

const (
	Mask  = 0xffffffff
	Shift = 5
)

type Bitmap struct {
	max    uint32
	bitmap []uint32
}

func NewBitmap(max uint32) *Bitmap {
	n := max/32 + 1
	bitmap := &Bitmap{
		max:    max,
		bitmap: make([]uint32, n, n),
	}
	return bitmap
}

func (bitmap *Bitmap) Set(n uint32) {
	if n > bitmap.max {
		panic(fmt.Errorf("out of maximum"))
	}
	bitmap.bitmap[n>>Shift] |= 1 << (n & Mask)
}

func (bitmap *Bitmap) Clr(n uint32) {
	bitmap.bitmap[n>>Shift] &^= 1 << (n & Mask)
}

func (bitmap *Bitmap) Test(n uint32) bool {
	i := n >> Shift
	if bitmap.bitmap[i]&(1<<(n&Mask)) == 1 {
		return true
	}
	return false
}

func main() {
	bitmap := NewBitmap(1024)
	bitmap.Set(100)

	fmt.Printf("test result: %+v\n", bitmap.Test(100))
	bitmap.Clr(100)
	fmt.Printf("test result: %+v\n", bitmap.Test(100))
}
