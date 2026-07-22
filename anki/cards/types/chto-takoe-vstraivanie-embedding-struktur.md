---
id: types_chto_takoe_vstraivanie_embedding
deck: Go
tags:
  - types
---

# Front

Что такое встраивание (embedding) структур? Это наследование?

# Back

Embedding — включение типа без имени поля:

```go
type Server struct {
    http.Handler
}
```

Методы встроенного типа продвигаются (promotion) — можно вызывать напрямую.

Это не наследование: нет полиморфизма через базовый тип, нет переопределения (только затенение). Это композиция.
