---
id: sql_chem_inner_join_otlichaetsya
deck: Go
tags:
  - sql
---

# Front

Чем `INNER JOIN` отличается от `LEFT JOIN`?

# Back

`INNER JOIN`: возвращает только строки с совпадением в обеих таблицах.

`LEFT JOIN`: возвращает все строки из левой таблицы + совпавшие из правой; если совпадения нет — столбцы правой заполняются `NULL`.
