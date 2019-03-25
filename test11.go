package main

import "fmt"

// 構造体を定義
type user struct {
	// field
	name  string
	score int
}

func main() {
	// newでuser構造体文の領域を確保して、初期化して、そのポインタを返す
	u1 := new(user)

	u1.name = "fujimoto"
	u1.score = 200

	fmt.Println(u1)

	// 初期化する時に直接値を与えることも可能
	// こちらの場合はポインタではない構造体のデータが入ってくる
	u2 := user{name: "arita", score: 300}
	fmt.Println(u2)
}
