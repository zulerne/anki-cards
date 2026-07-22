---
id: sql_nazovi_4_urovnya_izolyatsii
deck: Go
tags:
  - sql
---

# Front

Назови 4 уровня изоляции транзакций (от слабого к сильному).

# Back

1) Read Uncommitted — видны незафиксированные изменения других транзакций.2) Read Committed — видны только зафиксированные данные (default в PostgreSQL).3) Repeatable Read — повторное чтение одной строки возвращает тот же результат.4) Serializable — транзакции выполняются как последовательно, максимальная изоляция.
