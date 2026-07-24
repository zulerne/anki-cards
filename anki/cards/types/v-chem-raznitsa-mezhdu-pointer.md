---
id: types_v_chem_raznitsa_mezhdu
deck: Go
tags:
  - types
---

# Front

В чём разница между pointer receiver и value receiver?

# Back

Value receiver (`func (t T)`) получает копию — не может изменить оригинал.

Pointer receiver (`func (t *T)`) получает указатель — может изменять поля.

Addressable значение может вызвать pointer-метод: Go автоматически берёт его адрес. Не каждое выражение addressable. Интерфейс с таким методом реализует только `*T`, а не `T`.
