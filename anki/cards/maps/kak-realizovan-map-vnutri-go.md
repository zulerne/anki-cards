---
id: maps_kak_realizovan_map_vnutri
deck: Go
tags:
  - maps
---

# Front

Как реализован `map` внутри Go?

# Back

Детали реализации `map` не являются частью спецификации и могут меняться.

В стандартном runtime Go 1.24+ используется реализация на основе Swiss Tables: таблица состоит из групп по 8 слотов с control word, содержащим состояние слотов и часть хеша. Это implementation detail, поэтому нельзя полагаться на конкретный load factor, layout или стратегию роста.
