---
id: interfaces_kak_v_go_realizuetsya
deck: Go
tags:
  - interfaces
---

# Front

Как в Go реализуется интерфейс?

# Back

Неявно (duck typing). Тип реализует интерфейс, если имеет все его методы. Ключевого слова implements нет. Проверка — на этапе компиляции.
