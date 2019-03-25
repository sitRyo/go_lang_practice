// 配列

package main

import "fmt"

func main() {
	var a [5]int

	a[2] = 3
	a[4] = 10

	fmt.Println(a)

	// 宣言と代入を同時にする
	b := [3]int{1, 3, 6}
	fmt.Println(b)

	c := [...]int{2, 4, 7, 5, 5}
	fmt.Println(c)
}
