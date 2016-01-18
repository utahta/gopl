# 4. Composite Types

* array, slice, map, struct
  * array, struct は集成体型
  * array は同じ型、structs は色々な型
  * array, struct 固定
  * slice, map 可変
* JSON
* generate HTML from template


## 4.1 Arrays

* あんまり使わないらしい（slice 便利）
* `a := [...]int{1, 2, 3}` と書ける
* `a := [...]int{99: -1}` と書けるし, 100個要素が定義される
* `[2]int == [2]int` 型が同じなら比較できる

## 4.2 Slices

* pointer, length, capacity
* `s := array[i:j]` s は array を参照
* slice 同士の == 比較はできない。自前で実装する必要ある（けど止めた方がいいらしい）
* `len(a) == 0` slice が空か判断するとき

### 4.2.1 The append Function

* `s = append(s, "a", "b", "c")`
* capacity を超えたら、新しく capacity を増やした array が割り当てられる
* append 前と後で異なる array を参照してるかもしれない

### 4.2.2 In-Place Slice Techniques

* 用例

## 4.3 Maps

* 存在しない key にアクセスしても安全。zero value 返る
* `_ = &age['hoge']` 値のアドレスとれない
* range で取れる順番は不定
* `var ages[string]int`
  * `ages == nil // true`
  * `len(ages) == 0 // true`
* 存在確認
  * `age, ok := ages["hoge"]`
* `==` で比較できない

## 4.4 Structs

* Field
  * 定義の順番大事
  * 順番変えると別の構造体の定義
  * 先頭大文字で外からアクセス可
  * 自分自身を Field に定義できない（ポインタならOK）

### 4.4.1 Struct Literals

* `type Point struct {x, y int}`
  * `_ := Point{1, 2}`
  * `_ := Point{x: 1, y: 2}`
* `func call(p Point) {...}` 値渡し
* `func call(p *Point) {...}` 参照渡し

### 4.4.2 Comparing Structs

* `==` `!=` 使える
* map の key できる

### 4.4.3 Struct Embedding and Anonymous Fields

* struct を埋め込むことで `.` 呼び出しを省略できる
* embed できるのは struct に限らない
  * method ?

## 4.5 JSON

* `encoding/json` 便利
  * `json.Marshal`
  * `json.Unmarshal`
  * `json.NewDecoder(io.Reader).Decode(&result)`
  * `json.NewEncoder(io.Writter).Encode(&result)`

## 4.6 Text and HTML Templates






