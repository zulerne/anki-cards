---
id: sql_chto_takoe_seq_scan
deck: Go
tags:
  - sql
---

# Front

Что такое Seq Scan, Index Scan, Index Only Scan? Когда какой?

# Back

Seq Scan — полный перебор таблицы. Используется когда нет индекса или нужно >10-15% строк.

Index Scan — поиск по индексу + обращение к таблице за остальными столбцами.

Index Only Scan — все нужные столбцы есть в индексе, обращение к таблице не нужно. Самый быстрый.
