package main

import "fmt"

func main() {
	underlyingArray := [...]int{3, 1, 4, 1, 5}

	s := underlyingArray[:]

	// スライスの要素を書き換える
	for i := 0; i < len(s); i++ {
		s[i] = 9
	}

	// underlying array も書き換わっている！(というか、underlying arrayを書き換えた)
	fmt.Println(s)               // [9 9 9 9 9]
	fmt.Println(underlyingArray) // [9 9 9 9 9]
}
