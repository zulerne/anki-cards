---
id: redis_chto_takoe_ttl_v
deck: Go
tags:
  - redis
---

# Front

Что такое TTL в Redis? Как работает expire?

# Back

TTL (Time To Live) — время жизни ключа.

`EXPIRE key 60` — удалить через 60 секунд. `TTL key` — узнать оставшееся время.

Redis удаляет: лениво (при обращении) + активно (периодически сканирует случайные ключи).

`PERSIST key` — снять TTL.
