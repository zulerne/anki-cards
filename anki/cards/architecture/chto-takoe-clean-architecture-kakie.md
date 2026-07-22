---
id: architecture_chto_takoe_clean_architecture
deck: Go
tags:
  - architecture
---

# Front

Что такое Clean Architecture? Какие слои?

# Back

Архитектура с направлением зависимостей внутрь. Слои (изнутри наружу):

1) Entities — бизнес-объекты и правила.

2) Use Cases — бизнес-логика приложения.

3) Adapters — контроллеры, репозитории, презентеры.

4) Frameworks — БД, HTTP, внешние API.

Внутренние слои не знают о внешних. Зависимости через интерфейсы.
