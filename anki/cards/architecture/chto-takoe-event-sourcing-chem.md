---
id: architecture_chto_takoe_event_sourcing
deck: Go
tags:
  - architecture
---

# Front

Что такое Event Sourcing? Чем отличается от CRUD?

# Back

Хранение не текущего состояния, а последовательности событий (OrderCreated, ItemAdded, OrderPaid). Состояние восстанавливается воспроизведением событий. CRUD перезаписывает данные — история теряется. ES сохраняет полную историю, позволяет воспроизвести состояние на любой момент.
