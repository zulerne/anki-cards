---
id: types_v_chem_raznitsa_mezhdu
deck: Go
tags:
  - types
---

# Front

В чём разница между pointer receiver и value receiver?

# Back

Value receiver (func (t T)) получает копию — не может изменить оригинал. Pointer receiver (func (t *T)) получает указатель — может изменять поля. Значение может вызвать pointer-метод (Go автоматически берёт адрес), но интерфейс не удовлетворяется — только *T реализует интерфейс с pointer-методом.
