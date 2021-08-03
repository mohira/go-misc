package main

import "fmt"

func main() {
	months := [13]string{
		"", // 添字と対応させるための番兵
		"1月", "2月", "3月",
		"4月", "5月", "6月",
		"7月", "8月", "9月",
		"10月", "11月", "12月",
	}

	// 開始地点が同じなら、同じポインタだよね
	s1 := months[7:13]
	s2 := months[7:10]
	fmt.Printf("s1=%p\n", s1)
	fmt.Printf("s2=%p\n", s2)
}
