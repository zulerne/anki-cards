---
id: panic-recover_chto_takoe_error_sentinel
deck: Go
tags:
  - panic-recover
---

# Front

Что такое error sentinel? Зачем нужен?

# Back

Переменная-ошибка на уровне пакета: `var ErrNotFound = errors.New("not found")`.

Позволяет вызывающему коду проверять конкретный тип ошибки через `errors.Is(err, ErrNotFound)` вместо сравнения строк. Стабильный контракт пакета.
