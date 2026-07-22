---
id: context_dlya_chego_nuzhen_contextcontext
deck: Go
tags:
  - context
---

# Front

Для чего нужен `context.Context`?

# Back

Передача дедлайнов, сигналов отмены и request-scoped значений через цепочку вызовов.

Передаётся первым параметром функции: `func Do(ctx context.Context, ...)`.

Основной механизм graceful cancellation в Go — горутины слушают `ctx.Done()` и завершаются.
