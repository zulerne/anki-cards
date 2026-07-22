---
id: panic-recover_chto_vyvedet_kod_x
deck: Go
tags:
  - panic-recover
---

# Front

Что выведет код?

```go
x := 0
defer fmt.Println(x) // defer с аргументом
x = 1
defer func() { fmt.Println(x) }() // defer с closure
x = 2
```

# Back

Выведет 2, затем 0. Closure захватывает переменную по ссылке — видит финальное значение 2. Аргумент `defer` вычисляется в момент объявления — фиксирует 0. Порядок: LIFO.
