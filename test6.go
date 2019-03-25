// マップ
package main

import "fmt"

func main() {
	// mapを宣言する
	// m := make(map[string]int)
	// mapに値を設定
	// m["fujimoto"] = 200
	// m["arita"] = 300

	// 値を指定しながら宣言をする // key value
	m := map[string]int{"fujimoto": 100, "arita": 200}

	// キーの存在を調べる
	// if (ok == true) v <= value
	// else v <= 0
	v, ok := m["fujimoto"]

	fmt.Println(v)
	fmt.Println(ok)

	// 要素を削除する
	delete(m, "fujimoto")
	fmt.Println(m)
}
