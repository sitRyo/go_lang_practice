// スライス
package main

import "fmt"

func main() {
	a := [5]int{2, 10, 8, 15, 4}

	b := a[2:4]
	c := a[2:]
	d := a[:4]
	e := a[:]

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// スライスの長さを取得
	fmt.Println(len(c))

	// スライスの最大容量
	fmt.Println(cap(c))
}
