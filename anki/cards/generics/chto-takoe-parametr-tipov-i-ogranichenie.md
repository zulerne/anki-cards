---
id: generics_chto_takoe_parametr_tipov_i_ogranichenie
deck: Go
tags:
  - generics
---

# Front

Что такое параметр типа и constraint в Go?

# Back

Параметр типа позволяет написать функцию или тип, работающий с несколькими типами:

```go
func Max[T cmp.Ordered](a, b T) T { ... }
```

Constraint ограничивает допустимые типы и описывает доступные операции. Начиная с Go 1.18, constraints задаются интерфейсами; `any` разрешает любой тип, `comparable` — типы, значения которых можно сравнивать через `==`.
