---
id: interfaces_kogda_sravnenie_interfeysov_vyzovet
deck: Go
tags:
  - interfaces
---

# Front

Когда сравнение интерфейсов вызовет `panic`?

# Back

Когда оба интерфейса содержат значение несравнимого типа (`slice`, `map`, `function`).

Пример: два `any` со значениями `[]int{1}` — при `==` будет:

`panic: runtime error: comparing uncomparable type []int`.

Компилятор не может это проверить заранее, так как конкретный тип неизвестен.
