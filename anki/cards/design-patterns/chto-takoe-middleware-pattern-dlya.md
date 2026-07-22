---
id: design-patterns_chto_takoe_middleware_pattern
deck: Go
tags:
  - design-patterns
---

# Front

Что такое middleware pattern для HTTP?

# Back

Обёртка вокруг handler, добавляющая cross-cutting логику (логирование, аутентификация, CORS, rate limiting).

Сигнатура: `func(http.Handler) http.Handler`.

Цепочка middleware: `logging(auth(handler))`. Каждый middleware решает — вызвать следующий или прервать.
