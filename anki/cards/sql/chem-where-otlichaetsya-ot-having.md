---
id: sql_chem_where_otlichaetsya_ot
deck: Go
tags:
  - sql
---

# Front

Чем WHERE отличается от HAVING?

# Back

WHERE фильтрует строки до агрегации (GROUP BY).HAVING фильтрует группы после агрегации.В HAVING можно использовать агрегатные функции (HAVING COUNT(*) > 5), в WHERE — нельзя.
