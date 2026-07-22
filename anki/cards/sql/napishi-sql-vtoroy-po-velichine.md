---
id: sql_napishi_sql_vtoroy_po
deck: Go
tags:
  - sql
---

# Front

Напиши SQL: второй по величине salary.

# Back

Оконная функция:SELECT salary FROM (
    SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) AS rn
    FROM employees
) t WHERE rn = 2;Без оконной функции:SELECT MAX(salary) FROM employees
WHERE salary < (SELECT MAX(salary) FROM employees);
