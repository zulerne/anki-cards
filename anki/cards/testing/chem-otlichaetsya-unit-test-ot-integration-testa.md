---
id: testing_chem_otlichaetsya_unittest_ot
deck: Go
tags:
  - testing
---

# Front

Чем отличается unit-тест от integration-теста?

# Back

**Unit**: тестирует одну функцию/метод изолированно, зависимости замоканы, быстрый (мс).

**Integration**: проверяет взаимодействие компонентов (БД, HTTP, очереди), медленнее.

В Go: integration-тесты часто за build tag `//go:build integration` или `if testing.Short() { t.Skip() }`.
