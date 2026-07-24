---
id: maps_potokobezopasen_li_vstroennyy_map
deck: Go
tags:
  - maps
---

# Front

Потокобезопасен ли встроенный `map`?

# Back

Нет. Конкурентная запись и чтение, либо две конкурентные записи без синхронизации — data race и ошибка программы; runtime часто сообщает `fatal error: concurrent map read and map write` или `concurrent map writes`, но полагаться на конкретное сообщение нельзя.

Нужен `sync.Mutex`, `sync.RWMutex` или `sync.Map`.
