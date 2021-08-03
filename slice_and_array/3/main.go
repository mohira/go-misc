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

	// 連番であることを確認する
	for i := range months {
		s := months[i:]
		fmt.Printf("&months[%02d:]=%p\n", i, s)
	}
}
