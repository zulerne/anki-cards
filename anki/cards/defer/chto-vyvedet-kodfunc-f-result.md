---
id: defer_chto_vyvedet_kodfunc_f
deck: Go
tags:
  - defer
---

# Front

Что выведет код?func f() (result int) {
  defer func() { result++ }()
  return 0
}

# Back

1. Defer с closure может изменять именованные return values. Порядок: result = 0 (return) → defer меняет result на 1 → функция возвращает 1.
