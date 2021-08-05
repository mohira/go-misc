package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i64 int64
	var i32 int32

	// バイトサイズを確認する
	fmt.Printf("i64 size=%d\n", unsafe.Sizeof(i64))
	fmt.Printf("i32 size=%d\n", unsafe.Sizeof(i32))
}
