---
id: panic-recover_kak_rabotaet_recover_gde
deck: Go
tags:
  - panic-recover
---

# Front

Как работает `recover`? Где его нужно вызывать?

# Back

`recover()` перехватывает `panic` и возвращает переданное значение. Работает только при вызове внутри `defer`-функции. Вне `defer` возвращает `nil`.

После `recover` выполнение продолжается после вызвавшей функции (не с точки `panic`).
