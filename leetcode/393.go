// 原题链接：https://leetcode-cn.com/problems/utf-8-validation/
package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	// 思路：
	// 1) 首先比对字节前几个比特位确定UTF8字符字节数
	// 2）读取剩余字符判断前缀是否为10
	// 3）如果不匹配则失败退出
	num := len(data)
	valid := true
	for i := 0; i < num; i++ {
		if data[i] > 0xff {
			valid = false
			break
		}
		c := byte(data[i])
		// fmt.Printf("0x%x\n", c)
		n := 0
		if (c&0xf8)^0xf0 == 0 { // 四字节
			n = 3
		} else if (c&0xf0)^0xe0 == 0 { // 三字节
			n = 2
		} else if (c&0xe0)^0xc0 == 0 { // 双字节
			n = 1
		} else if (c & 0x80) == 0 { // 单字节
			continue
		} else {
			valid = false
			break
		}
		// fmt.Printf("i=%d, n=%d\n", i, n)
		if i+n >= num {
			valid = false
			break
		}

		// 连续读取n个字节，并判断是否以10开头
		for j := i + 1; j <= i+n; j++ {
			// fmt.Printf("0x%x\n", data[j])
			if data[j] > 0xff {
				valid = false
				break
			}
			if (byte(data[j])&0xc0)^0x80 != 0 {
				valid = false
				break
			}
		}
		// 字符无效直接退出循环
		if !valid {
			break
		}
		// 跳过n个字节
		i += n
	}
	return valid
}

// func IsValidUtf8(data []byte) bool {
// 	num := len(data)
// 	if num == 0 {
// 		return false
// 	}
// 	var n int
// 	var c byte
// 	var valid bool
// 	for i := 0; i < num; i++ {
// 		c = data[i]
// 		if (c&0xf8)^0xf0 == 0 { // 四字节
// 			n = 3
// 		} else if (c&0xf0)^0xe0 == 0 { // 三字节
// 			n = 2
// 		} else if (c&0xe0)^0xc0 == 0 { // 双字节
// 			n = 1
// 		} else if (c & 0x80) == 0 { // 单字节
// 			continue
// 		} else {
// 			valid = false
// 			break
// 		}
//
// 		if i+n >= num {
// 			valid = false
// 			break
// 		}
// 		// 连续读取n个字节，并判断是否以10开头
// 		for j := i + 1; j <= i+n; j++ {
// 			if (byte(data[j])&0xc0)^0x80 != 0 {
// 				valid = false
// 				break
// 			}
// 		}
// 		// 字符无效直接退出循环
// 		if !valid {
// 			break
// 		}
// 		// 跳过n个字节
// 		i += n
// 	}
// 	return valid
// }

func main() {
	data := []int{237}
	// data := []int{235, 140, 4}
	// data := []int{197, 130, 1}
	valid := validUtf8(data)
	fmt.Printf("Result: %+v\n", valid)
	// bytes := []byte("程序猿01号")
	// fmt.Printf("%+v\n", bytes)
	// fmt.Printf("Result: %+v\n", IsValidUtf8(bytes))
}
