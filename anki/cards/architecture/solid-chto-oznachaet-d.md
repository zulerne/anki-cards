---
id: architecture_solid_chto_oznachaet_d
deck: Go
tags:
  - architecture
---

# Front

SOLID: что означает D — Dependency Inversion Principle?

# Back

Модули верхнего уровня не зависят от нижнего — оба зависят от абстракций (интерфейсов).

Пример: сервис зависит от интерфейса `Repository`, а не от `PostgresRepo`. Можно подменить реализацию без изменения бизнес-логики.
