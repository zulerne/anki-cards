---
id: interfaces_chto_takoe_pustoy_interfeys
deck: Go
tags:
  - interfaces
---

# Front

Что такое пустой интерфейс (any)? Из чего состоит внутренне?

# Back

any (алиас interface{}) принимает значение любого типа. Внутренне eface — два указателя: _type (информация о типе) + data (указатель на данные). Непустой интерфейс — iface: itab (тип + таблица методов) + data.
