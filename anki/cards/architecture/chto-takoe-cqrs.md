---
id: architecture_chto_takoe_cqrs
deck: Go
tags:
  - architecture
---

# Front

Что такое CQRS?

# Back

Command Query Responsibility Segregation — разделение моделей чтения и записи. Команды (write) и запросы (read) используют разные модели/хранилища.

Оправдан: когда паттерны чтения и записи сильно отличаются, нужна разная оптимизация. Усложняет систему — не для простых CRUD.
