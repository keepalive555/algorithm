// Google Varint编码
// 参考代码：
// 编码：https://github.com/golang/protobuf/blob/master/proto/encode.go
// 解码：https://github.com/golang/protobuf/blob/master/proto/decode.go
package main

import (
	"fmt"
)

const maxVarintBytes = 10

// @param x uint64: 数字
// @return: Varint编码之后的字节流
func EncodeVarint(x uint64) []byte {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7f)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[:n]
}

func DecodeVarint(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n > len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		x |= (b & 0x7f) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}
	return 0, 0
}

func main() {
	var x uint64 = 1024
	b := EncodeVarint(x)
	fmt.Printf("encode varint <%d>, result => len <%d>, data <%+v>\n", x, len(b), b)
	n, _ := DecodeVarint(b)
	fmt.Printf("decode varint: %d\n", n)
}
