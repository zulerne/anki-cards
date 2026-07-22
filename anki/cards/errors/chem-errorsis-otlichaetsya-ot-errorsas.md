---
id: errors_chem_errorsis_otlichaetsya_ot
deck: Go
tags:
  - errors
---

# Front

Чем `errors.Is` отличается от `errors.As`?

# Back

`errors.Is(err, target)` — ищет совпадение в цепочке ошибок; ошибка может
определить собственную логику через метод `Is`. Частый случай — sentinel
errors.

`errors.As(err, &target)` — проверяет тип и извлекает значение в target. Для кастомных типов ошибок.
