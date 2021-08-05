package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [...]int64{3, 1, 4}
	s := a[:]

	fmt.Printf("len=%d cap=%d size=%d a=%v\n", len(a), cap(a), unsafe.Sizeof(a), a)
	fmt.Printf("len=%d cap=%d size=%d s=%v\n", len(s), cap(s), unsafe.Sizeof(s), s)

	// sliceは3つの情報(underlying arrayのポインタ, len, cap)しか持たないので、
	// どれだけappendしようが、24Byte(==3 * 8Byte)で変わらない
	s = append(s, 1)
	s = append(s, 5)
	s = append(s, 9)
	s = append(s, 2)
	fmt.Printf("len=%d cap=%d size=%d s=%v\n", len(s), cap(s), unsafe.Sizeof(s), s)
}
