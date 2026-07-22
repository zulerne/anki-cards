---
id: design-patterns_chto_takoe_pattern_repository
deck: Go
tags:
  - design-patterns
---

# Front

Что такое паттерн Repository? Зачем отделять от бизнес-логики?

# Back

Интерфейс для доступа к данным (CRUD). Бизнес-логика зависит от интерфейса, а не от конкретной БД.

Можно подменить PostgreSQL на in-memory в тестах. Разделение ответственности: сервис не знает про SQL.
