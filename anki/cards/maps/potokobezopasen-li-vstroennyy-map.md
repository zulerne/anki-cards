---
id: maps_potokobezopasen_li_vstroennyy_map
deck: Go
tags:
  - maps
---

# Front

Потокобезопасен ли встроенный `map`?

# Back

Нет. Конкурентная запись и чтение без синхронизации вызовет `fatal error: concurrent map read and map write`.

Нужен `sync.Mutex`, `sync.RWMutex` или `sync.Map`.
