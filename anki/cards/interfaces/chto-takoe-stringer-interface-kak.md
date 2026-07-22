---
id: interfaces_chto_takoe_stringer_interface
deck: Go
tags:
  - interfaces
---

# Front

Что такое Stringer interface? Как кастомизировать вывод fmt.Println?

# Back

Интерфейс fmt.Stringer с методом String() string.Если тип реализует его, fmt.Println, fmt.Sprintf("%v") и т.д. вызовут этот метод. Аналог toString() из Java.Определён в пакете fmt.
