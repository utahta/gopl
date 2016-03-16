# 7.8 The error Interface

```
type error Interface {
  Error() string
}
```

- `errors.New("hoge")` で生成できる。あんまり使わない。
- `fmt.Errorf("hoge")` わりと使う。
- `syscall.Errno(2)` 例を示してる
    - `syscall.Syscall` の実装は一体どこに？

# 7.9 Example: Expression Evaluator

簡単な算術式を評価するものをつくる。

- 二項演算子 `+` `-` `*` `/`
- 単項演算子 `+x` `-x`
- 関数呼び出し `pow(x, y)` `sin(x)` `sqrt(x)` 
- 変数
- 値は float64 型

gopl.ioサンプルコード

- 全容
    - https://github.com/adonovan/gopl.io/blob/master/ch7/eval/
- 定義 
    - https://github.com/adonovan/gopl.io/blob/master/ch7/eval/ast.go
- テスト
    - https://github.com/adonovan/gopl.io/blob/master/ch7/eval/eval_test.go
- Parse
    - https://github.com/adonovan/gopl.io/blob/master/ch7/eval/parse.go
- Check
    - https://github.com/adonovan/gopl.io/blob/master/ch7/eval/check.go
    - https://github.com/adonovan/gopl.io/blob/master/ch7/surface/surface.go
        
# 7.10 Type Assertions

C++ の dynamic cast みたいな感じがある

https://github.com/utahta/gopl/blob/master/ch7/type_assertions.go
        
# ズンドコ

https://gist.github.com/utahta/69f91fbb3d6d18ea392a
