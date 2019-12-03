// 原题链接：https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/submissions/
package main

import (
    "fmt"
)

func countPrimeSetBits(L int, R int) int {
    // 1)思路一：输入数据类型为int类型，比特位个数为：[0, 32]（符号位计入在内），此区间内质数可以枚举。
    // 使用for遍历[L, R]，求出每个数的置位个数，并查表判断是否为质数，即可获得答案。
    if L > R {
        return 0
    }
    // 质数表
    table := map[int]bool{
        2: true,
        3: true,
        5: true,
        7: true,
        11: true,
        13: true,
        17: true,
        19: true,
        23: true,
        29: true,
        31: true,
    }
    var total int

    for i:=L; i<=R; i++ {
        count := 0
        n := i
        for n > 0 {
            count++
            n = n & (n - 1)
        }
        fmt.Println(count)

        if _, ok := table[count]; ok {
            total++
        }
    }
    return total
}

func main() {
    r := countPrimeSetBits(842, 888)
    fmt.Println(r)
}
