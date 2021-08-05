package main

import (
	"fmt"
)

// スライスのメモリ上の配置のイメージをつかむ
func main() {
	// 8Byteずつずれる(int64で、len=1, capacity=1だから)
	fmt.Printf("pointer=%p\n", make([]int, 1)) // 0xc0000ae010
	fmt.Printf("pointer=%p\n", make([]int, 1)) // 0xc0000ae018
	fmt.Printf("pointer=%p\n", make([]int, 1)) // 0xc0000ae020

	// 16Byteずつずれる(int64で、len=2, capacity=2だから)
	fmt.Printf("pointer=%p\n", make([]int, 2)) // 0xc0000ae030
	fmt.Printf("pointer=%p\n", make([]int, 2)) // 0xc0000ae040
	fmt.Printf("pointer=%p\n", make([]int, 2)) // 0xc0000ae050

	// 4Byteずつずれる(int32で、len=1, capacity=1だから)
	fmt.Printf("pointer=%p\n", make([]int32, 1)) // 0xc0000ae028
	fmt.Printf("pointer=%p\n", make([]int32, 1)) // 0xc0000ae02c
	fmt.Printf("pointer=%p\n", make([]int32, 1)) // 0xc0000ae030

	// string: 16Byteずつずれる
	fmt.Printf("pointer=%p\n", make([]string, 1)) // 0xc000096220
	fmt.Printf("pointer=%p\n", make([]string, 1)) // 0xc000096230
	fmt.Printf("pointer=%p\n", make([]string, 1)) // 0xc000096240
}
