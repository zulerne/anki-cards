---
id: sql_chto_takoe_prepared_statement
deck: Go
tags:
  - sql
---

# Front

Что такое prepared statement? Как защищает от SQL injection?

# Back

Prepared statement — запрос с плейсхолдерами, компилируемый отдельно от параметров: SELECT * FROM users WHERE id = $1. БД парсит запрос один раз, параметры передаются отдельно и экранируются автоматически. SQL injection невозможен — параметры не могут изменить структуру запроса.
