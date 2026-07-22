---
id: sql_chto_takoe_foreign_key
deck: Go
tags:
  - sql
---

# Front

Что такое foreign key? Какие ON DELETE стратегии?

# Back

Foreign key — ссылка на первичный ключ другой таблицы, обеспечивает ссылочную целостность.ON DELETE стратегии:CASCADE — удалить зависимые строки.SET NULL — установить NULL.SET DEFAULT — установить значение по умолчанию.RESTRICT — запретить удаление (сразу).NO ACTION — запретить удаление (в конце транзакции, по умолчанию).
