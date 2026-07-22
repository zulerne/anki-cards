---
id: design-patterns_chem_builder_otlichaetsya_ot
deck: Go
tags:
  - design-patterns
---

# Front

Чем Builder отличается от Functional Options? Когда что использовать?

# Back

Builder: цепочка методов `b.SetPort(8080).SetTimeout(5s).Build()` — мутабельный объект, нужен метод `Build`.

Functional Options: `New(WithPort(8080))` — иммутабельная конфигурация в конструкторе.

FO идиоматичнее в Go, Builder — для сложной пошаговой сборки с валидацией.
