---
id: stdlib_chto_izmenilos_v_servemux_go_1_22
deck: Go
tags:
  - http
  - stdlib
---

# Front

Что изменилось в `http.ServeMux` в Go 1.22?

# Back

В Go 1.22 стандартный `ServeMux` получил более выразительные patterns с
методом, путём и wildcard, например `GET /items/{id}`.

Значение wildcard можно получить через `r.PathValue("id")`. Это встроенный
вариант маршрутизации; для более сложных требований можно использовать
сторонний router.
