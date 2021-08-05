package main

import (
	"fmt"
)

// appendの動き
// 	- capacityが足りないときは、およそ2倍のcapacityを確保する
// 	- capacityが増えるタイミングでポインタが変わる
func main() {
	s := []int{0}
	capacity := cap(s)

	fmt.Printf("init: pointer=%p len=%d cap=%d\n", s, len(s), cap(s))

	for i := 1; i <= (1 << 20); i++ {
		s = append(s, i)

		if capacity < cap(s) {
			capacity = cap(s)
			fmt.Printf("i=%6d: pointer=%p len=%6d cap=%7d\n", i, s, len(s), cap(s))
		}
	}
}
