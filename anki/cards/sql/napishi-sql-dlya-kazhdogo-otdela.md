---
id: sql_napishi_sql_dlya_kazhdogo
deck: Go
tags:
  - sql
---

# Front

Напиши SQL: для каждого отдела — сотрудник с максимальной зарплатой.

# Back

```sql
SELECT department_id, name, salary
FROM (
    SELECT *, ROW_NUMBER() OVER (
        PARTITION BY department_id ORDER BY salary DESC
    ) AS rn FROM employees
) t WHERE rn = 1;
```
