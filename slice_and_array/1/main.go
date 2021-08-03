package main

import "fmt"

func main() {
	// months := [...]string のほうが楽な記述
	months := [13]string{
		"", // 添字と対応させるための番兵
		"1月", "2月", "3月",
		"4月", "5月", "6月",
		"7月", "8月", "9月",
		"10月", "11月", "12月",
	}

	summer := months[6:9]
	q2 := months[4:7] // 第2四半期

	// それぞれの構成要素に着目(ポインタも違うよね)
	fmt.Printf("summer=%v, pointer=%p len=%d cap=%d\n", summer, summer, len(summer), cap(summer))
	fmt.Printf("    q2=%v, pointer=%p len=%d cap=%d\n", q2, q2, len(q2), cap(q2))
}
