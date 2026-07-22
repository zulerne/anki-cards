---
id: sql_napishi_sql_naydi_dublikaty
deck: Go
tags:
  - sql
---

# Front

Напиши SQL: найди дубликаты email в таблице users.

# Back

```sql
SELECT email, COUNT(*) AS cnt
FROM users
GROUP BY email
HAVING COUNT(*) > 1;
```
