package main

import "fmt"

func main() {
	underlyingArray := [...]int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	reverse := func(nums []int) {
		// 両サイドから入れ替えていくアルゴリズム
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	s := underlyingArray[:]
	reverse(s)

	// 当然、underlying arrayも書き換わっている
	fmt.Println(s)               // [5 6 2 9 5 1 4 1 3]
	fmt.Println(underlyingArray) // [5 6 2 9 5 1 4 1 3]
}
