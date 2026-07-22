---
id: architecture_solid_chto_oznachaet_l
deck: Go
tags:
  - architecture
---

# Front

SOLID: что означает L — Liskov Substitution Principle?

# Back

Объекты подтипа должны быть заменяемы на объекты базового типа без нарушения корректности. В Go: любая реализация интерфейса должна соблюдать контракт. Пример: если io.Reader возвращает ошибку — должна быть реальная ошибка, а не нормальный случай.
