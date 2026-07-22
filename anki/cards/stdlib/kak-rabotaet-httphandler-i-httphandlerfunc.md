---
id: stdlib_kak_rabotaet_httphandler_i
deck: Go
tags:
  - stdlib
---

# Front

Как работает `http.Handler` и `http.HandlerFunc`? В чём разница?

# Back

`http.Handler` — интерфейс с методом `ServeHTTP(w, r)`.

`http.HandlerFunc` — тип-адаптер: `type HandlerFunc func(w, r)` с методом `ServeHTTP`, который вызывает саму функцию.

Позволяет использовать обычную функцию как `Handler` без создания структуры.
