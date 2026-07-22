---
id: panic-recover_kogda_dopustimo_ispolzovat_panic
deck: Go
tags:
  - panic-recover
---

# Front

Когда допустимо использовать `panic`?

# Back

1) При инициализации программы, если невозможно продолжать (не открылась БД, невалидный конфиг).

2) В must-функциях (`regexp.MustCompile`, `template.Must`) — compile-time known данные.

3) Логические ошибки программиста (неправильное использование API).

В обычном бизнес-коде `panic` не допустим — используй `error`.
