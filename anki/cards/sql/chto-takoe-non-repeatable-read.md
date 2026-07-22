---
id: sql_chto_takoe_nonrepeatable_read
deck: Go
tags:
  - sql
---

# Front

Что такое non-repeatable read?

# Back

Транзакция A дважды читает одну строку и получает разные значения, потому что транзакция B успела изменить и зафиксировать данные между чтениями.

Предотвращается уровнем Repeatable Read и выше.
