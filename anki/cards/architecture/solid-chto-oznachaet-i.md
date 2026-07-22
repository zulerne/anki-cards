---
id: architecture_solid_chto_oznachaet_i
deck: Go
tags:
  - architecture
---

# Front

SOLID: что означает I — Interface Segregation Principle?

# Back

Много маленьких интерфейсов лучше одного большого. Клиент не должен зависеть от методов, которые не использует.

Пример в Go: `io.Reader`, `io.Writer`, `io.Closer` — отдельные, а `io.ReadWriteCloser` — композиция при необходимости.
