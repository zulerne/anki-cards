---
id: generics_kak_otlichayutsya_any_i_comparable
deck: Go
tags:
  - generics
---

# Front

Чем `any` отличается от `comparable` в generics?

# Back

`any` — алиас пустого интерфейса и допускает любой тип.

`comparable` — специальный constraint для типов, значения которых можно сравнивать оператором `==`; поэтому его можно использовать как ограничение для ключей `map` или generic-функции сравнения.
