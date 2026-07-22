---
id: sql_chto_takoe_mvcc_v
deck: Go
tags:
  - sql
---

# Front

Что такое MVCC в PostgreSQL? Зачем нужен `VACUUM`?

# Back

MVCC (Multi-Version Concurrency Control) — каждая транзакция видит свой снимок данных. При `UPDATE` создаётся новая версия строки, старая помечается как мёртвая (dead tuple).

`VACUUM` удаляет мёртвые строки и освобождает место. Без `VACUUM` таблица раздувается (table bloat).
