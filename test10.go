package main

import "fmt"

func main() {
	// スライス
	s := []int{2, 3, 8}

	for i, v := range s {
		// i -> index, v -> value
		v = 10
		fmt.Println(i, v)
	}

	// rangeが返すvはイテレータ？

	// ブランク修飾子
	// _にしておくと何が入ってきてもそれを廃棄してくれるという仕様
	for _, v := range s {
		fmt.Println(v)
	}

	// マップ
	m := map[string]int{"fujimoto": 200, "arita": 300}

	for k, v := range m {
		// key, value
		fmt.Println(k, v)
	}
}
