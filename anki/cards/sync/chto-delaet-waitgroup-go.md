---
id: sync_chto_delaet_waitgroup_go
deck: Go
tags:
  - sync
  - concurrency
---

# Front

Что делает `WaitGroup.Go`?

# Back

`WaitGroup.Go(f)` запускает `f` в новой горутине, увеличивает счётчик и
уменьшает его после возврата `f`.

Метод добавлен в Go 1.25. Функция `f` не должна паниковать. Для старых версий
используй `Add(1)`, `go f()` и `defer Done()`.
