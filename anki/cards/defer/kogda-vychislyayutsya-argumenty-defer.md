---
id: defer_kogda_vychislyayutsya_argumenty_defer
deck: Go
tags:
  - defer
---

# Front

Когда вычисляются аргументы defer?

# Back

В момент объявления defer, НЕ в момент выполнения. defer fmt.Println(x) — значение x фиксируется сразу.Замыкания (defer func()) читают переменные при выполнении
