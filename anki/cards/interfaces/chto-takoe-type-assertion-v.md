---
id: interfaces_chto_takoe_type_assertion
deck: Go
tags:
  - interfaces
---

# Front

Что такое type assertion в Go?

# Back

Извлечение конкретного типа из интерфейса: `v := i.(T)` — если `i` не содержит тип `T`, будет `panic`.

Безопасная форма: `v, ok := i.(T)` — `ok == false` без паники.

Работает только на интерфейсных типах.
