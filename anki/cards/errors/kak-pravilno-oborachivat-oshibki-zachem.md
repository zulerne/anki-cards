---
id: errors_kak_pravilno_oborachivat_oshibki
deck: Go
tags:
  - errors
---

# Front

Как правильно оборачивать ошибки? Зачем %w?

# Back

fmt.Errorf("context: %w", err). %w сохраняет цепочку — errors.Is/As могут найти оригинальную ошибку через Unwrap. %v теряет цепочку — обёрнутая ошибка недоступна.
