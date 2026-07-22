---
id: design-patterns_chto_takoe_functional_options
deck: Go
tags:
  - design-patterns
---

# Front

Что такое Functional Options pattern?

# Back

Функции-опции для конфигурации:

```go
type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) { s.port = p }
}

srv := NewServer(WithPort(8080), WithTimeout(5*time.Second))
```

Плюсы: читаемый API, опциональные параметры, обратная совместимость. Идиоматичный Go.
