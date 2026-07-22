---
id: sql_chto_takoe_partitsionirovanie_tablits
deck: Go
tags:
  - sql
---

# Front

Что такое партиционирование таблиц? Когда применять?

# Back

Разбиение большой таблицы на части (partitions) по ключу: RANGE (по дате), LIST (по значению), HASH. Запросы сканируют только нужные партиции (partition pruning). Применять: таблицы >10M строк, запросы всегда фильтруют по ключу партиции, нужна быстрая очистка старых данных (DROP PARTITION).
