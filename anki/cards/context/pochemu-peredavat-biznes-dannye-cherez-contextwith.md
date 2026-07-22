---
id: context_pochemu_peredavat_biznesdannye_cherez
deck: Go
tags:
  - context
---

# Front

Почему передавать бизнес-данные через context.WithValue — антипаттерн?

# Back

Нетипизированный key-value (any → any): нет проверки компилятором, легко ошибиться с ключом.Неявная зависимость — из сигнатуры не видно что функция ожидает.Усложняет тестирование.WithValue — для request-scoped метаданных (request ID, trace ID), не для бизнес-параметров. Бизнес-данные передавай явно через аргументы.
