---
id: interfaces_chto_takoe_type_switch
deck: Go
tags:
  - interfaces
---

# Front

Что такое type switch в Go?

# Back

Конструкция `switch v := i.(type)` для ветвления по конкретному типу интерфейса.

В каждом `case` переменная `v` имеет соответствующий тип.

Удобнее цепочки `if v, ok := i.(T)`.

`default` обрабатывает все остальные типы.
