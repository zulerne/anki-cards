---
id: types_chto_takoe_iota_privedi
deck: Go
tags:
  - types
---

# Front

Что такое `iota`? Приведи пример использования в `const`.

# Back

Генератор последовательных целых чисел в блоке `const`. Начинается с 0, увеличивается на 1 для каждой строки.

```go
const (
    Read  = 1 << iota // 1
    Write             // 2
    Exec              // 4
)
```
