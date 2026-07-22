---
id: stdlib_chto_takoe_slog_chem
deck: Go
tags:
  - stdlib
---

# Front

Что такое slog? Чем лучше log.Printf?

# Back

Пакет log/slog (Go 1.21+) — структурированное логирование. Логи как пары ключ-значение, а не строки. Поддерживает уровни (Debug, Info, Warn, Error), JSON-формат, замену handler. log.Printf — только текст, нет уровней, неудобно парсить.
