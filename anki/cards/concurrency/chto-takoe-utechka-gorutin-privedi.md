---
id: concurrency_chto_takoe_utechka_gorutin
deck: Go
tags:
  - concurrency
---

# Front

Что такое утечка горутин? Приведи пример и способ предотвращения.

# Back

Горутина заблокирована навсегда (ждёт канал, который никто не закроет; бесконечный цикл без выхода). Потребляет память и стек (~2KB+).

Предотвращение: всегда обеспечить путь выхода — context с cancel/timeout, done-канал, `select` с `ctx.Done()`.

Обнаружение: `runtime.NumGoroutine()`, goleak в тестах.
