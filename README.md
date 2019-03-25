初めてのgo lang練習  
参考  
https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7

golangの特徴  
- コンパイル型の言語  
- シンプル  
- 型がある  
- 並行プログラミングに強いらしい  

基本知識
- 拡張子は .go  
- プログラムは必ず何かのpackageに属している必要があり、そのうちの1つは必ずmainでないといけない。パッケージが名前空間の役割を持つ  
- mainパッケージにmain関数があれば、それを実行する  
- セミコロンなし

実行  
- go build  
  - mainパッケージを読み込み、実行ファイルを作る(コンパイル)  
- go run <file>  
  - fileを実行する  
  
### 簡単なコード解説  
- 前提
  - 変数の定義
  - var val <type>  
  - val := <value> (var 省略)  
  - var a, b int  
  - a, b := 10, 14  
  - var ( c int \n d string )  

- test1.go, test2.go  
  - 関数の宣言方法と引数の取り方  
  - 関数もデータ型なので, `f := func()` の書き方ができる。  
  - 即時関数的な書き方(JSみたい) `func(msg string) {...} ("message")`  
  
- test3.go  
  - 配列の書き方  
  - golangの場合、要素数が固定になっている、値渡しが基本(参照渡しにすればいいのでは?)みたいな理由で配列を使うのはあんまりよくないっぽい  
  
- test4.go, test5.go  
  - スライス  
  - `a := [5]int{2,10,8,15,4}`  
  - `b := a[2:4]`  
  - `e := a[:]`  
  - 組み込み関数: `len, cap, make, append, copy`  
  - `make([]int, 3)`  
  - `append(s, 8, 2, 10)`  
  - `copy(参照先, 参照元)`  
  
- test7.go  
  - Map  
  - `m := make(map[string]int)`  
  - `m["fujimoto"] = 200`  
  - `m["arita"] = 300`  
  - `map[string][int]{"fujimoto": 100, "arita": 200}`  
  - `v, ok := m["fujimoto"]` vにvalue, okに存在の有無  
  - ブランク修飾子 `_` はなにが入ってきても廃棄してくれる。  `_, ok := m[key]`  

- test8.go, test9.go  
  - if and switch  
  - if文にそのスコープでのみ使用する変数を定義できる  
  
- test10.go  
  - 繰り返し `for`  
  - ex) `for i := 0; i < 10; i++ {}`  
  - while文はない  
  
- test11.go  
  - range, 配列、スライス、マップの要素分だけなんらかの処理を繰り返し行いたい時に使える命令  
  - `s := []int{2, 3, 8}`  `for i, v := range s {...}` iにはインデックスやキー, vには値が入る  

- test12.go  
  - 構造体の定義とフィールド
  - `type user struct {}`  
  
- test13.go  
  - メソッド, go langではクラスなどで定義されるメソッドはサポートしていない  
  - メソッドを構造体に紐づけるためにレシーバを使う  
  - `func (u user) show() {}`  

- test14.go  
  - メソッドの一覧を定義したデータ型  
  - それらのメソッドをある構造体が実装していればその構造体をそのインタフェース型として扱っても良いという特徴がある  
  - `type greeter interface { greet() }`  

- test15.go  
  - 空インターフェース。なんでも受け取れる型。TSの`any`?  
  - 型アサーションを使用して動的に型を知ることができる。  
  - `_, ok := t.(type)`, (型、bool)  

- test16.go test17.go
  - go routine, channel  
  - 並行プログラミングに使用するらしい。  
  
  
  

