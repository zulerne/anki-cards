---
id: sql_v_chem_raznitsa_mezhdu
deck: Go
tags:
  - sql
---

# Front

В чём разница между `DELETE` и `TRUNCATE`?

# Back

`DELETE` — построчное удаление, пишет в WAL, может иметь `WHERE`, срабатывают триггеры, можно откатить в транзакции.

`TRUNCATE` — быстрая очистка всей таблицы; в PostgreSQL не пишет построчные
изменения в WAL, вызывает `ON TRUNCATE` triggers и может быть откатен в
транзакции. Identity сбрасывается только с `RESTART IDENTITY`.
