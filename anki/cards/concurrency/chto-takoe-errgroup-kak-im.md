---
id: concurrency_chto_takoe_errgroup_kak
deck: Go
tags:
  - concurrency
---

# Front

Что такое errgroup? Как им пользоваться?

# Back

Пакет `golang.org/x/sync/errgroup`. Запуск параллельных горутин с обработкой ошибок.

`g.Go(fn)` запускает задачу.

`g.Wait()` ждёт все и возвращает первую ошибку.

`errgroup.WithContext` — создаёт group с контекстом, который отменяется при первой ошибке.

`g.SetLimit(n)` ограничивает параллелизм.
