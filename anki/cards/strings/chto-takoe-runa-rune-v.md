---
id: strings_chto_takoe_runa_rune
deck: Go
tags:
  - strings
---

# Front

Что такое руна (`rune`) в Go?

# Back

Алиас для `int32`. Представляет один Unicode code point.

Символ может занимать 1-4 байта в UTF-8, но руна всегда 4 байта.

`len("Привет")` вернёт количество байт (12).

`len([]rune("Привет"))` — количество символов (6).
