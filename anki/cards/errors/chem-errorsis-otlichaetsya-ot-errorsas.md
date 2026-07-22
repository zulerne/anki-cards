---
id: errors_chem_errorsis_otlichaetsya_ot
deck: Go
tags:
  - errors
---

# Front

Чем errors.Is отличается от errors.As?

# Back

errors.Is(err, target) — проверяет равенство (совпадает ли ошибка или любая обёрнутая с target). Для sentinel errors.errors.As(err, &target) — проверяет тип и извлекает значение в target. Для кастомных типов ошибок.
