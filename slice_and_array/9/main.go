package main

import (
	"fmt"
)

// メモリ上の配置のイメージをつかむ
// Q. make([]int, 0)を使うと初期のポインタは同じになるのはなぜ？？
func main() {
	fmt.Printf("pointer=%p\n", make([]int, 0)) // 0x118e370
	fmt.Printf("pointer=%p\n", make([]int, 0)) // 0x118e370
	fmt.Printf("pointer=%p\n", make([]int, 0)) // 0x118e370
}
