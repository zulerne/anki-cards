---
id: concurrency_kak_realizovat_graceful_shutdown
deck: Go
tags:
  - concurrency
---

# Front

Как реализовать graceful shutdown HTTP-сервера в Go?

# Back

1) Запустить `srv.ListenAndServe()` в горутине.

2) Ждать сигнал ОС (`signal.NotifyContext` или `signal.Notify` с `SIGINT`/`SIGTERM`).

3) Вызвать `srv.Shutdown(ctx)` — перестаёт принимать новые соединения, дожидается завершения текущих в пределах таймаута контекста.
