// 関数の書き方練習2 即時関数など

package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(swap(5, 2))

	f := func(a int, b int) (int, int) {
		return b, a
	}

	fmt.Println(f(3, 8))

	// 即時関数的な(JSぽい)
	func(msg string) {
		fmt.Println(msg)
	}("seriru")
}
