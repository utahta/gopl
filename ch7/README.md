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
