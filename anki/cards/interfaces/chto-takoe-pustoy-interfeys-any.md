---
id: interfaces_chto_takoe_pustoy_interfeys
deck: Go
tags:
  - interfaces
---

# Front

Что такое пустой интерфейс (`any`)? Из чего состоит внутренне?

# Back

`any` — алиас пустого интерфейса и принимает значение любого типа.

Концептуально значение интерфейса содержит dynamic type и dynamic value.
Конкретное внутреннее представление (`eface`, `iface`, `itab`) — деталь
реализации конкретного Go runtime и не должно использоваться как контракт.
