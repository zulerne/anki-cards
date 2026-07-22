---
id: sql_chto_takoe_dirty_read
deck: Go
tags:
  - sql
---

# Front

Что такое dirty read?

# Back

Транзакция A читает данные, изменённые транзакцией B, которая ещё не зафиксирована (COMMIT).Если B откатится (ROLLBACK), A прочитала «грязные» — несуществующие данные.Предотвращается уровнем Read Committed и выше.
