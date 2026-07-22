---
id: sql_kakie_anomalii_predotvrashchaet_kazhdyy
deck: Go
tags:
  - sql
---

# Front

Какие аномалии предотвращает каждый уровень изоляции транзакций?

# Back

Read Uncommitted: ничего не предотвращает.

Read Committed: предотвращает dirty read.

Repeatable Read: предотвращает dirty read + non-repeatable read.

Serializable: предотвращает все (dirty read, non-repeatable read, phantom read).

Чем выше уровень — тем больше блокировок, ниже производительность.
