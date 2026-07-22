---
id: sql_chto_takoe_cte_chem
deck: Go
tags:
  - sql
---

# Front

Что такое CTE? Чем отличается от подзапроса?

# Back

CTE (Common Table Expression):

```sql
WITH name AS (SELECT ...) SELECT ... FROM name
```

Даёт имя подзапросу, можно ссылаться несколько раз. Читабельнее вложенных подзапросов. Подзапрос — анонимный, инлайнится. CTE может быть рекурсивным (`WITH RECURSIVE`).
