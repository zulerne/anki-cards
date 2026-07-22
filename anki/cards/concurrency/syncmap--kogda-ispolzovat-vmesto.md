---
id: concurrency_syncmap__kogda_ispolzovat
deck: Go
tags:
  - concurrency
---

# Front

sync.Map — когда использовать вместо map + Mutex?

# Back

Два основных сценария:

1) Ключи стабильны — много чтений, редкие записи (append-only кэш).

2) Горутины работают с непересекающимися наборами ключей.

В остальных случаях `map` + `sync.RWMutex` быстрее. `sync.Map` не типизирован (ключи и значения — `any`).
